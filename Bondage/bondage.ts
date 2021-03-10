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
            utf8ToHex(endpoint),
            dots)
            .send({ from, gas, gasPrice });
        if (cb) {
            promiEvent.on('transactionHash', (transactionHash: string) => cb(null, transactionHash));
            promiEvent.on('error', (error: any) => cb(error));
        }

        return promiEvent;
    }