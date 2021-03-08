import { ZapToken } from '../ZapToken/zaptoken';

import { ethers } from 'ethers';

import {
    testZapProvider,
    HardhatServerOptions,
    HardhatProvider
}
    from '../Utils/utils';

const expect = require('chai').expect;

describe('ZapToken Test', () => {

    let accounts: Array<string> = [],
        HardhatServer: any,
        zapTokenWrapper: any,
        testProvider = testZapProvider,
        zapTokenOwner: string,
        options: any = {
            networkId: HardhatServerOptions.network_id,
            networkProvider: HardhatProvider
        };

    const allocateAmount = '1000';

    it('Should initiate wrapper', async () => {

        zapTokenWrapper = new ZapToken({
            networkId: options.network_id,
            networkProvider: options.networkProvider
        });

        console.log(zapTokenWrapper)

        expect(zapTokenWrapper).to.be.ok;

        // zapTokenOwner = await zapTokenWrapper.getContractOwner();
    });



});

