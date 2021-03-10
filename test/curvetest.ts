import {Curve} from '../Curve/curve';

const expect = require('chai').expect;

describe('Zap Curve Test', function() {
    let terms: number[];
    let curve:Curve;


    it('1) Try to initialize a simple curve', () => {
        terms = [3, 0, 2, 1, 100]; // y=2x+x^2 from [1,100]
        curve = new Curve(terms);
    });

    it('2) Perform a simple calculation', () => {

        const curve = new Curve(terms);

        const x = 10;
        //expect(curve.getPrice(x)).to.equal(2 * x ** 2);
    });

    it('3) Perform calculations with a curve with multiple pieces', () => {
        terms = [4, 0, 1, 0, 1, 100000]; // x+xx^3 range [1,100000]
        curve = new Curve(terms);

        // test piece 1
        let x = 10;
        //expect(curve.getPrice(x)).to.equal(2 * x ** 2);
        // test piece 2
        x = 40;
        //expect(curve.getPrice(x)).to.equal(3 * x ** 3);
    });
    
    it('4) Perform calculations with a curve with multiple parts to a piece', () => {
        terms = [4, 0, 0, 2, 3, 100000]; // 2x^2+3x^3
        curve = new Curve(terms);
        const x = 5;
        //expect(curve.getPrice(x)).to.equal(2 * x ** 2 + 3 * x ** 3);
    });

    it('5) Perform calculations with a curve with multiple parts to multiple pieces', () => {

    });

    it('6) Attempt to create curves with incorrect number of arguments', () => {
        terms = [2, 2, 0, 3, 3, 0, 1]; // wrong number of constants


        const bad = function() { new Curve(terms); };
        //expect(bad).to.throw(/Invalid number of constants arguments/);

    });

    it('7) Attempt to create curves with bad arguments', () => {

    });

    it('8) Attempt to convert curve values to string representation', () => {
        const terms = [1, 1, 10, 2, 0, 1, 20, 3, 0, 0, 1, 30, 4, 0, 0, 0, 1, 40, 5, 0, 0, 0, 0, 1, 50];
        const curveString = 'x^0; limit = 10 & x^1; limit = 20 & x^2; limit = 30 & x^3; limit = 40 & x^4; limit = 50';
        expect(Curve.curveToString(terms)).to.equal(curveString);
    });

    it('9) Attempt to convert single curve term to string', () => {
        const terms = [2, 0, 10000000000000000, 7000000000000000000];
        const curveString = '10000000000000000*x^1; limit = 7000000000000000000';
        expect(Curve.termToString(terms)).to.equal(curveString);
    });

    it('10) Attempt to create Curve from string formula', () => {
        const formula = '10+x^3';
        const end = 100;
        expect(Curve.convertToCurve(end, formula)).to.deep.equal([4, 10, 0, 0, 1, end]);
    });

});
