import { BigNumber } from 'bignumber.js';

/**
 * This class represents a Zap piecewise curve. Provides the functionality to parse Zap's piecewise function encoding and calculate the price of dots at any given point on the curve.
 */
export class Curve {
    public values: number[];
    public max = 0;

    /**
     * Initializes a wrapper class for function and structurizes the provided curves.
     * @constructor
     * @param {Array<number>} values Curve array
     */
    constructor(curve: number[] = []) {
        this.values = curve;
        this.checkValidity();
    }

    /**
     * Checks whether the piecewise curve encoding is valid and throws an error if invalid.
     */
    private checkValidity(): void {
        let prevEnd = 1;
        let index = 0;
        // Validate the curve
        while (index < this.values.length) {
            // Validate the length of the piece
            const len: number = this.values[index];
            if (len <= 0) throw ('Invalid curve length');

            // Validate the end index of the piece
            const endIndex: number = index + len + 1;
            if (endIndex >= this.values.length) throw ('Piece is out of bounds');

            // Validate that the end is continuous
            const end: number = this.values[endIndex];
            if (end <= prevEnd) throw ('Piece domains are overlapping');

            prevEnd = end;
            index += len + 2;
        }
        this.max = prevEnd;
    }

    /**
     * Gets the price of the nth dot. e.g. the price of a single dot to a curve with no dots issued would be calculated at n=1.
     * @param {number} total_x n - Where the new dot will be the nth dot to be bonded.
     * @returns {number} Returns the price (in Zap) of the nth dot.
     */
    public getPrice(total_x: number): number {
        if (total_x <= 0 || total_x > this.max) {
            throw ('Invalid curve supply position');
        }
        if (!this.values) {
            throw ('Curve is not initialized');
        }

        let index = 0;
        while (index < this.values.length) {
            const len: number = this.values[index];
            const end: number = this.values[index + len + 1];

            if (total_x > end) {
                // move onto the next piece
                index += len + 2;
                continue;
            }

            // calculate at this piece
            let sum = 0;
            for (let i = 0; i < len; i++) {
                const coeff: number = this.values[index + i + 1];
                sum += coeff * Math.pow(total_x, i);
            }
            return sum;
        }
        return -1;
    }

    // buying n dots starting at the ath dot
    public getZapRequired(a: number, n: number): number {
        let sum = 0;
        for (let i: number = a; i < a + n; i++) {
            sum += this.getPrice(i);
        }
        return sum;
    }

    /**
     * Converts this curve constants, parts, dividers into Array of Bignumbers
     * @returns {Array<Array<BigNumber>>}
     */
    public convertToBNArrays(): BigNumber[] {
        return this.values.map((item: number | string) => {
            return new BigNumber(item);
        });
    }

    /**
     * @ignore
     */
    public valuesToString(): string[] {
        return this.values.map((item: number) => { return '' + item; });
    }


    /**
     * Create a string representing a piecewise function
     * @param curve The curve to be stringified
     * @returns The stringified curve
     */
    public static curveToString(values: number[]): string {
        return Curve.splitCurveToTerms(values).map(Curve.termToString).join(' & ');
    }
    private static splitCurveToTerms(curve: number[]): number[][] {
        if (curve.length <= 0) return [];
        const res = [];
        let startIndex = 0;
        let currentLength = curve[0];
        let endIndex = currentLength + 2;
        while (startIndex < curve.length) {
            res.push(curve.slice(startIndex, endIndex));
            startIndex += currentLength + 2;
            currentLength = curve[endIndex];
            endIndex = startIndex + currentLength + 2;
        }
        return res;
    }
    public static termToString(term: number[]): string {
        const limit = term[term.length - 1];
        const parts = [];
        for (let i = 1; i <= term[0]; i++) {
            if (term[i] === 0) continue;
            if (term[i] === 1) {
                parts.push('x^' + (i - 1));
            } else {
                parts.push(term[i] + '*' + 'x^' + (i - 1));
            }
        }
        return parts.join('+') + '; limit = ' + limit;
    }



    public static convertToCurve(end: number, curve: string): number[] {
        if (!end || isNaN(end)) throw new Error('Start and end must be numbers');
        const tokenRegex = /\s*(x|tether|gether|mether|kether|grand|kether|zap|ether|finney|milliether|milli|szabo|microether|micro|gwei|shannon|nanoether|nano|mwei|lovelace|picoether|kwei|babbage|femtoether|wei|[0-9.]+|\S)\s*/gi;
        const terms: string[] = curve.split('+').map(term => term.trim());
        const current_curve: number[] = [];

        for (const term of terms) {
            let coef = 1;
            let exp = 0;

            const tokens: string[] = [];
            let m;
            while ((m = tokenRegex.exec(term)) !== null) {
                tokens.push(m[1]);
            }

            for (let i = 0; i < tokens.length; i++) {
                const token = tokens[i];

                if (!isNaN(+token)) {
                    coef *= +token;

                    if (i < tokens.length - 1) {
                        // https://web3js.readthedocs.io/en/1.0/web3-utils.html#fromwei
                        switch (tokens[i + 1].toLowerCase()) {
                            case 'tether':
                                coef *= 1e30;
                                i++;
                                continue;
                            case 'gether':
                                coef *= 1e27;
                                i++;
                                continue;
                            case 'mether':
                                coef *= 1e24;
                                i++;
                                continue;
                            case 'kether':
                            case 'grand':
                                coef *= 1e21;
                                i++;
                                continue;
                            case 'kether':
                                coef *= 1e21;
                                i++;
                                continue;
                            case 'zap':
                            case 'ether':
                                coef *= 1e18;
                                i++;
                                continue;
                            case 'finney':
                            case 'milliether':
                            case 'milli':
                                coef *= 1e15;
                                i++;
                                continue;
                            case 'szabo':
                            case 'microether':
                            case 'micro':
                                coef *= 1e12;
                                i++;
                                continue;
                            case 'gwei':
                            case 'shannon':
                            case 'nanoether':
                            case 'nano':
                                coef *= 1e9;
                                i++;
                                continue;
                            case 'mwei':
                            case 'lovelace':
                            case 'picoether':
                                coef *= 1e6;
                                i++;
                                continue;
                            case 'kwei':
                            case 'babbage':
                            case 'femtoether':
                                coef *= 1e3;
                                i++;
                                continue;
                            case 'wei':
                                coef *= 1;
                                i++;
                                continue;
                        }
                    }
                }
                else if (token == 'x') {
                    exp = 1;
                }
                else if (token == '*') {
                    continue;
                }
                else if (token == '^') {
                    if (i == tokens.length - 1) {
                        throw new Error('Must specify an exponent.');
                    }

                    const exponent: string = tokens[++i];

                    if (isNaN(+exponent)) {
                        throw new Error('Exponent must be a number');
                    }

                    exp = +exponent;
                }
            }

            while (current_curve.length < exp) {
                current_curve.push(0);
            }

            current_curve[exp] = coef;
        }

        return [current_curve.length, ...current_curve, end];
    }
}

export * from './types';
