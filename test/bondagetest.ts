import { ZapBondage } from '../Bondage/bondage';

import { join } from 'path';

import { ethers } from 'ethers';

import {
    testZapProvider,
    HardhatServerOptions,
    HardhatProvider
}
    from '../Utils/utils';
import { doesNotReject } from 'node:assert';
import { Utils } from '../BaseContract/utils';

const expect = require('chai').expect;

describe('Bondage Test', () => {

    let hardhatHttpProvider: any;
    let hardhatAccounts: any;
    let signerOne: any
    let signerTwo: any
    let signerThree: any

    let accounts: Array<string> = [],
        HardhatServer: any,
        bondageWrapper: any,
        testProvider = testZapProvider,
        options: any = {
            networkId: HardhatServerOptions.network_id,
            networkProvider: HardhatProvider
        };

    beforeEach(async () => {

        hardhatHttpProvider = new ethers.providers.JsonRpcProvider(options.networkProvider);

        hardhatAccounts = await hardhatHttpProvider.listAccounts();

        signerOne = await hardhatHttpProvider.getSigner(hardhatAccounts[0]);

        signerTwo = await hardhatHttpProvider.getSigner(hardhatAccounts[1]);

        signerThree = await hardhatHttpProvider.getSigner(hardhatAccounts[2]);

    });

    after(() => {

        console.log('Done running Bondage tests');
    });

    it('Should be able to create bondageWrapper', async () => {

        bondageWrapper = new ZapBondage(options);

        bondageWrapper = bondageWrapper.contract.connect(signerOne)

        expect(bondageWrapper).to.be.ok;

    });

    // it('Should initiate bondageWrapper through Coordinator address', async () => {

    //     options.ZapCoordinator = bondageWrapper.ZapCoordinator._address;
    //     bondageWrapper = new ZapBondage(options);
    //     //await Utils.delay(3000);

    //     expect(bondageWrapper).to.be.ok;
    // })

    // it('A newly created provider should have no bound dots', async () => {
    //     const boundDots = await bondageWrapper.getBoundDots({subscriber: accounts[2], provider: accounts[0], endpoint: testZapProvider.endpoint});
    //     expect(boundDots).to.equal('0');
    // });

    // it('Ensure that the total zap bound to an unbonded provider is 0', async function() {
    //     //const boundZap = await bondageWrapper.getZapBound
    // });

        it('Should calculate the correct amount of zap for 5 dots on a particular endpoint', async () => {
            // requiredZap = await bondageWrapper.calcZapForDots({
            //     provider: accounts[0],
            //     endpoint: testZapProvider.endpoint,
            //     dots: 5
            // });
    
            // expect(requiredZap).to.equal('85');
        });

        it("Should bond (without broker) the correct amount of zap to recieve 5 dots", async () => {
            // requiredZap = await bondageWrapper.calcZapForDots({
            //     provider: accounts[0],
            //     endpoint: testZapProvider.endpoint,
            //     dots: 5
            // });
            // // approve
            // await deployedToken.contract.methods.approve(deployedBondage.contract._address, requiredZap).send({from: accounts[2], gas:200000 });
            // const bonded = await bondageWrapper.bond({
            //     provider: accounts[0],
            //     endpoint: testZapProvider.endpoint,
            //     dots: 5,
            //     from: accounts[2],
            //     gas: 600000
            // }, (err: any, txid: string) => expect(txid).to.be.a('string'));
            // const numZap = bonded.events.Bound.returnValues.numZap;
            // const numDots = bonded.events.Bound.returnValues.numDots;
    
            // const boundDots = await bondageWrapper.getBoundDots({
            //     subscriber: accounts[2],
            //     provider: accounts[0],
            //     endpoint: testZapProvider.endpoint
            // });
            // expect(numZap).to.equal('85');
            // expect(numDots).to.equal('5');
            // return;
        });

        it('Should bond (with broker) the correct amount of zap to recieve 5 dots', async () => {
            // requiredZap = await bondageWrapper.calcZapForDots({
            //     provider: accounts[0],
            //     endpoint: endpointB,
            //     dots: 5
            // });
            // // approve
            // await deployedToken.contract.methods.approve(deployedBondage.contract._address, requiredZap).send({from: broker, gas: 600000 });
            // const bonded = await bondageWrapper.bond({
            //     provider: accounts[0],
            //     endpoint: endpointB,
            //     dots: 5,
            //     from: broker,
            //     gas: 600000
            // }, (err: any, txid: string) => expect(txid).to.be.a('string'));
            // const numZap = bonded.events.Bound.returnValues.numZap;
            // const numDots = bonded.events.Bound.returnValues.numDots;
    
            // const boundDots = await bondageWrapper.getBoundDots({
            //     subscriber: accounts[2],
            //     provider: accounts[0],
            //     endpoint: endpointB
            // });
            // expect(numZap).to.equal('85');
            // expect(numDots).to.equal('5');
            // return;
        });

        it('Should unbond (without broker) 1 dot(s) and return the correct amount of zap recieved', async () => {
            // const preAmt = await deployedToken.contract.methods.balanceOf(accounts[2]).call().valueOf();

            // const unbonded = await bondageWrapper.unbond({
            //     provider: accounts[0],
            //     endpoint: testZapProvider.endpoint,
            //     dots: 1,
            //     from: accounts[2],
            //     gas: 600000
            // }, (err: any, txid: string) => expect(txid).to.be.a('string'));
    
            // const postAmt = await deployedToken.contract.methods.balanceOf(accounts[2]).call();
            // const diff = new BigNumber(postAmt).minus(new BigNumber(preAmt)).toString();
            // expect(diff).to.equal('35');
        });

        it('Should unbond (with broker) 1 dot(s) and return the correct amount of zap recieved', async () => {
        //     const preAmt = await deployedToken.contract.methods.balanceOf(accounts[2]).call().valueOf();
        // const preAmtBroker = await deployedToken.contract.methods.balanceOf(broker).call().valueOf();

        // const unbonded = await bondageWrapper.unbond({
        //     provider: accounts[0],
        //     endpoint: endpointB,
        //     dots: 1,
        //     from: broker,
        //     gas: 600000
        // }, (err: any, txid: string) => expect(txid).to.be.a('string'));

        // const postAmt = await deployedToken.contract.methods.balanceOf(accounts[2]).call();
        // const postAmtBroker = await deployedToken.contract.methods.balanceOf(broker).call();
        // const diff = new BigNumber(postAmt).minus(new BigNumber(preAmt)).toString();
        // const brokerDiff = new BigNumber(postAmtBroker).minus(new BigNumber(preAmtBroker)).toString();
        // expect(diff).to.equal('0');
        // expect(brokerDiff).to.equal('35');
        });

        it('Should fail to unbond endpoint with broker assigned from subscriber ', async () => {
            // try {
            //     await bondageWrapper.unbond({
            //         provider: accounts[0],
            //         endpoint: endpointB,
            //         dots: 1,
            //         from: accounts[2],
            //         gas:600000
            //     });
            // } catch (e){
            //     expect(e).to.deep.equal(`Broker address ${broker} needs to call unbonding for this endpoint`);
            // }
        });

        it("Check that issued dots increase with every bond performed", async () => {
        //     const startDots = await bondageWrapper.getBoundDots({subscriber: accounts[2], provider: accounts[0], endpoint: testZapProvider.endpoint });

        // await deployedToken.contract.methods.approve(deployedBondage.contract._address, 50).send({from: accounts[2], gas: Utils.Constants.DEFAULT_GAS });
        // const bonded = await bondageWrapper.bond({
        //     provider: accounts[0],
        //     endpoint: testZapProvider.endpoint,
        //     dots: 1,
        //     from: accounts[2]
        //});
        // const finalDots = await bondageWrapper.getBoundDots({subscriber: accounts[2], provider: accounts[0], endpoint: testZapProvider.endpoint });
        // expect(finalDots - startDots).to.equal(1);
        });

        it("Check that issued dots will decrease with every unbond", async () => {
            // const startDots = await bondageWrapper.getBoundDots({subscriber: accounts[2], provider: accounts[0], endpoint: testZapProvider.endpoint });
            // const unbonded = await bondageWrapper.unbond({
            //     provider: accounts[0],
            //     endpoint: testZapProvider.endpoint,
            //     dots: 1,
            //     from: accounts[2]
            // });
    
            // const finalDots = await bondageWrapper.getBoundDots({subscriber: accounts[2], provider: accounts[0], endpoint: testZapProvider.endpoint });
            // expect(finalDots - startDots).to.equal(-1);
        });

        it("Check that you cannot unbond more dots than you have", async () => {
            // const startDots:number = await bondageWrapper.getBoundDots({
            //     subscriber: accounts[2],
            //     provider: accounts[0],
            //     endpoint: testZapProvider.endpoint
            // });
            // try {
            //     await bondageWrapper.unbond({
            //         provider: accounts[0],
            //         endpoint: testZapProvider.endpoint,
            //         dots: 100,
            //         from: accounts[2]
            //     });
            // } catch (e){
            //     expect(e.toString()).to.include('revert');
            //     const finalDots = await bondageWrapper.getBoundDots({subscriber: accounts[2], provider: accounts[0], endpoint: testZapProvider.endpoint });
            //     expect(finalDots).to.equal(startDots);
            // }
    
        });

        it("Check that you can delegateBond", async () => {
        //     const startDots = await bondageWrapper.getBoundDots({subscriber: accounts[1], provider: accounts[0], endpoint: testZapProvider.endpoint });

        // await deployedToken.contract.methods.approve(deployedBondage.contract._address, 50).send({from: accounts[2], gas: Utils.Constants.DEFAULT_GAS });
        // const bonded = await bondageWrapper.delegateBond({
        //     provider: accounts[0],
        //     endpoint: testZapProvider.endpoint,
        //     dots: 1,
        //     subscriber: accounts[1],
        //     from: accounts[2]
        // });

        // const finalDots = await bondageWrapper.getBoundDots({subscriber: accounts[1], provider: accounts[0], endpoint: testZapProvider.endpoint });
        // expect(finalDots - startDots).to.equal(1);
        });

        it("Should be able to bond more than 10^23 wei zap", async () => {
            // const dotsLimit = await bondageWrapper.getDotsLimit({
            //     provider: accounts[0],
            //     endpoint: testZapProvider.endpoint
            // });
            // const dotsIssued = await bondageWrapper.getDotsIssued({provider: accounts[0], endpoint: testZapProvider.endpoint });
            // const availableDots = dotsIssued * 1000000;
            // const zapForDots:string = await bondageWrapper.calcZapForDots({
            //     provider: accounts[0],
            //     endpoint: testZapProvider.endpoint,
            //     dots: availableDots
            // });
            // await deployedToken.contract.methods.approve(deployedBondage.contract._address, toHex(zapForDots)).send({from: accounts[3], gas: Utils.Constants.DEFAULT_GAS });
            // const bond = await bondageWrapper.bond({
            //     provider: accounts[0],
            //     endpoint: testZapProvider.endpoint,
            //     dots: availableDots,
            //     from: accounts[3]
            // });
        });
        // it("13) Check that bonding without approval will fail", async() => {
        //     let allowance = await deployedToken.contract.methods.allowance(accounts[2],deployedBondage.contract._address).call().valueOf();
        //     if(allowance > 0) await deployedToken.contract.methods.decreaseApproval(deployedBondage.contract, allowance, {from: accounts[2]});
        //     // will revert
        //     await expect(function(){
        //         bondageWrapper.bond({
        //             provider:accounts[0],
        //             endpoint:testZapProvider.endpoint,
        //             zapNum:100,
        //             from:accounts[2]
        //         });
        //     }).to.throw(Error);
        // });
})

