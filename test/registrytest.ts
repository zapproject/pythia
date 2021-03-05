import { ZapRegistry } from '../Registry/registry';

import { BaseContract } from '../BaseContract/basecontract';

import { join } from 'path';

import { ethers } from 'ethers';

import {
    testZapProvider,
    HardhatServerOptions,
    HardhatProvider
}
    from '../Utils/utils';

import { sign } from 'crypto';

const expect = require('chai').expect;

describe('Registry Test', () => {

    let hardhatHttpProvider: any;
    let hardhatAccounts: any;
    let signerOne: any
    let signerTwo: any
    let signerThree: any

    let accounts: Array<string> = [],
        HardhatServer: any,
        registryWrapper: any,
        ethersjs,
        testArtifacts,
        testProvider = testZapProvider,
        buildDir: string = join(__dirname, 'contracts'),
        options: any = {
            artifactsDir: '',
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

        console.log('Done running Registry tests');
    });

    it('Should be able to create registryWrapper', async () => {

        registryWrapper = new ZapRegistry(options);

        registryWrapper = registryWrapper.contract.connect(signerOne);

        expect(registryWrapper).to.be.ok;

    });

    it('Should initiate provider in zap registry contract', async () => {

        let initProviderTx: any;

        try {

            initProviderTx = await registryWrapper.initiateProvider(
                testProvider.pubkey,
                ethers.utils.formatBytes32String(testProvider.title)
            );

            const initProviderReceipt = await initProviderTx.wait();

            expect(initProviderReceipt).to.include.keys('events');

            expect(initProviderReceipt.events[0].event).to.equal('NewProvider');

            expect(initProviderReceipt.events[0]).to.include.keys('args');

            const args = initProviderReceipt.events[0].args;

            expect(args).to.include.keys('provider', 'title');

            expect(testZapProvider.title).to.equal(ethers.utils.parseBytes32String(args.title));

            expect(args.provider).to.equal(signerOne._address);

            const title = await registryWrapper.getProviderTitle(signerOne._address);

            expect(title).to.be.equal(ethers.utils.formatBytes32String(testProvider.title));

            const pubkey = await registryWrapper.getProviderPublicKey(signerOne._address);

            expect(parseInt(pubkey)).to.be.equal(testProvider.pubkey);

        } catch (err) {

            console.log(signerOne._address + ': ' + 'Is already initiated as a provider');

        }

    });

    it('Should initiate Provider curve  with 0x0 broker in zap registry contract', async () => {

        let initProviderCurveTx: any;

        try {

            initProviderCurveTx = await registryWrapper.initiateProviderCurve(

                ethers.utils.formatBytes32String(testZapProvider.endpoint),
                testZapProvider.curve.values,
                testProvider.broker
            );

            const receipt = await initProviderCurveTx.wait();

            expect(receipt).to.include.keys('events');

            expect(receipt.events[0].event).to.equal('NewCurve');

            expect(receipt.events[0]).to.include.keys('args');

            const args = receipt.events[0].args;

            expect(args).to.include.keys('provider', 'endpoint', 'curve', 'broker');

            expect(args.broker).to.equal(testProvider.broker);

            expect(args.provider).to.equal(signerOne._address);

            const getTxCurve = args.curve.map((num: any) => parseInt(num));

            const testCurve = testProvider.curve.values;

            expect(testZapProvider.endpoint).to.equal(ethers.utils.parseBytes32String(args.endpoint));

            expect(testCurve).to.eql(getTxCurve);

        } catch (err: any) {

            console.log(signerOne._address + ': ' + 'Curve is already initiated');
        }

    });

    it('Should get all provider params', async () => {

        const params = await registryWrapper.getAllProviderParams(signerOne._address);

        console.log(params)

    });

    it('Should set new title', async () => {

        const title = ethers.utils.formatBytes32String('NEWTITLE');

        await registryWrapper.setProviderTitle(title);

        const newTitle = await registryWrapper.getProviderTitle(signerOne._address);

        expect(newTitle).to.equal(title);

    });


    it('Should initiate Provider curve with valid broker address in zap registry contract', async () => {

        registryWrapper = registryWrapper.connect(signerTwo)

        let initProviderTwoTx: any

        try {

            initProviderTwoTx = await registryWrapper.initiateProvider(
                testZapProvider.pubkey,
                ethers.utils.formatBytes32String(testZapProvider.title),
            );

            const receipt = await tx.wait();

            const initCurveTx = await registryWrapper.initiateProviderCurve(
                ethers.utils.formatBytes32String(testZapProvider.endpoint),
                testZapProvider.curve.values,
                signerThree._address,
            );

        } catch (err) {

            console.log(signerTwo._address + ': ' + 'Provider and Curve is already initiated');
        }

    });

    it('Should set endpoint endpointParams in zap registry contract', async () => {

        registryWrapper = registryWrapper.connect(signerOne)

        const result = await registryWrapper.setEndpointParams(
            ethers.utils.formatBytes32String(testZapProvider.endpoint),
            testZapProvider.endpoint_params.map((params: string) => ethers.utils.formatBytes32String(params)),
        )

    });

    it('Should clear endpoint', async () => {

        await registryWrapper.clearEndpoint(
            ethers.utils.formatBytes32String(testZapProvider.endpoint)
        );

        const clearedEndpoints = await registryWrapper.getProviderEndpoints(signerOne._address);

        clearedEndpoints.forEach((endpoint: any) => expect(endpoint)
            .to.equal('0x0000000000000000000000000000000000000000000000000000000000000000'));

    });

})