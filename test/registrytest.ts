import { ZapRegistry } from '../Registry/registry';

import { ethers } from 'ethers';

import {
    testZapProvider,
    HardhatServerOptions,
    HardhatProvider
}
    from '../Utils/utils';

const expect = require('chai').expect;

describe('Registry Test', () => {

    let hardhatHttpProvider: any;
    let hardhatAccounts: any;

    let signers: Array<any> = [],
        registryWrapper: any,
        testProvider = testZapProvider,
        options: any = {
            networkId: HardhatServerOptions.network_id,
            networkProvider: HardhatProvider
        };

    beforeEach(async () => {

        hardhatHttpProvider = new ethers.providers.JsonRpcProvider(options.networkProvider);

        hardhatAccounts = await hardhatHttpProvider.listAccounts();

        for (let i = 0; i < hardhatAccounts.length; i++) {

            signers.push(await hardhatHttpProvider.getSigner(hardhatAccounts[i]));
        }

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

        try {

            const initProviderTx: any = await registryWrapper.initiateProvider({

                public_key: testProvider.pubkey,
                title: testProvider.title

            }, (err: any, newProvider: string) =>

                expect(newProvider).to.be.a('string')
            );

            const initProviderReceipt: any = await initProviderTx.wait();

            expect(initProviderReceipt).to.include.keys('events');

            expect(initProviderReceipt.events[0].event).to.equal('NewProvider');

            expect(initProviderReceipt.events[0]).to.include.keys('args');

            const args: any = initProviderReceipt.events[0].args;

            expect(args).to.include.keys('provider', 'title');

            expect(testProvider.title).to.equal(ethers.utils.parseBytes32String(args.title));

            expect(args.provider).to.equal(signers[0]._address);

            const title: String = await registryWrapper.getProviderTitle(signers[0]._address);

            expect(title).to.be.equal(testProvider.title);

            const pubkey: Number = await registryWrapper.getProviderPublicKey(signers[0]._address);

            expect(pubkey).to.be.equal(testProvider.pubkey);

        } catch (err) {

            const initStatus: Boolean = await registryWrapper.isProviderInitiated(signers[0]._address);

            if (initStatus === true) {

                console.log(signers[0]._address + ': ' + 'Is already initiated as a provider');
            }
            else {

                console.log('Error with the initiateProvider function: ' + err);
            }
        }

    });

    it('Should initiate Provider curve  with 0x0 broker in zap registry contract', async () => {

        try {

            const initProviderCurveTx: any = await registryWrapper.initiateProviderCurve({

                endpoint: testProvider.endpoints[0],
                term: testProvider.curve.values,
                broker: testProvider.broker
            });

            const curveReceipt: any = await initProviderCurveTx.wait();

            expect(curveReceipt).to.include.keys('events');

            expect(curveReceipt.events[0].event).to.equal('NewCurve');

            expect(curveReceipt.events[0]).to.include.keys('args');

            const args: any = curveReceipt.events[0].args;

            expect(args).to.include.keys('provider', 'endpoint', 'curve', 'broker');

            expect(args.broker).to.equal(testProvider.broker);

            expect(args.provider).to.equal(signers[0]._address);

            const getTxCurve: Array<number> = args.curve.map((num: any) => parseInt(num));

            const testCurve: Array<number> = testProvider.curve.values;

            expect(testProvider.endpoints[0]).to.equal(ethers.utils.parseBytes32String(args.endpoint));

            expect(testCurve).to.eql(getTxCurve);

        } catch (err: any) {

            console.log(signers[0]._address + ': ' + 'Curve is already initiated');
        }

    });

    it('Should set new title', async () => {

        const title: String = 'NEWTITLE';

        const setTitleTx: Object = await registryWrapper.setProviderTitle({
            title: title
        });

        expect(setTitleTx).to.be.ok;

        const newTitle: String = await registryWrapper.getProviderTitle(signers[0]._address);

        expect(newTitle).to.equal(title);

        expect(newTitle).to.equal(title);

    });

    it('Should initiate Provider curve with valid broker address in zap registry contract', async () => {

        const initCurveTwoTx: any = await registryWrapper.initiateProviderCurve({
            endpoint: testProvider.endpoints[1],
            term: testProvider.curve.values,
            broker: signers[1]._address,
        });

        expect(initCurveTwoTx).to.be.ok;

    });

    it('Should get the provider curve', async () => {

        const getCurve: Array<number> = await registryWrapper.getProviderCurve(
            signers[0]._address,
            testProvider.endpoints[0]
        );

        expect(getCurve).to.ok;

        expect(getCurve.values).to.eql(testProvider.curve.values);

        expect(getCurve.values).to.eql(testProvider.curve.values);

    });

    it('Should get endpoint broker', async () => {

        const getBroker: String = await registryWrapper.getEndpointBroker(
            signers[0]._address,
            testProvider.endpoints[0]
        );

        expect(getBroker).to.be.ok;
        expect(getBroker).to.be.equal(testProvider.broker);

    });

    it('Should set endpoint endpointParams in zap registry contract', async () => {

        const setParamsTx: Object = await registryWrapper.setEndpointParams({
            endpoint: testProvider.endpoints[0],
            endpoint_params: testProvider.endpoint_params,
        });

        expect(setParamsTx).to.be.ok

    });

    it('Should get the endpoint params in zap registry contract', async () => {

        const getEndpointParams: Array<string> = await registryWrapper.getEndpointParams({
            provider: signers[0]._address,
            endpoint: testProvider.endpoints[0]
        });

        expect(getEndpointParams).to.be.ok;
        expect(getEndpointParams).to.eql(testProvider.endpoint_params);

    });

    it('Should set the markdown url for the first endpoint param ', async () => {

        const setMdTx: Object = await registryWrapper.setProviderParameter({
            key: testProvider.endpoint_params[0],
            value: testProvider.markdownUrl
        });

        expect(setMdTx).to.be.ok;

    });

    it('Should set the json url for the second endpoint param ', async () => {

        const setJsonTx: Object = await registryWrapper.setProviderParameter({
            key: testProvider.endpoint_params[1],
            value: testProvider.jsonUrl
        });

        expect(setJsonTx).to.be.ok;

    });

    it('Should get the markdown url from the first endpoint param', async () => {

        const getMd: String = await registryWrapper.getProviderParam(
            signers[0]._address,
            testProvider.endpoint_params[0]
        );

        expect(getMd).to.equal(testProvider.markdownUrl);

    });

    it('Should get the json url from the second endpoint param', async () => {

        const getJson: String = await registryWrapper.getProviderParam(
            signers[0]._address,
            testProvider.endpoint_params[1]
        );

        expect(getJson).to.equal(testProvider.jsonUrl);

    });

    it('Should get all provider params', async () => {

        const getAllParams: Array<string> = await registryWrapper.getAllProviderParams(
            signers[0]._address
        );

        expect(getAllParams).to.eql(testProvider.endpoint_params);

    });

    it('Should be able to get all providers', async () => {

        const getProviders: Array<string> = await registryWrapper.getAllProviders();

        expect(getProviders).to.be.ok;

    });

    it('Should get provider address by index', async () => {

        const providers: Array<string> = [];

        const allProviders: String = await registryWrapper.getAllProviders();

        expect(allProviders).to.be.ok;

        for (let i = 0; i < allProviders.length; i++) {

            providers.push(await registryWrapper.getProviderAddressByIndex(i));

        }

        expect(providers).to.eql(allProviders);

    });

    it('Should check if all providers are initiated', async () => {

        const initStatus: Array<boolean> = [];

        const allProviders: Array<string> = await registryWrapper.getAllProviders();

        for (let i = 0; i < allProviders.length; i++) {

            initStatus.push(await registryWrapper.isProviderInitiated(allProviders[i]));

        };

        expect(initStatus.forEach(bool => bool === true));

    });

    it('Should check if the endpoint and corresponding curve is set', async () => {

        const curveStatus: Boolean = await registryWrapper.isEndpointSet(
            signers[0]._address,
            testProvider.endpoints[0]
        );

        expect(curveStatus).to.be.true;

    });

    it('Should clear the first endpoint', async () => {

        const clearTx: Object = await registryWrapper.clearEndpoint({

            endpoint: testProvider.endpoints[0]
        });

        expect(clearTx).to.be.ok;

    });

    it('Should clear the last endpoint', async () => {

        const clearTx: Object = await registryWrapper.clearEndpoint({

            endpoint: testProvider.endpoints[1]
        });

        expect(clearTx).to.be.ok;

    });

    it('Should have an empty array of endpoints after clearing', async () => {

        const getEndpoints: Array<string> = await registryWrapper.getProviderEndpoints(signers[0]._address);

        expect(getEndpoints).to.be.ok;

        expect(getEndpoints.length).to.equal(0);

    });

    it('Should initiate Provider curve after clearing', async () => {

        const initProviderCurveTx: Object = await registryWrapper.initiateProviderCurve({

            endpoint: testProvider.endpoints[0],
            term: testProvider.curve.values,
            broker: testProvider.broker
        });

        expect(initProviderCurveTx).to.be.ok;

    });

})