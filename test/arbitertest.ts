const expect = require('chai')
    .use(require('chai-as-promised'))
    .use(require('chai-bignumber'))
    .expect;
import {BaseContract } from '../BaseContract/basecontract';
import {ZapArbiter } from '../Arbiter/arbiter';
const Web3 = require('web3');
import {join } from 'path';
import {bootstrap } from './utils/setup_test';

import {testZapProvider } from '../Utils/utils';
import {DEFAULT_GAS} from '../Utils/constants';

async function configureEnvironment(func:()=>void) {
    await func();
}

const {utf8ToHex, toBN } = require('web3-utils');

describe('Arbiter Test', () => {
    let accounts :Array<string> = [],
        deployedToken:any,
        deployedRegistry:any,
        deployedBondage:any,
        arbiterWrapper:any,
        testArtifacts:any,
        ganacheServer:any,
        web3:any;
    const buildDir = join(__dirname, 'contracts')
    const options = {
        artifactsDir: buildDir,
        networkId: Utils.Constants.ganacheServerOptions.network_id,
        networkProvider: Utils.Constants.ganacheProvider,
        coordinator: undefined
    };


    it('Should set env', async ()=> {
            ganacheServer = Utils.startGanacheServer();
            web3 = new Web3(Utils.Constants.ganacheProvider);
            accounts = await web3.eth.getAccounts();
            //delete require.cache[require.resolve('/contracts')];
            await Utils.migrateContracts(buildDir);
            testArtifacts = Utils.getArtifacts(buildDir);
            deployedBondage = new BaseContract(Object.assign(options, {artifactName: 'BONDAGE' }));
            deployedRegistry = new BaseContract(Object.assign(options, {artifactName: 'REGISTRY' }));
            deployedToken = new BaseContract(Object.assign(options, {artifactName: 'ZAP_TOKEN' }));
            await Utils.delay(3000);
    });

    after(function(){
        console.log('Done running Arbiter tests');
        process.exit();
    });

    it('Should have all pre conditions set up for dispatch to work', async () => {
        const res = await bootstrap(testZapProvider, accounts, deployedRegistry, deployedToken, deployedBondage);
        await expect(res).to.be.equal('done');
    });

    it('Should initiate zapArbiter wrapper', async function() {
        arbiterWrapper = new ZapArbiter(options);
        await Utils.delay(3000);
        expect(arbiterWrapper).to.be.ok;
    });

    it('2) Should initiate zapArbiter wrapper with coordinator', async function() {
        options.coordinator = arbiterWrapper.coordinator._address;
        arbiterWrapper = new ZapArbiter(Object.assign(options, {artifactName: 'ARBITER' }));
        await Utils.delay(3000);
        expect(arbiterWrapper).to.be.ok;
    });

    it('3) Should initiate subscription', async function() {
        const subscription = await arbiterWrapper.initiateSubscription({
            provider: accounts[0],
            endpoint: testZapProvider.endpoint,
            endpoint_params: testZapProvider.endpoint_params,
            blocks: 4,
            pubkey: Utils.Constants.testZapProvider.pubkey,
            from: accounts[2],
            gas: DEFAULT_GAS
        }, (err: any, txid: string) => expect(txid).to.be.a('string'));
    });
    it('4) Should get subscription', async()=>{
        const subscription = await arbiterWrapper.getSubscription({
            provider: accounts[0],
            subscriber: accounts[2],
            endpoint: Utils.Constants.testZapProvider.endpoint
        });
        expect(subscription).to.have.keys('0', '1', '2', 'dots', 'blockStart', 'preBlockEnd');
        expect(subscription.dots).to.equal('4');

    });
    it('5) Should allow unscubscription from provider', async ()=>{
        const unsubscription = await arbiterWrapper.endSubscriptionProvider({
            subscriber: accounts[2],
            endpoint: testZapProvider.endpoint,
            from: accounts[0]
        }, (err: any, txid: string) => expect(txid).to.be.a('string'));
        const event = unsubscription.events;

    });
    it('6) Should allow unsubscription from subscriber', async()=>{
        await arbiterWrapper.initiateSubscription({
            provider: accounts[0],
            endpoint: testZapProvider.endpoint,
            endpoint_params: testZapProvider.endpoint_params,
            blocks: 4,
            pubkey: testZapProvider.pubkey,
            from: accounts[2],
            gas: DEFAULT_GAS
        }, (err: any, txid: string) => expect(txid).to.be.a('string'));
        await arbiterWrapper.endSubscriptionSubscriber({
            provider: accounts[0],
            endpoint: testZapProvider.endpoint,
            from: accounts[2]
        }, (err: any, txid: string) => expect(txid).to.be.a('string'));
    });
});