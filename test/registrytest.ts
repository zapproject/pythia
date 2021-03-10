import { ZapRegistry } from '../Registry/registry';

import { join } from 'path';

import { ethers } from 'ethers';

import {
    testZapProvider,
    HardhatServerOptions,
    HardhatProvider
}
    from '../Utils/utils';
import { doesNotReject } from 'node:assert';

const expect = require('chai').expect;

describe('Registry Test', () => {

    let hardhatHttpProvider: any;
    let hardhatAccounts: any;
    let signerOne: any
    let signerTwo: any
    let signerThree: any

    let account: Array<string> = [],
        registryWrapper: any,
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

        console.log('Done running Registry tests');
    });

    it('Should be able to create registryWrapper', async () => {

        registryWrapper = new ZapRegistry(options);

        expect(registryWrapper).to.be.ok;

    });

    it('Should initiate provider in zap registry contract', async () => {

        let initProviderTx: any;

        try {

            initProviderTx = await registryWrapper.initiateProvider({

                public_key: testProvider.pubkey,
                title: testProvider.title

            }, (err: any, newProvider: string) =>

                expect(newProvider).to.be.a('string')
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

            expect(title).to.be.equal(testProvider.title);

            const pubkey = await registryWrapper.getProviderPublicKey(signerOne._address);

            expect(pubkey).to.be.equal(testProvider.pubkey);

        } catch (err) {

            const initStatus = await registryWrapper.isProviderInitiated(signerOne._address);

            if (initStatus === true) {

                console.log(signerOne._address + ': ' + 'Is already initiated as a provider');
            }
            else {

                console.log('Error with the initiateProvider function: ' + err);
            }
        }

    });

    it('Should initiate Provider curve  with 0x0 broker in zap registry contract', async () => {

        let initProviderCurveTx: any;

        try {

            initProviderCurveTx = await registryWrapper.initiateProviderCurve({

                endpoint: testZapProvider.endpoint,
                term: testZapProvider.curve.values,
                broker: testProvider.broker
            });

            const curveReceipt = await initProviderCurveTx.wait();

            expect(curveReceipt).to.include.keys('events');

            expect(curveReceipt.events[0].event).to.equal('NewCurve');

            expect(curveReceipt.events[0]).to.include.keys('args');

            const args = curveReceipt.events[0].args;

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

    it('Should set new title', async () => {

        const title = 'NEWTITLE';

        await registryWrapper.setProviderTitle({

            title: title
        });

        const newTitle = await registryWrapper.getProviderTitle(signerOne._address);

        expect(newTitle).to.equal(title);

    });


    it('Should initiate Provider curve with valid broker address in zap registry contract', async () => {

        let initProviderTwoTx: any;
        let newEndpoint: string = 'newEndpoint';

        try {

            const initCurveTwotx = await registryWrapper.initiateProviderCurve({
                endpoint: newEndpoint,
                term: testZapProvider.curve.values,
                broker: signerThree._address,
            });

        } catch (err) {

            console.log(signerTwo._address + ': ' + 'Provider and Curve is already initiated');
        }

    });

    it('Should set endpoint endpointParams in zap registry contract', async () => {

        const result = await registryWrapper.setEndpointParams(
            ethers.utils.formatBytes32String(testZapProvider.endpoint),
            testZapProvider.endpoint_params.map((params: string) => ethers.utils.formatBytes32String(params)),
        )

    });

    // it('Should clear endpoint', async () => {



    // });

})