import { ZapBondage } from '../Bondage/bondage';

import { join } from 'path';

import { ethers } from 'ethers';

import {
    testZapProvider,
    HardhatServerOptions,
    HardhatProvider
}
    from '../Utils/utils';
import { doesNotReject } from 'node:assert';
import { Utils } from '../BaseContract/utils';

const expect = require('chai').expect;

describe('Bondage Test', () => {

    let hardhatHttpProvider: any;
    let hardhatAccounts: any;
    let signerOne: any
    let signerTwo: any
    let signerThree: any

    let accounts: Array<string> = [],
        HardhatServer: any,
        bondageWrapper: any,
        testProvider = testZapProvider,
        options: any = {
            networkId: HardhatServerOptions.network_id,
            networkProvider: HardhatProvider
        };

    beforeEach(async () => {

        hardhatHttpProvider = new ethers.providers.JsonRpcProvider(options.networkProvider);

        hardhatAccounts = await hardhatHttpProvider.listAccounts();

        signerOne = await hardhatHttpProvider.getSigner(hardhatAccounts[0]);

        signerTwo = await hardhatHttpProvider.getSigner(hardhatAccounts[1]);

        signerThree = await hardhatHttpProvider.getSigner(hardhatAccounts[2]);

    });

    after(() => {

        console.log('Done running Bondage tests');
    });

    it('Should be able to create bondageWrapper', async () => {

        bondageWrapper = new ZapBondage(options);

        bondageWrapper = bondageWrapper.contract.connect(signerOne)

        expect(bondageWrapper).to.be.ok;

    });

    it('Should initiate bondageWrapper through Coordinator address', async () => {

        options.ZapCoordinator = bondageWrapper.ZapCoordinator._address;
        bondageWrapper = new ZapBondage(options);
        //await Utils.delay(3000);

        expect(bondageWrapper).to.be.ok;
    })

})