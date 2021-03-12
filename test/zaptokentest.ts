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

        expect(zapTokenOwner).to.be.ok;
        expect(zapTokenWrapper).to.be.ok;

    });

    it('Should initiate wrapper with coordinator ', async () => {

        zapTokenWrapper = new ZapToken({
            networkId: HardhatServerOptions.network_id,
            networkProvider: HardhatProvider,
            coordinator: zapTokenWrapper.coordinator.address
        });

        expect(zapTokenWrapper).to.be.ok;

    });

    it('Should get zapToken owner', async () => {

        await zapTokenWrapper.getContractOwner()

            .then((owner: String) => {

                expect(owner).to.be.equal(signerOne._address);

            })
            .catch((err: Object) => {

                return err;
            })
    });

    it('Should get balance of zapToken from wrapper', async () => {

        await zapTokenWrapper.balanceOf(signerOne._address)

            .then((balance: Number) => {

                expect(balance).to.be.ok;
                expect(balance).to.be.gt(0);
            })
            .catch((err: Object) => {

                return err;
            })
    });

    it('Should update balance, and get updated balance of zap token', async () => {

        const beforeAllocation = await zapTokenWrapper.balanceOf(signerTwo._address);

        await zapTokenWrapper.allocate({
            to: signerTwo._address,
            amount: allocateAmount
        })
            .then(async (allocateTx: Object) => {

                const afterAllocation = await zapTokenWrapper.balanceOf(signerTwo._address);
                expect(beforeAllocation).to.be.ok;
                expect(afterAllocation).to.be.ok;
                expect(allocateTx).to.be.ok;
                expect(afterAllocation).to.be.equal(allocateAmount + beforeAllocation);

            })
            .catch((err: Object) => {

                return err;
            })
    });

});
