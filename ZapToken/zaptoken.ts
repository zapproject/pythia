import { BaseContract } from '../BaseContract/basecontract';

import { Util } from './tokenutils';

import {
    TransferType,
    address,
    txid,
    NetworkProviderOptions,
    TransactionCallback,
    NumType
} from '../Types/types';

import { toHex } from 'web3-utils';


/**
 * @class
 * Provides Represents an interface to the Zap Token ERC20 contract. Enables token transfers, balance lookups, and approvals.
 */

export class ZapToken extends BaseContract {

    /**
     * @constructor
     * @param obj
     * @param {string} n.artifactsDir - Directory where contract ABIs are located
     * @param {number|string} n.networkId - Select which network the contract is located on (mainnet, testnet, private)
     * @param {any} n.networkProvider - Ethereum network provider (e.g. Infura)
     * @example new ZapToken({networkId : 42, networkProvider : web3})
     */

    constructor(obj?: NetworkProviderOptions) {
        super(Object.assign(obj || {}, { artifactName: 'ZAP_TOKEN' }));
    }


    /**
     * Get the Zap Token balance of a given address.
     * @param {address} address  Address to check
     * @returns {Promise<number>} Returns a Promise that will eventually resolve into a Zap balance (wei)
     */
    async balanceOf(address: address): Promise<string | number> {

        return await this.contract.balanceOf(address);
    }

    /**
     * Transfers Zap from an address to another address.
     * @param {TransferType} t. {to, amount, from,gas=Util.DEFAULT_GAS}
     * @param {address} t.to - Address of the recipient
     * @param {number} t.amount - Amount of Zap to transfer (wei)
     * @param {address} t.from - Address of the sender
     * @param {number} t.gas - Sets the gas limit for this transaction (optional)
     * @param {Function} cb - Callback for transactionHash event
     * @returns {Promise<txid>} Returns a Promise that will eventually resolve into a transaction hash
     */
    async send({ to, amount, from, gasPrice, gas = Util.DEFAULT_GAS }: TransferType, cb?: TransactionCallback): Promise<txid> {
        amount = toHex(amount);
        const promiEvent = this.contract.transfer(to, amount).send({ from, gas, gasPrice });
        if (cb) {
            promiEvent.on('transactionHash', (transactionHash: string) => cb(null, transactionHash));
            promiEvent.on('error', (error: any) => cb(error));
        }

        return promiEvent;
    }

    /**
     * Allocates Zap Token from the Zap contract owner to an address (ownerOnly).
     * @param {TransferType} t. {to, amount, from,gas=Util.DEFAULT_GAS}
     * @param {address} t.to - Address of the recipient
     * @param {number} t.amount - Amount of Zap to allocate (wei)
     * @param {address} t.from - Address of the sender (must be owner of the Zap contract)
     * @param {number} t.gas - Sets the gas limit for this transaction (optional)
     * @param {Function} cb - Callback for transactionHash event
     * @returns {Promise<txid>} Returns a Promise that will eventually resolve into a transaction hash
     */
    async allocate({ to, amount, from, gasPrice, gas = Util.DEFAULT_GAS }: TransferType, cb?: TransactionCallback): Promise<txid> {
        amount = toHex(amount);
        const promiEvent = this.contract.allocate(to, amount).send({ from, gas, gasPrice });

        if (cb) {
            promiEvent.on('transactionHash', (transactionHash: string) => cb(null, transactionHash));
            promiEvent.on('error', (error: any) => cb(error));
        }

        return promiEvent;
    }

    /**
     * Approves the transfer of Zap Token from a holder to another account. Enables the bondage contract to transfer Zap during the bondage process.
     * @param {TransferType} t. {to, amount, from, gas=Util.DEFAULT_GAS}
     * @param {address} t.to - Address of the recipient
     * @param {number} t.amount - Amount of Zap to approve (wei)
     * @param {address} t.from - Address of the sender
     * @param {number} t.gas - Sets the gas limit for this transaction (optional)
     * @param {Function} cb - Callback for transactionHash event
     * @returns {Promise<txid>} Returns a Promise that will eventually resolve into a transaction hash
     */
    async approve({ to, amount, from, gasPrice, gas = Util.DEFAULT_GAS }: TransferType, cb?: TransactionCallback): Promise<txid> {
        amount = toHex(amount);

        const _success = this.contract.approve(to, amount).send({ from, gas, gasPrice });

        if (cb) {
            _success.on('transactionHash', (transactionHash: string) => cb(null, transactionHash));
            _success.on('error', (error: any) => cb(error));
        }

        const success = await _success;

        if (!success) {
            throw new Error('Failed to approve Bondage transfer');
        }

        return success;
    }
}
