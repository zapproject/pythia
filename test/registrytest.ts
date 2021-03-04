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

    })

    after(() => {

        console.log('Done running Registry tests');
    });

    it('Should be able to create registryWrapper', async () => {

        registryWrapper = new ZapRegistry(options);

        expect(registryWrapper).to.be.ok;

    });

    it('Should initiate provider in zap registry contract', async () => {

        registryWrapper = registryWrapper.contract.connect(signerOne);

        const tx = await registryWrapper.initiateProvider(
            testProvider.pubkey,
            ethers.utils.formatBytes32String(testProvider.title)
        )

        expect(tx).to.be.a('object');

        expect(registryWrapper).to.include.keys('filters');

        expect(registryWrapper.filters).to.include.keys('NewProvider');

        expect(registryWrapper.filters.NewProvider(
            signerOne._address,
            ethers.utils.formatBytes32String(testProvider.title)
        )).to.include.keys('topics');

        const title = await registryWrapper.getProviderTitle(signerOne._address);

        expect(title).to.be.equal(ethers.utils.formatBytes32String(testProvider.title));

        const pubkey = await registryWrapper.getProviderPublicKey(signerOne._address);

        expect(parseInt(pubkey)).to.be.equal(testProvider.pubkey);

    });

    it('Should initiate Provider curve  with 0x0 broker in zap registry contract', async () => {

        let tx: any;

        tx = await registryWrapper.initiateProviderCurve(
            ethers.utils.formatBytes32String(testZapProvider.endpoint),
            testZapProvider.curve.values,
            testProvider.broker
        );

        const receipt = await tx.wait();

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

        const tx = await registryWrapper.initiateProvider(
            testZapProvider.pubkey,
            ethers.utils.formatBytes32String(testZapProvider.title),
        );

        const receipt = await tx.wait();

        const initCurveTx = await registryWrapper.initiateProviderCurve(
            ethers.utils.formatBytes32String(testZapProvider.endpoint),
            testZapProvider.curve.values,
            signerThree._address,
        )


    });


})