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

    let hardhatHttpProvider: any;
    let hardhatAccounts: any;
    let signerOne: any
    let signerTwo: any
    let signerThree: any


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

    beforeEach(async () => {

        hardhatHttpProvider = new ethers.providers.JsonRpcProvider(options.networkProvider);

        hardhatAccounts = await hardhatHttpProvider.listAccounts();

        signerOne = await hardhatHttpProvider.getSigner(hardhatAccounts[0]);

        signerTwo = await hardhatHttpProvider.getSigner(hardhatAccounts[1]);

        signerThree = await hardhatHttpProvider.getSigner(hardhatAccounts[2]);

    });

    it('Should initiate wrapper', async () => {

        zapTokenWrapper = new ZapToken(options)

        zapTokenOwner = await zapTokenWrapper.getContractOwner();

        expect(zapTokenWrapper).to.be.ok;

    });

    it('Should initiate wrapper with coordinator ', async () => {

        zapTokenWrapper = new ZapToken({
            networkId: HardhatServerOptions.network_id,
            networkProvider: HardhatProvider,
            coordinator: zapTokenWrapper.coordinator.address
        });

        zapTokenOwner = await zapTokenWrapper.getContractOwner();

        expect(zapTokenWrapper).to.be.ok;

    });

    it('Should get zapToken owner', async () => {

        const owner = await zapTokenWrapper.getContractOwner();

        expect(owner).to.be.equal(signerOne._address);

    });


});

