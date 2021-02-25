import BN from 'bn.js';

/**
 * @ignore
 */
export class Util {
    /**
     * @param {number} num
     * @returns {any}
     */
    static toZapBase(num: number) {
        return new BN(num).mul(new BN(10).pow(new BN(18)));
    }

    /**
     *
     * @param {number | string} num
     * @returns {number}
     */
    static fromZapBase(num: number|string):number {
        return new BN(num).div(new BN(10).pow(new BN(18))).toNumber();
    }
    static DEFAULT_GAS = 6000000;
}

