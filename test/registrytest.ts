import { ZapRegistry } from '../Registry/registry';

import { join } from 'path';

import { ethers } from 'ethers';

import {
    delay,
    testZapProvider,
    HardhatServerOptions,
    HardhatProvider
}
    from '../Utils/utils';
import { doesNotReject } from 'node:assert';
import { collapseTextChangeRangesAcrossMultipleVersions } from 'typescript';

const expect = require('chai').expect;

describe('Registry Test', () => {

    let hardhatHttpProvider: any;
    let hardhatAccounts: any;
    let signerOne: any
    let signerTwo: any

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

    });

    after(() => {

        console.log('Done running Registry tests');

    });

    it('Should be able to create registryWrapper', async () => {

        registryWrapper = new ZapRegistry(options);

        expect(registryWrapper).to.be.ok;

    });

    it('should be able to create registryWrapper with coordinator', async () => {

        registryWrapper = new ZapRegistry({
            networkId: HardhatServerOptions.network_id,
            networkProvider: HardhatProvider,
            coordinator: registryWrapper.coordinator.address
        });

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

            expect(testProvider.title).to.equal(ethers.utils.parseBytes32String(args.title));

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

                endpoint: testProvider.endpoints[0],
                term: testProvider.curve.values,
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

            expect(testProvider.endpoints[0]).to.equal(ethers.utils.parseBytes32String(args.endpoint));

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

        try {

            const initCurveTwotx = await registryWrapper.initiateProviderCurve({
                endpoint: testProvider.endpoints[1],
                term: testProvider.curve.values,
                broker: signerTwo._address,
            });

        } catch (err) {

            console.log(signerTwo._address + ': ' + 'Provider and Curve is already initiated');
        }

    });

    it('Should get the provider curve', async () => {

        const getCurve = await registryWrapper.getProviderCurve(
            signerOne._address,
            testProvider.endpoints[0]
        );

        expect(getCurve).to.ok;
        expect(getCurve.values).to.eql(testProvider.curve.values);

    });

    it('Should get endpoint broker', async () => {

        const getBroker = await registryWrapper.getEndpointBroker(
            signerOne._address,
            testProvider.endpoints[0]
        );

        expect(getBroker).to.be.ok;
        expect(getBroker).to.be.equal(testProvider.broker);

    });

    it('Should set endpoint endpointParams in zap registry contract', async () => {

        const setParamsTx = await registryWrapper.setEndpointParams({
            endpoint: testProvider.endpoints[0],
            endpoint_params: testProvider.endpoint_params,
        });

    });

    it('Should get the endpoint params in zap registry contract', async () => {

        const endpointParams = await registryWrapper.getEndpointParams({
            provider: signerOne._address,
            endpoint: testProvider.endpoints[0]
        });

        expect(endpointParams).to.be.ok;

        expect(endpointParams).to.eql(testProvider.endpoint_params);

    });

    it('Should set the markdown url for the first endpoint param ', async () => {

        const setMdTx = await registryWrapper.setProviderParameter({
            key: testProvider.endpoint_params[0],
            value: testProvider.markdownUrl
        });

        expect(setMdTx).to.be.ok;

    });

    it('Should set the json url for the second endpoint param ', async () => {

        const setJsonTx = await registryWrapper.setProviderParameter({
            key: testProvider.endpoint_params[1],
            value: testProvider.jsonUrl
        });

        expect(setJsonTx).to.be.ok;

    });

    it('Should get the markdown url from the first endpoint param', async () => {

        const markdownParam = await registryWrapper.getProviderParam(
            signerOne._address,
            testProvider.endpoint_params[0]
        );

        expect(markdownParam).to.equal(testProvider.markdownUrl);

    });

    it('Should get the json url from the second endpoint param', async () => {

        const jsonParam = await registryWrapper.getProviderParam(
            signerOne._address,
            testProvider.endpoint_params[1]
        );

        expect(jsonParam).to.equal(testProvider.jsonUrl);

    });


    it('Should get all provider params', async () => {

        const getProviderParamsTx = await registryWrapper.getAllProviderParams(signerOne._address);

        expect(getProviderParamsTx).to.eql(testProvider.endpoint_params);

    });

    it('Should be able to get all providers', async () => {

        const getProviderTx = await registryWrapper.getAllProviders();

        expect(getProviderTx).to.be.ok;

    });

    it('Should get provider address by index', async () => {

        const providers = [];

        const allProviders = await registryWrapper.getAllProviders();

        for (let i = 0; i < allProviders.length; i++) {

            providers.push(await registryWrapper.getProviderAddressByIndex(i));

        }

        expect(providers).to.eql(allProviders);

    });

    it('Should check if all providers are initiated', async () => {

        const initStatus = [];

        const allProviders = await registryWrapper.getAllProviders();

        for (let i = 0; i < allProviders.length; i++) {

            initStatus.push(await registryWrapper.isProviderInitiated(allProviders[i]));

        };

        expect(initStatus.forEach(bool => bool === true));

    });

    it('Should check if the endpoint and corresponding curve is set', async () => {

        const unset = await registryWrapper.isEndpointSet(
            signerOne._address,
            testProvider.endpoints[0]
        );

        expect(unset).to.be.true;

    });

    it('Should clear the first endpoint', async () => {

        const clearFirstEndpnt = await registryWrapper.clearEndpoint({

            endpoint: testProvider.endpoints[0]
        });

        expect(clearFirstEndpnt).to.be.ok;

    });

    it('Should clear the last endpoint', async () => {

        const clearScndEndpnt = await registryWrapper.clearEndpoint({

            endpoint: testProvider.endpoints[1]
        })

        expect(clearScndEndpnt).to.be.ok;

    });

    it('Should have an empty array of endpoints after clearing', async () => {

        const getEndpnts = await registryWrapper.getProviderEndpoints(signerOne._address);

        expect(getEndpnts).to.be.ok;
        expect(getEndpnts.length).to.equal(0);
    });

})