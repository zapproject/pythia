import { BaseContract } from '../BaseContract/basecontract';

import {
    BondageArgs, BondArgs, UnbondArgs, DelegateBondArgs, BondFilter,
    Filter, txid, NetworkProviderOptions, NULL_ADDRESS, NumType, TransactionCallback
} from '../Types/types';

import { Utils } from './utils';

const ethers = require('ethers');

import { assert } from 'chai';

// unsure if the line below is needed
const DEFAULT_GAS = 300000;

export class ZapBondage extends BaseContract {

    /**
     * Initializes a subclass of BaseContract that can access the methods of the Bondage contract.
     * @constructor
     * @param {NetworkProviderOptions} net. {artifactsDir ?:string|undefined,networkId?: number|undefined,networkProvider: any}
     * @param {string} net.artifactsDir - Directory where contract ABIs are located
     * @param {string} net.networkId -Select which network the contract is located on (mainnet, testnet, private)
     * @param  {Object} net.networkProvider - Ethereum network provider
     * @example new ZapBondage({networkId : 42, networkProvider : web3})
     */
    constructor(obj?: NetworkProviderOptions) {
        super(Object.assign(obj || {}, { artifactName: 'BONDAGE' }));
    }

    /**
     * @function
     * Bonds a given number of dots from a subscriber to a provider's endpoint.
     * Note: this requires that at least zapNum has been approved from the subscriber to be transferred by the Bondage contract.
     * @param  {BondArgs} bond. {provider,endpoint,dots,from,gas?=DEFAULT_GAS}
     * @param {string} bond.provider - Provider's address
     * @param {number} bond.dots Number of dots to bond to this provider
     * @param {address} bond.from - Subscriber's owner (0 broker)  or broker's address
     * @param {number} bond.gas - Sets the gas limit for this transaction (optional)
     * @param {TransactionCallback} cb - Callback for transactionHash event
     * @returns {Promise<txid>} Transaction hash.
     */
    public async bond({ provider, endpoint, dots, from, gasPrice, gas = DEFAULT_GAS }: BondArgs, cb?: TransactionCallback): Promise<txid> {
        assert(dots && dots > 0, 'Dots to bond must be greater than 0.');
        const broker = await this.contract.methods.getEndpointBroker(provider, ethers.utils.formatBytes32String(endpoint)).call();
        if (broker != NULL_ADDRESS) {
            if (from !== broker) {
                throw new Error(`Broker address ${broker} needs to call delegate bonding`);
            }
        }

        const promiEvent = this.contract.methods.bond(
            provider,
            ethers.utils.formatBytes32String(endpoint),
            //toHex(dots))
            // is formatBytes32String the proper function??
            ethers.utils.hexlify(dots))
            .send({ from, gas, gasPrice }); // not sure if we keep this or not
        if (cb) {
            promiEvent.on('transactionHash', (transactionHash: string) => cb(null, transactionHash));
            promiEvent.on('error', (error: any) => cb(error));
        }

        return promiEvent;
    }

    /**
     * Bonds a given number of dots from an account to a subscriber. This would be used to bond to a provider on behalf of another account, such as a smart contract.
     * @param {DelegateBondArgs} delegate. {provider, endpoint, dots, subscriber, from, gas= DEFAULT_GAS}
     * @param {string} delegate.provider - Provider's address
     * @param {string} delegate.endpoint - Data endpoint of the provider
     * @param {number} delegate.dots - Number of dots to bond to this provider
     * @param {address} delegate.subscriber - Address of the intended holder of the dots (subscriber)
     * @param {address} delegate.from - Address of the data subscriber
     * @param {number} [delegate.gas] - Sets the gas limit for this transaction (optional)
     * @param {TransactionCallback} cb - Callback for transactionHash event
     * @returns {Promise<txid>} Transaction hash
     */
    public async delegateBond({ provider, endpoint, dots, subscriber, from, gasPrice, gas = DEFAULT_GAS }: DelegateBondArgs, cb?: TransactionCallback): Promise<txid> {
        assert(dots && dots > 0, 'Dots to bond must be greater than 0.');
        dots = ethers.utils.hexlify(dots);
        const broker = await this.contract.methods.getEndpointBroker(provider, ethers.utils.formatBytes32String(endpoint)).call();
        if (broker != NULL_ADDRESS) {
            if (from !== broker) {
                throw new Error(`Broker address ${broker} needs to call delegate bonding for this endpoint`);
            }
        }
        const promiEvent = this.contract.methods.delegateBond(
            subscriber,
            provider,
            ethers.utils.formatBytes32String(endpoint),
            dots)
            .send({ from, gas, gasPrice }); // not sure if this is neccessary
        if (cb) {
            promiEvent.on('transactionHash', (transactionHash: string) => cb(null, transactionHash));
            promiEvent.on('error', (error: any) => cb(error));
        }

        return promiEvent;
    }

    /**
     * Unbonds a given number of dots from a provider's endpoint and transfers the appropriate amount of Zap to the subscriber.
     * @param {UnbondArgs} unbond. {provider, endpoint, dots, from, gas= DEFAULT_GAS}
     * @param {address} unbond.provider - Address of the data provider
     * @param {string} unbond.endpoint - Data endpoint of the provider
     * @param {number} unbond.dots - The number of dots to unbond from the contract
     * @param {address} unbond.from  - Address of the data subscriber
     * @param {number} unbond.gas - Sets the gas limit for this transaction (optional)
     * @param {Function} cb - Callback for transactionHash event
     * @returns {Promise<txid>} Transaction hash
     */
    public async unbond({ provider, endpoint, dots, from, gasPrice, gas = DEFAULT_GAS }: UnbondArgs, cb?: TransactionCallback): Promise<txid> {
            assert(dots && dots > 0, 'Dots to unbond must be greater than 0');
            dots = ethers.utils.hexlify(dots);
            const broker = await this.contract.methods.getEndpointBroker(provider, ethers.utils.formatBytes32String(endpoint)).call();
            if (broker != NULL_ADDRESS) {
                if (from !== broker) {
                    throw `Broker address ${broker} needs to call unbonding for this endpoint`;
                }
            }
            const promiEvent = this.contract.methods.unbond(
                provider,
                ethers.utils.formatBytes32String(endpoint),
                dots);
            if (cb) {
                promiEvent.on('transactionHash', (transactionHash: string) => cb(null, transactionHash));
                promiEvent.on('error', (error: any) => cb(error));
            }

            return promiEvent;
    }


    /************************* GETTERS ****************************/

    /**
     * Gets the number of dots that are bounded to a provider's endpoint for the current subscriber.
     * @param {BondageArgs} bond. {subscriber, provider, endpoint}
     * @param {address} bond.subscriber - Address of the data subscriber
     * @param {address} bond.provider  - Address of the data provider
     * @param {string} bond.endpoint - Data endpoint of the provider
     * @returns {Promise<string|BigNumber>} Number of bound dots to this provider's endpoint
     */
    public async getBoundDots({ subscriber, provider, endpoint }: BondageArgs): Promise < string | number > {
        return await this.contract.methods.getBoundDots(
            subscriber,
            provider,
            ethers.utils.formatBytes32String(endpoint)
        ).call();
    }

    /**
     * Calculates the amount of Zap required to bond a given number of dots to a provider's endpoint.
     * @param {BondageArgs} bondage. {provider, endpoint, dots}
     * @param {address} bondage.provider - Address of the data provider
     * @param {string} bondage.endpoint - Endpoint to calculate zap
     * @param {number} bondage.dots - Number of dots to calculate the price (in Zap) for
     * @returns {Promise<string|BigNumber>} Price (in Zap) for the given number of dots
     */
    public async calcZapForDots({ provider, endpoint, dots }: BondageArgs): Promise<string | number> {
        return await this.contract.methods.calcZapForDots(
            provider,
            ethers.utils.formatBytes32String(endpoint),
            ethers.utils.hexlify(dots)).call();
    }