import { BaseContract } from '../BaseContract/basecontract';
import {
    Filter, txid, DEFAULT_GAS, NetworkProviderOptions,
    NumType,
    DataPurchaseEvent,
    SubscriptionEndEvent,
    ParamsPassedEvent,
    SubscriptionInit,
    SubscriptionEnd,
    SubscriptionType,
    SubscriptionParams,
    TransactionCallback
} from '../Types/types';

import ethers from 'ethers';

/**
 * @class
 * Provides an interface to the Arbiter contract for managing temporal subscriptions to oracles.
 */
export class ZapArbiter extends BaseContract {

    /**
     * Initializes a subclass of BaseContract that can access the methods of the Arbiter contract.
     * @constructor
     * @augments BaseContract
     * @example new ZapArbiter({networkId : 42, networkProvider : web3})
     * @param obj: NetworkProviderOptions
     */
    constructor(obj?: NetworkProviderOptions) {
        super(Object.assign(obj || {}, { artifactName: 'ARBITER' }));
    }

    /**
     * Initializes a subscription with a given provider, endpoint, and endpoint parameters.
     * @param {SubscriptionInit} r.  {provider, endpoint, endpoint_params, blocks, pubkey, from, gas=DEFAULT_GAS}
     * @param {address} r.provider - Address of the data provider
     * @param {string} r.endpoint - Data endpoint of the provider
     * @param {Array<string>} r.endpoint_params - Params passed to endpoint
     * @param {number} r.blocks - Number of blocks that the subscription will last for
     * @param {number} r.provider - Public key of provider
     * @param {address} r.from - Subscriber's address
     * @param {number} r.gas - Set the gas limit for this transaction (optional)
     * @param {Function} cb - Callback for transactionHash event
     * @returns {Promise<txid>} Transaction hash
     */
    async initiateSubscription(
        { provider, endpoint, endpoint_params, blocks, pubkey, from, gasPrice, gas = DEFAULT_GAS }: SubscriptionInit, cb?: TransactionCallback): Promise<txid> {
        endpoint_params = endpoint_params.map((i: string) => {
            if (!ethers.utils.isHexString(i)) {
                return ethers.utils.formatBytes32String(i);
            }
            else return i;
        });
        const promiEvent = this.contract.methods.initiateSubscription(
            provider,
            ethers.utils.formatBytes32String(endpoint),
            endpoint_params,
            pubkey,
            blocks).send({ from, gas, gasPrice });
        if (cb) {
            promiEvent.on('transactionHash', (transactionHash: string) => cb(null, transactionHash));
            promiEvent.on('error', (error: any) => cb(error));
        }

        return promiEvent;
    }

    /**
     * Ends a currently active subscription for a given subscriber and endpoint from the subscriber.
     * @param {SubscriptionEnd} s. {provider, endpoint, from, gas=DEFAULT_GAS}
     * @param {address} s.provider - Address of the data provider
     * @param {string} s.endpoint - Data endpoint of the provider
     * @param {address} s.from - Address of the subscriber
     * @param {number} s.gas - Gas limit of this transaction
     * @param {Function} cb - Callback for transactionHash event
     * @returns {Promise<txid>} Transaction hash
     */
    async endSubscriptionSubscriber({ provider, endpoint, from, gasPrice, gas = DEFAULT_GAS }: SubscriptionEnd, cb?: TransactionCallback): Promise<txid> {
        const promiEvent = this.contract.methods.endSubscriptionSubscriber(
            provider,
            ethers.utils.formatBytes32String(endpoint))
            .send({ from, gas, gasPrice });
        if (cb) {
            promiEvent.on('transactionHash', (transactionHash: string) => cb(null, transactionHash));
            promiEvent.on('error', (error: any) => cb(error));
        }

        return promiEvent;
    }

    /**
     * Ends a currently active subscription for a given subscriber and endpoint from the provider.
     * @param {SubscriptionEnd} s. {subscriber, endpoint, from, gas=DEFAULT_GAS}
     * @param {address} s.subscriber - Address of the subscriber
     * @param {string} s.endpoint - Data endpoint of the provider
     * @param {address} s.from - Address of the provider
     * @param {number} s.gas - Gas limit of this transaction
     * @param {Function} cb - Callback for transactionHash event
     * @returns {Promise<txid>} Transaction hash
     */
    async endSubscriptionProvider({ subscriber, endpoint, from, gasPrice, gas = DEFAULT_GAS }: SubscriptionEnd, cb?: TransactionCallback): Promise<txid> {
        const promiEvent = this.contract.methods.endSubscriptionProvider(
            subscriber,
            ethers.utils.formatBytes32String(endpoint))
            .send({ from, gas, gasPrice });
        if (cb) {
            promiEvent.on('transactionHash', (transactionHash: string) => cb(null, transactionHash));
            promiEvent.on('error', (error: any) => cb(error));
        }

        return promiEvent;
    }

    /**
     * broadcast parameters from sender to offchain receiver
     * @param {SubscriptionParams} s. {receiver, endpoint, params, from, gas=DEFAULT_GAS}
     * @param {address} s.receiver - Address to receive parameters
     * @param {string} s.endpoint - Data endpoint of the provider
     * @param {Array<string>} s.params - Params passed to reciever
     * @param {number} s.gas - Gas limit of this transaction
     * @param {Function} cb - Callback for transactionHash event
     * @returns {Promise<txid>} Transaction hash
     */
    async passParams({ receiver, endpoint, params, from, gasPrice, gas = DEFAULT_GAS }: SubscriptionParams, cb?: TransactionCallback): Promise<txid> {
        params = params.map((i: string) => {
            if (!ethers.utils.isHexString(i)) {
                return ethers.utils.formatBytes32String(i);
            }
            else return i;
        });
        const promiEvent = this.contract.methods.passParams(receiver, ethers.utils.formatBytes32String(endpoint), params).send({ from, gas, gasPrice });
        if (cb) {
            promiEvent.on('transactionHash', (transactionHash: string) => cb(null, transactionHash));
            promiEvent.on('error', (error: any) => cb(error));
        }

        return promiEvent;
    }

    /************************* GETTERS *****************************/

    /**
     * Gets the subscription status for a given provider, subscriber, and endpoint.
     * @param {SubscriptionType} s. {provider,subscriber,endpoint}
     * @param {address} s.provider - Address of the data provider
     * @param {address} s.subscriber - Address of the subscriber
     * @param {string} s.endpoint - Data endpoint of the provider
     * @returns {Promise<string>} Information on the currently active subscription (dots,blockStart,blockEnd)
     */
    async getSubscription({ provider, subscriber, endpoint }: SubscriptionType): Promise<number[] | string[]> {
        return await this.contract.methods.getSubscription(provider, subscriber, ethers.utils.formatBytes32String(endpoint)).call();
    }

    /**
     * Get subscriber dots remaining for specified provider endpoint
     * @param provider
     * @param subscriber
     * @param endpoint
     * @returns {Promise<number|string>} Number of dots remaining
     */
    async getDots({ provider, subscriber, endpoint }: SubscriptionType): Promise<NumType> {
        return await this.contract.methods.getDots(provider, subscriber, endpoint).call();
    }

    /**
     * Get first subscription block number
     * @param provider
     * @param subscriber
     * @param endpoint
     * @returns {Promise<number|string>} First subscribed block number
     */
    async getBlockStart({ provider, subscriber, endpoint }: SubscriptionType): Promise<NumType> {
        return await this.contract.methods.getBlockStart(provider, subscriber, endpoint).call();
    }

    /**
     * Get subscription last block number
     * @param provider
     * @param subscriber
     * @param endpoint
     * @returns Block Number that subscription will end
     */
    async getPreBlockEnd({ provider, subscriber, endpoint }: SubscriptionType): Promise<NumType> {
        return await this.contract.methods.getPreBlockEnd(provider, subscriber, endpoint).call();
    }



    /******************************* EVENTS ************************************/

    /**
     * Listen for "DataSubscriptionEnd" unsubscription events
     * @param {Filter} filters
     * @param {Function} callback
     */
    listenSubscriptionEnd(filters: SubscriptionEndEvent = {}, callback: TransactionCallback): void {
        this.contract.events.DataSubscriptionEnd(
            filters,
            { fromBlock: filters.fromBlock ? filters.fromBlock : 0, toBlock: 'latest' },
            callback);
    }

    /**
     * Listen for "DataPurchase" subscription events
     * @param {Filter} filters
     * @param {()=>void} callback
     */
    listenDataPurchase(filters: DataPurchaseEvent = {}, callback: TransactionCallback): void {
        this.contract.events.DataPurchase(
            filters,
            { fromBlock: filters.fromBlock ? filters.fromBlock : 0, toBlock: 'latest' },
            callback);
    }

    listenParamsPassedEvent(filters: ParamsPassedEvent = {}, callback: TransactionCallback): void {
        this.contract.events.ParamsPassed(filters, callback);
    }


    /**
     * Listen for all Arbiter contract events based on a given filter.
     * @param {Filter} filter
     * @param {()=>void} callback
     */
    listen(filters: Record<string, unknown> = {}, callback: TransactionCallback): void {
        return this.contract.events.allEvents({ fromBlock: 0, toBlock: 'latest' }, callback);
    }

}