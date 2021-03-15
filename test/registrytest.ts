import { ZapRegistry } from '../Registry/registry';

import { ethers } from 'ethers';

import {
    testZapProvider,
    HardhatServerOptions,
    HardhatProvider
}
    from '../Utils/utils';
import { AnyARecord } from 'node:dns';

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

            expect(args.provider).to.equal(signers[0]._address);

            const title = await registryWrapper.getProviderTitle(signers[0]._address);

            expect(title).to.be.equal(testProvider.title);

            const pubkey = await registryWrapper.getProviderPublicKey(signers[0]._address);

            expect(pubkey).to.be.equal(testProvider.pubkey);

        } catch (err) {

            const initStatus = await registryWrapper.isProviderInitiated(signers[0]._address);

            if (initStatus === true) {

                console.log(signers[0]._address + ': ' + 'Is already initiated as a provider');
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

            expect(args.provider).to.equal(signers[0]._address);

            const getTxCurve = args.curve.map((num: any) => parseInt(num));

            const testCurve = testProvider.curve.values;

            expect(testProvider.endpoints[0]).to.equal(ethers.utils.parseBytes32String(args.endpoint));

            expect(testCurve).to.eql(getTxCurve);

        } catch (err: any) {

            console.log(signers[0]._address + ': ' + 'Curve is already initiated');
        }

    });

    it('Should set new title', async () => {

        const title = 'NEWTITLE';

        await registryWrapper.setProviderTitle({
            title: title
        })
            .then(async (setTitle: Object) => {

                expect(setTitle).to.be.ok;

                const newTitle = await registryWrapper.getProviderTitle(signers[0]._address);

                expect(newTitle).to.equal(title);
            })
            .catch((err: Object) => {

                return err;
            })
    });

    it('Should initiate Provider curve with valid broker address in zap registry contract', async () => {

        try {

            await registryWrapper.initiateProviderCurve({
                endpoint: testProvider.endpoints[1],
                term: testProvider.curve.values,
                broker: signers[1]._address,
            })
                .then((initCurveTwoTx: Object) => {

                    expect(initCurveTwoTx).to.be.ok;

                })
                .catch((err: Object) => {

                    return err;
                })

        } catch (err) {

            console.log(signers[1]._address + ': ' + 'Provider and Curve is already initiated');
        }

    });

    it('Should get the provider curve', async () => {

        await registryWrapper.getProviderCurve(
            signers[1]._address,
            testProvider.endpoints[0]
        )
            .then((getCurve: Array<number>) => {

                expect(getCurve).to.ok;

                expect(getCurve.values).to.eql(testProvider.curve.values);
            })
            .catch((err: Object) => {
                return err;
            })

    });

    it('Should get endpoint broker', async () => {

        await registryWrapper.getEndpointBroker(
            signers[1]._address,
            testProvider.endpoints[0]
        )
            .then((getBroker: String) => {

                expect(getBroker).to.be.ok;
                expect(getBroker).to.be.equal(testProvider.broker);
            })
            .catch((err: Object) => {

                return err;
            })
    });

    it('Should set endpoint endpointParams in zap registry contract', async () => {

        await registryWrapper.setEndpointParams({
            endpoint: testProvider.endpoints[0],
            endpoint_params: testProvider.endpoint_params,
        })
            .then((setParamsTx: Object) => {
                expect(setParamsTx).to.be.ok
            })
            .catch((err: Object) => {

                return err;
            })
    });

    it('Should get the endpoint params in zap registry contract', async () => {

        await registryWrapper.getEndpointParams({
            provider: signers[1]._address,
            endpoint: testProvider.endpoints[0]
        })
            .then((endpointParams: Array<string>) => {

                expect(endpointParams).to.be.ok;
                expect(endpointParams).to.eql(testProvider.endpoint_params);

            })
            .catch((err: Object) => {
                return err;
            })
    });

    it('Should set the markdown url for the first endpoint param ', async () => {

        await registryWrapper.setProviderParameter({
            key: testProvider.endpoint_params[0],
            value: testProvider.markdownUrl
        })
            .then((mdTx: Object) => {

                expect(mdTx).to.be.ok;
            })
            .catch((err: Object) => {
                return err;
            })
    });

    it('Should set the json url for the second endpoint param ', async () => {

        await registryWrapper.setProviderParameter({
            key: testProvider.endpoint_params[1],
            value: testProvider.jsonUrl
        })
            .then((jsonTx: Object) => {

                expect(jsonTx).to.be.ok;
            })
            .catch((err: Object) => {

                return err;
            })
    });

    it('Should get the markdown url from the first endpoint param', async () => {

        await registryWrapper.getProviderParam(
            signers[0]._address,
            testProvider.endpoint_params[0]
        )
            .then((markdownParam: String) => {
                expect(markdownParam).to.equal(testProvider.markdownUrl);
            })
            .catch((err: Object) => {
                return err;
            })
    });

    it('Should get the json url from the second endpoint param', async () => {

        await registryWrapper.getProviderParam(
            signers[0]._address,
            testProvider.endpoint_params[1]
        )
            .then((jsonParam: String) => {

                expect(jsonParam).to.equal(testProvider.jsonUrl);
            })
            .catch((err: Object) => {
                return err;
            })
    });

    it('Should get all provider params', async () => {

        await registryWrapper.getAllProviderParams(signers[0]._address)

            .then((getProviderParams: Array<string>) => {

                expect(getProviderParams).to.eql(testProvider.endpoint_params);

            })
            .catch((err: Object) => {

                return err;
            })
    });

    it('Should be able to get all providers', async () => {

        await registryWrapper.getAllProviders()

            .then((getProviders: Array<string>) => {

                expect(getProviders).to.be.ok;
            })
            .catch((err: Object) => {
                return err;
            })
    });

    it('Should get provider address by index', async () => {

        const providers = [];

        const allProviders = await registryWrapper.getAllProviders();

        expect(allProviders).to.be.ok;

        for (let i = 0; i < allProviders.length; i++) {

            providers.push(await registryWrapper.getProviderAddressByIndex(i));

        }

        expect(providers).to.eql(allProviders);

    });

    it('Should check if all providers are initiated', async () => {

        const initStatus = [];

        const allProviders = await registryWrapper.getAllProviders()

        for (let i = 0; i < allProviders.length; i++) {

            initStatus.push(await registryWrapper.isProviderInitiated(allProviders[i]));

        };

        expect(initStatus.forEach(bool => bool === true));

    });

    it('Should check if the endpoint and corresponding curve is set', async () => {

        await registryWrapper.isEndpointSet(
            signers[0]._address,
            testProvider.endpoints[0]
        )
            .then((unset: Boolean) => {
                expect(unset).to.be.true;
            })
            .catch((err: Object) => {
                return err;
            })

    });

    it('Should clear the first endpoint', async () => {

        await registryWrapper.clearEndpoint({

            endpoint: testProvider.endpoints[0]
        })
            .then((clearEndpointTx: Object) => {

                expect(clearEndpointTx).to.be.ok;
            })
            .catch((err: Object) => {

                return err;
            })
    });

    it('Should clear the last endpoint', async () => {

        await registryWrapper.clearEndpoint({

            endpoint: testProvider.endpoints[1]
        })
            .then((clearEndpointTx: Object) => {

                expect(clearEndpointTx).to.be.ok;

            })
            .catch((err: Object) => {

                return err;
            })
    });

    it('Should have an empty array of endpoints after clearing', async () => {

        await registryWrapper.getProviderEndpoints(signers[0]._address)

            .then((getEndpoints: Array<string>) => {

                expect(getEndpoints).to.be.ok;
                expect(getEndpoints.length).to.equal(0);
            })
            .catch((err: Object) => {

                return err;
            })
    })

})