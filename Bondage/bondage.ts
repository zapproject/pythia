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

    constructor(obj?: NetworkProviderOptions) {
        super(Object.assign(obj || {}, { artifactName: 'BONDAGE' }));
    }

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
            ethers.utils.formatBytes32String(dots))
            .send({ from, gas, gasPrice });
        if (cb) {
            promiEvent.on('transactionHash', (transactionHash: string) => cb(null, transactionHash));
            promiEvent.on('error', (error: any) => cb(error));
        }

        return promiEvent;
    
}


public async delegateBond({ provider, endpoint, dots, subscriber, from, gasPrice, gas = DEFAULT_GAS }: DelegateBondArgs, cb?: TransactionCallback): Promise<txid> {
    assert(dots && dots > 0, 'Dots to bond must be greater than 0.');
    dots = ethers.utils.formatBytes32String(dots);
    const broker = await this.contract.methods.getEndpointBroker(provider, ethers.utils.formatBytes32String(endpoint)).call();
    if (broker != NULL_ADDRESS) {
        if (from !== broker) {
            throw new Error(`Broker address ${broker} needs to call delegate bonding for this endpoint`);
        }
    }
    const promiEvent = this.contract.methods.delegateBond(
        subscriber,
        provider,
        (endpoint),
        dots)
        .send({ from, gas, gasPrice });
    if (cb) {
        promiEvent.on('transactionHash', (transactionHash: string) => cb(null, transactionHash));
        promiEvent.on('error', (error: any) => cb(error));
    }

    return promiEvent;
    }

    
public async unbond({ provider, endpoint, dots, from, gasPrice, gas = DEFAULT_GAS }: UnbondArgs, cb?: TransactionCallback): Promise<txid> {
        assert(dots && dots > 0, 'Dots to unbond must be greater than 0');
        dots = toHex(dots);
        const broker = await this.contract.methods.getEndpointBroker(provider, utf8ToHex(endpoint)).call();
        if (broker != NULL_ADDRESS) {
            if (from !== broker) {
                throw `Broker address ${broker} needs to call unbonding for this endpoint`;
            }
        }
        const promiEvent = this.contract.methods.unbond(
            provider,
            ethers.utils.formatBytes32String(endpoint),
            dots)
            .send({ from, gas, gasPrice });
        if (cb) {
            promiEvent.on('transactionHash', (transactionHash: string) => cb(null, transactionHash));
            promiEvent.on('error', (error: any) => cb(error));
        }

        return promiEvent;
    }

    public async getBoundDots({ subscriber, provider, endpoint }: BondageArgs): Promise<string | number> {
        return await this.contract.methods.getBoundDots(
            subscriber,
            provider,
            ethers.utils.formatBytes32String(endpoint)
        ).call();
    }

    public async calcZapForDots({ provider, endpoint, dots }: BondageArgs): Promise<string | number> {
        return await this.contract.methods.calcZapForDots(
            provider,
            ethers.utils.formatBytes32String(endpoint),
            toHex(dots)).call();
    }

    public async currentCostOfDot({ provider, endpoint, dots }: BondageArgs): Promise<string | number> {
        dots = toHex(dots);
        return this.contract.methods.currentCostOfDot(
            provider,
            ethers.utils.formatBytes32String(endpoint),
            dots
        ).call();
    }

    public async getDotsLimit({ provider, endpoint }: BondageArgs): Promise<string | number> {
        return await this.contract.methods.dotLimit(provider, ethers.utils.formatBytes32String(endpoint)).call().valueOf();
    }

    public async getDotsIssued({ provider, endpoint }: BondageArgs): Promise<string | number> {
        return await this.contract.methods.getDotsIssued(
            provider,
            ethers.utils.formatBytes32String(endpoint)
        ).call();
    }

    public async getBrokerAddress({ provider, endpoint }: BondageArgs): Promise<string> {
        return await this.contract.methods.getEndpointBroker(provider, ethers.utils.formatBytes32String(endpoint)).call();
    }

    public async getZapBound({ provider, endpoint }: BondageArgs): Promise<string | number> {
        return await this.contract.methods.getZapBound(provider, ethers.utils.formatBytes32String(endpoint)).call();
    }

    async getNumEscrow({ provider, endpoint, subscriber }: BondageArgs): Promise<number | string> {
        return await this.contract.methods.getNumEscrow(subscriber, provider, endpoint).call();
    }

    public listen(filters: BondFilter = {}, callback: TransactionCallback): void {
        this.contract.events.allEvents(filters, { fromBlock: 0, toBlock: 'latest' }, callback);
    }

    public listenBound(filters: BondFilter = {}, callback: TransactionCallback): void {
        this.contract.events.Bound(filters, { toBlock: 'latest' }, callback);
    }

    public listenUnbound(filters: BondFilter = {}, callback: TransactionCallback): void {
        this.contract.events.Unbond(filters, { toBlock: 'latest' }, callback);
    }

    public listenEscrowed(filters: BondFilter = {}, callback: TransactionCallback): void {
        this.contract.events.Escrowed(filters, { toBlock: 'latest' }, callback);
    }

    public listenReleased(filters: BondFilter = {}, callback: TransactionCallback): void {
        this.contract.events.Released(filters, { toBlock: 'latest' }, callback);
    }

    public listenReturned(filters: BondFilter = {}, callback: TransactionCallback): void {
        this.contract.events.Returned(filters, { toBlock: 'latest' }, callback);
    }
}