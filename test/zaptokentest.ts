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

    const allocateAmount: any = 1000;

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

        const owner: String = await zapTokenWrapper.getContractOwner();

        expect(owner).to.be.equal(signers[0]._address);

    });

    it('Should get balance of zapToken from wrapper', async () => {

        const getBalance: Number = await zapTokenWrapper.balanceOf(signers[0]._address)

        expect(getBalance).to.be.ok;
        expect(getBalance).to.be.gt(0);

    });

    it('Should update balance, and get updated balance of zap token', async () => {

        const beforeAllocation: any = await zapTokenWrapper.balanceOf(signers[1]._address);

        const allocateTx: Object = await zapTokenWrapper.allocate({
            to: signers[1]._address,
            amount: allocateAmount
        });

        const afterAllocation: any = await zapTokenWrapper.balanceOf(signers[1]._address);

        expect(beforeAllocation).to.be.ok;

        expect(afterAllocation).to.be.ok;

        expect(allocateTx).to.be.ok;

        expect(afterAllocation).to.be.equal(beforeAllocation + allocateAmount);

    });

    it('Should make transfer to another account', async () => {

        const beforeTransfer: any = await zapTokenWrapper.balanceOf(signers[1]._address);

        const transferTx: Object = await zapTokenWrapper.send({
            to: signers[1]._address,
            amount: allocateAmount,
        })

        const afterTransfer: any = await zapTokenWrapper.balanceOf(signers[1]._address);

        expect(transferTx).to.be.ok;

        expect(afterTransfer).to.be.ok;

        expect(afterTransfer).to.be.equal(allocateAmount + beforeTransfer);

    });

    it('Should approve to transfer from one to the another account', async () => {

        const approveTx: Object = await zapTokenWrapper.approve({
            to: signers[1]._address,
            amount: allocateAmount,
        })

        expect(approveTx).to.be.ok;

    });

});
