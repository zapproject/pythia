import { parseBytes32String } from 'ethers/lib/utils';

import { BaseContract } from '../BaseContract/basecontract';

import { Curve } from '../Curve/curve';

import {
    InitProvider, InitCurve, NextEndpoint, EndpointParams, SetProviderParams, SetProviderTitle, Endpoint,
    Filter, txid, address, NetworkProviderOptions, DEFAULT_GAS, NULL_ADDRESS, TransactionCallback
} from '../Types/types';

const ethers = require('ethers');

/**
 * Manage Providers and Curves registration
 */

export class ZapRegistry extends BaseContract {

    /**
     * @extends BaseContract
     * @param {any} artifactsDir Directory where contract ABIs are located
     * @param {any} networkId Select which network the contract is located on (mainnet, testnet, private)
     * @param {any} networkProvider Ethereum network provider (e.g. Infura)
     * @example new ZapRegistry({networkId : 42, networkProvider : web3})
     */
    constructor(obj?: NetworkProviderOptions) {
        super(Object.assign(obj || {}, { artifactName: 'REGISTRY' }));
    }

    /*************************** REGISTRY STORAGE CALLS FOR ALL PROVIDER *************************/

    /**
     * Get all providers in Registry contract.
     * @returns {Promise<Object>} Returns a Promise that will eventually resolve into list of oracles
     */
    async getAllProviders(): Promise<Record<string, any>> {
        return await this.contract.methods.getAllOracles().call();
    }

    /**
     * Look up provider's address by its index in registry storage
     * @param index
     * @returns Promise<address> address of indexed provider
     */
    async getProviderAddressByIndex(index: number | string): Promise<address> {
        return await this.contract.methods.getOracleAddress(index).call();
    }

    /****************** PROVIDER SPECIFIC CALLS **********************/

    /**
     * Initializes a brand endpoint in the Registry contract, creating an Oracle entry if needed.
     * @param {InitProvider} i. {public_key, title, from, gas=DEFAULT_GAS}
     * @param {string} i.public_key - A public identifier for this oracle
     * @param {string} i.title - A descriptor describing what data this oracle provides
     * @param {address} i.from - Ethereum Address of the account that is initializing this provider
     * @param {BigNumber} i.gas - Sets the gas limit for this transaction (optional)
     * @param {()=>void} cb - Callback for transactionHash event
     * @returns {Promise<txid>} Returns a Promise that will eventually resolve into a transaction hash
     */
    async initiateProvider({ public_key, title }: InitProvider, cb?: TransactionCallback): Promise<txid> {

        const promiEvent = this.contract.initiateProvider(

            ethers.BigNumber.from(public_key).toString(),

            ethers.utils.formatBytes32String(title)
        )

        if (cb) {
            promiEvent.on('transactionHash', (transactionHash: string) => cb(null, transactionHash));
            promiEvent.on('error', (error: any) => cb(error));
        }

        return promiEvent;
    }

    /**
     * Get a provider's public key from the Registry contract.
     * @param {address} provider - The address of this provider
     * @returns {Promise<number>} Returns a Promise that will eventually resolve into public key number
     */
    async getProviderPublicKey(provider: address): Promise<number | string> {
        const pubKey: string = await this.contract.methods.getProviderPublicKey(provider).call();
        return pubKey.valueOf(); //string
    }

    /**
     * Get a provider's title from the Registry contract.
     * @param {address} provider The address of this provider
     * @returns {Promise<string>} Returns a Promise that will eventually resolve into a title string
     */
    async getProviderTitle(provider: address): Promise<string> {
        const title = await this.contract.getProviderTitle(provider);
        return ethers.utils.parseBytes32String(title);
    }

    /**
    * Set new provider's title.
    * @param {address} provider The address of this provider
    * @param {string} title The new title of this provider
    * @param {()=>void} cb - Callback for transactionHash event
    * @returns {Promise<string>} Transaction hash
    */
    async setProviderTitle({ from, title, gas = DEFAULT_GAS }: SetProviderTitle, cb?: TransactionCallback): Promise<txid> {

        const promiEvent = this.contract.methods.setProviderTitle(

            ethers.utils.formatBytes32String(title)
        )
            .send({ from: from, gas });

        if (cb) {
            promiEvent.on('transactionHash', (transactionHash: string) => cb(null, transactionHash));
            promiEvent.on('error', (error: any) => cb(error, null));
        }

        return promiEvent;

    }

    /**
     * Gets whether this provider has already been created
     * @param {address} provider The address of this provider
     * @returns {Promise<boolean>} Returns a Promise that will eventually resolve a true/false value.
     */
    async isProviderInitiated(provider: address): Promise<boolean> {
        const created: boolean = await this.contract.methods.isProviderInitiated(provider);
        return created;
    }

    /**
     * Set the parameter of a provider
     * @param {SetProviderParams} s. { key, value, from, gas=DEFAULT_GAS }
     * @param {string} s.key - The key to be set
     * @param {string} s.value - The value to set the key to
     * @param {address} s.from - The address of the provider
     * @param {BN} s.gas - The amount of gas to use.
     * @returns {Promise<txid>} Transaction hash
     */
    async setProviderParameter({ key, value, from, gas = DEFAULT_GAS }: SetProviderParams): Promise<txid> {
        return await this.contract.methods.setProviderParameter(
            ethers.utils.formatBytes32String(key),
            ethers.utils.formatBytes32String(value)
        ).send({ from, gas });
    }

    /**
     * Get a parameter from a provider
     * @param {address} provider The address of the provider
     * @param {string} key The key you're getting
     * @returns {Promise<string>} A promise that will be resolved with the value of the key
     */
    async getProviderParam(provider: address, key: string): Promise<string> {
        return await this.contract.methods.getProviderParameter(
            provider,
            ethers.utils.formatBytes32String(key)
        ).call();
    }

    /**
     * Get all the parameters of a provider
     * @param {address} provider The address of the provider
     * @returns {Promise<string[]>} A promise that will be resolved with all the keys
     */
    async getAllProviderParams(provider: address): Promise<string[]> {
        const allParams = await this.contract.methods.getAllProviderParams(provider).call();
        return allParams;
    }

    /**
     * Get the endpoints of a given provider
     * @param {address} provider The address of this provider
     * @returns {Promise<string[]>} Returns a Promise that will be eventually resolved with the endpoints of the provider.
     */
    async getProviderEndpoints(provider: address): Promise<string[]> {
        let endpoints = await this.contract.methods.getProviderEndpoints(provider).call();
        const validEndpoints = [];
        endpoints = endpoints.map(ethers.utils.parseBytes32String);
        for (const e of endpoints) {
            if (e != '') {
                validEndpoints.push(e);
            }
        }
        return validEndpoints;
    }


    /**********************PROVIDER'S SPECIFIC ENDPOINT CALLS ***************************/


    /**
     * Initializes a piecewise curve for a given provider's endpoint. Note: curve can only be set once per endpoint.
     * @param {InitCurve} i. {endpoint, term, broker, from, gas=DEFAULT_GAS}
     * @param {string} i.endpoint - Data endpoint of the provider
     * @param {CurveType} i.curve - A curve object representing a piecewise curve
     * @param {address} i.broker - The address allowed to bond/unbond. If 0, any address allowed
     * @param {address} i.from - The address of the owner of this oracle
     * @param {BigNumber} i.gas - Sets the gas limit for this transaction (optional)
     * @param {()=>void} cb - Callback for transactionHash event
     * @returns {Promise<txid>} Returns a Promise that will eventually resolve into a transaction hash
     */
    async initiateProviderCurve({ endpoint, term, broker = NULL_ADDRESS, from, gasPrice, gas = DEFAULT_GAS }: InitCurve, cb?: TransactionCallback): Promise<txid> {
        const hex_term: string[] = [];
        for (const i in term) {
            hex_term[i] = ethers.utils.hexlify(term[i]);
        }
        const promiEvent = this.contract.methods.initiateProviderCurve(
            ethers.utils.formatBytes32String(endpoint), hex_term, broker
        )
            .send({ from, gas, gasPrice });
        if (cb) {
            promiEvent.on('transactionHash', (transactionHash: string) => cb(null, transactionHash));
            promiEvent.on('error', (error: any) => cb(error));
        }

        return promiEvent;
    }

    /**
     * Clear endpoint
     * @param {string} from The address of this provider
     * @param {string} endpoint Data endpoint of the provider
     * @param {()=>void} cb - Callback for transactionHash event
     * @returns {Promise<txid>} Transaction Hash
     */
    async clearEndpoint({ endpoint, from, gasPrice, gas = DEFAULT_GAS }: Endpoint, cb?: TransactionCallback): Promise<txid> {
        const promiEvent = this.contract.methods.clearEndpoint(
            ethers.utils.formatBytes32String(endpoint)).send({ from, gas, gasPrice }
            );
        if (cb) {
            promiEvent.on('transactionHash', (transactionHash: string) => cb(null, transactionHash));
            promiEvent.on('error', (error: any) => cb(error));
        }

        return promiEvent;
    }

    /**
     * Get a provider's endpoint's curve from the Registry contract.
     * @param {string} provider The address of this provider
     * @param {string} endpoint Data endpoint of the provider
     * @returns {Promise<CurveType>} Returns a Promise that will eventually resolve into a Curve object
     */
    async getProviderCurve(provider: string, endpoint: string): Promise<Curve> {
        const term: string[] = await this.contract.methods.getProviderCurve(
            provider,
            ethers.utils.formatBytes32String(endpoint)
        ).call();
        return new Curve(term.map((i: string) => { return parseInt(i); }));
    }

    /**
     * Initialize endpoint params for an endpoint. Can only be called by the owner of this oracle.
     * @param {EndpointParams} e. {endpoint, endpoint_params, from, gas=DEFAULT_GAS}••••••••••
     * @param {string} e.endpoint - Data endpoint of the provider
     * @param {string[]} e.endpoint_params - The parameters that this endpoint accepts as query arguments
     * @param {address} e.from - The address of the owner of this oracle
     * @param {BigNumber} e.gas - Sets the gas limit for this transaction (optional)
     * @param {()=>void} cb - Callback for transactionHash event
     * @returns {Promise<txid>} Returns a Promise that will eventually resolve into a transaction hash
     */
    async setEndpointParams({ endpoint, endpoint_params = [], from, gasPrice, gas = DEFAULT_GAS }: EndpointParams, cb?: TransactionCallback): Promise<txid> {
        const params = ZapRegistry.encodeParams(endpoint_params);
        const promiEvent = this.contract.methods.setEndpointParams(
            ethers.utils.formatBytes32String(endpoint),
            params
        ).send({ from, gas, gasPrice });
        if (cb) {
            promiEvent.on('transactionHash', (transactionHash: string) => cb(null, transactionHash));
            promiEvent.on('error', (error: any) => cb(error));
        }

        return promiEvent;
    }

    /**
     * Get a provider endpoint's broker address
     * @param {address} provider The address of this provider
     * @param {string} endpoint - Endpoint to query broker's address
     * @returns {Promise<address>} Returns a Promise of broker's address or Null_address if there is
     * no broker for this endpoint
     */
    async getEndpointBroker(provider: address, endpoint: string): Promise<string> {
        const broker = await this.contract.methods.getEndpointBroker(
            provider, ethers.utils.formatBytes32String(endpoint)
        ).call();
        return ethers.utils.parseBytes32String(broker);
    }

    /**
     * Gets whether this endpoint and its corresponding curve have already been set
     * @param {address} provider The address of this provider
     * @param {string} endpoint - Endpoint's string
     * @returns {Promise<boolean>} Returns a Promise of a true/false value.
     */
    async isEndpointSet(provider: address, endpoint: string): Promise<boolean> {
        const unset: boolean = await this.contract.methods.getCurveUnset(
            provider, ethers.utils.formatBytes32String(endpoint)
        ).call();
        return !unset;
    }

    /**
     * Get the endpoint params at a certain index of a provider's endpoint.
     * @param {address} provider The address of this provider
     * @param {string} endpoint Data endpoint of the provider
     * @returns {Promise<string>} Returns a Promise of the endpoint's all params
     */
    async getEndpointParams({ provider, endpoint }: NextEndpoint): Promise<string[]> {
        const params: string[] = await this.contract.methods.getEndpointParams(
            provider,
            ethers.utils.formatBytes32String(endpoint)
        ).call();

        return ZapRegistry.decodeParams(params);
    }

    /**
     * Used to encode params that has length more than 32 bytes
     * Saved into indexed ordered array
     * @param endpointParams
     */
    private static encodeParams(endpointParams: string[] = []): string[] {
        const hexParams = endpointParams.map(el => el.indexOf('0x') === 0 ? el : ethers.utils.formatBytes32String(el));
        const bytesParams: number[][] = hexParams.map(ethers.utils.arrayify);
        const params: string[] = [];
        bytesParams.forEach((el: number[]) => {
            if (el.length <= 32) {
                params.push(ethers.utils.hexlify(el));
                return;
            }
            const chunksLength = Math.ceil((el.length + 2) / 32);
            const paramBytesWithLength = [0, chunksLength].concat(el);
            for (let i = 0; i < chunksLength; i++) {
                const start = i * 32;
                const end = start + 32;
                params.push(ethers.utils.hexlify(paramBytesWithLength.slice(start, end)));
            }
        });
        return params;
    }

    /**
     * Used to decode params that is longer than 32 bytes
     * Retrieve if the params is indexed ordered array
     * @param rawParams
     */
    private static decodeParams(rawParams: string[] = []): string[] {
        const bytesParams: number[][] = rawParams.map(ethers.utils.arrayify);
        const params: string[] = [];
        let i = 0;
        const len = bytesParams.length;
        while (i < len) {
            // check if the first byte is 0, the second byte is number and is larger than 1 (chuncks count), and the length must be 32 bytes
            const isStartOfChunks = bytesParams[i][0] === 0 && bytesParams[i][1] > 1 && bytesParams[i].length === 32;
            if (!isStartOfChunks) {
                params.push(ethers.utils.hexlify(bytesParams[i]));
                i++;
                continue;
            }
            const chunksLength = bytesParams[i][1];
            const end = i + chunksLength;
            // strip zero byte and chunks length from the beginning
            let bytes = bytesParams[i].slice(2);
            i++;
            while (i < end) {
                bytes = bytes.concat(bytesParams[i]);
                i++;
            }
            params.push(ethers.utils.hexlify(bytes));
        }
        return params.map(hex => {
            try {
                return ethers.utils.parseBytes32String(hex);
            } catch (e) {
                console.log(e);
            }
            return hex;
        });
    }



    /********************** EVENTS ***************/

    /**
     * Listen to all Registry contract events
     * @param {()=>void} callback Callback function that is called whenever an event is emitted
     */
    async listen(callback: () => void): Promise<void> {
        this.contract.events.allEvents(callback);
    }

    /**
     * Listen to Registry contracts events for new providers
     * @param filters : {provider:address,title:bytes32}
     * @param callback function
     */
    async listenNewProvider(filters: Filter = {}, callback: TransactionCallback): Promise<void> {
        this.contract.events.NewProvider(filters, callback);
    }

    /**
     * Listen to Registry contract's events for new providers' curve
     * @param filters : {provider:address, endpoint:bytes32, curve:int[], broker:address}
     * @param callback
     */
    async listenNewCurve(filters: Filter, callback: TransactionCallback): Promise<void> {
        this.contract.events.NewCurve(filters, callback);
    }


}
