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

    let accounts: Array<string> = [],
        HardhatServer: any,
        registryWrapper: any,
        ethersjs,
        testArtifacts,
        provider = testZapProvider,
        buildDir: string = join(__dirname, 'contracts'),
        options: any = {
            artifactsDir: '',
            networkId: HardhatServerOptions.network_id,
            networkProvider: HardhatProvider
        };

    it('Should be able to create registryWrapper', async () => {

        expect(new ZapRegistry(options)).to.be.ok

    })

});
