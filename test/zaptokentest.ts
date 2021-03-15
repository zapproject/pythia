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

    let signers: Array<any> = [],
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

        for (let i = 0; i < hardhatAccounts.length; i++) {

            signers.push(await hardhatHttpProvider.getSigner(hardhatAccounts[i]));
        }

    });

    after(() => {

        console.log('Done running ZapToken tests');

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

                expect(owner).to.be.equal(signers[0]._address);

            })
            .catch((err: Object) => {

                return err;
            })
    });

    it('Should get balance of zapToken from wrapper', async () => {

        await zapTokenWrapper.balanceOf(signers[0]._address)

            .then((balance: Number) => {

                expect(balance).to.be.ok;
                expect(balance).to.be.gt(0);
            })
            .catch((err: Object) => {

                return err;
            })
    });

    it('Should update balance, and get updated balance of zap token', async () => {

        const beforeAllocation = await zapTokenWrapper.balanceOf(signers[1]._address);

        await zapTokenWrapper.allocate({
            to: signers[1]._address,
            amount: allocateAmount
        })
            .then(async (allocateTx: Object) => {

                const afterAllocation = await zapTokenWrapper.balanceOf(signers[1]._address);
                expect(beforeAllocation).to.be.ok;
                expect(afterAllocation).to.be.ok;
                expect(allocateTx).to.be.ok;
                expect(afterAllocation).to.be.equal(allocateAmount + beforeAllocation);

            })
            .catch((err: Object) => {

                return err;
            })
    });

    it('Should make transfer to another account', async () => {

        const beforeTransfer = await zapTokenWrapper.balanceOf(signers[1]._address);

        await zapTokenWrapper.send({
            to: signers[1]._address,
            amount: allocateAmount,
        })
            .then(async (transferTx: Object) => {

                const afterTransfer = await zapTokenWrapper.balanceOf(signers[1]._address);

                expect(transferTx).to.be.ok;
                expect(afterTransfer).to.be.ok;
                expect(afterTransfer).to.be.equal(allocateAmount + beforeTransfer);

            })
            .catch((err: Object) => {

                return err;
            })
    });

    it('Should approve to transfer from one to the another account', async () => {

        await zapTokenWrapper.approve({
            to: signers[1]._address,
            amount: allocateAmount,
        })
            .then((approveTx: Object) => {
                expect(approveTx).to.be.ok;
            })
            .catch((err: Object) => {

                return err;
            })
    });

});
