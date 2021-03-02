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
    let signer: any;

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

        signer = await hardhatHttpProvider.getSigner(hardhatAccounts[0])

    })

    it('Should be able to create registryWrapper', async () => {

        registryWrapper = new ZapRegistry(options);

        expect(registryWrapper).to.be.ok;

    });

    it('Should initiate provider in zap registry contract', async () => {

        const tx = await registryWrapper.initiateProvider({
            public_key: testProvider.pubkey,
            title: testProvider.title,
            from: signer._address

        })

    });

});
