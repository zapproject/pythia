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

const expect = require('chai').expect;

describe('Registry Test', () => {

    let hardhatHttpProvider: any;
    let hardhatAccounts: any;
    let signerOne: any

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
            .then((txId: object) => {
                expect(txId).to.be.a('object')
            })
            .catch((err: any) => {
                return err;
            })
    });

});
