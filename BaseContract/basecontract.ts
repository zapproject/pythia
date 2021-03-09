// Imports the contract JSON files
import { Console } from 'node:console';
import { Artifacts } from '../Artifacts/artifacts';

import { BaseContractType } from '../Types/types';

import { Utils } from './utils';

const ethers = require('ethers');

export class BaseContract {
    provider: any;
    ethersjs: any;
    contract: any;
    networkId: number;
    coordinator: any;
    artifact: any;
    name: string;
    address: string | undefined;

    /**
     * Creates a contract class wrapper for a given contract.
     * @constructor
     * @param {BaseContractType} b. {artifactsDir,artifactName,networkId,networkProvider}
     * @param {string | null} b.artifactsDir - Directory where contract ABIs are located
     * @param {string} b.artifactName - Contract name for this contract object
     * @param {number | null} b.networkId - Select which network the contract is 
     * located on (mainnet, testnet, private)
     * @param {any | null} b.networkProvider - Ethereum network provider (e.g. Infura)
     */
    constructor(
        {
            artifactsDir,
            artifactName,
            networkId,
            networkProvider,
            coordinator,
            address,
            ethersjs
        }: BaseContractType) {

        let coorArtifact;

        this.name = artifactName;

        try {

            if (!artifactsDir) {
                this.artifact = Artifacts[artifactName];
                coorArtifact = Artifacts['ZAPCOORDINATOR'];
            }
            else {
                const artifacts: any = Utils.getArtifacts(artifactsDir);
                this.artifact = artifacts[artifactName];
                coorArtifact = artifacts['ZAPCOORDINATOR'];
            }

            this.provider = new ethers.providers.JsonRpcProvider('http://localhost:8545')

            this.networkId = networkId || 1;

            this.coordinator = new ethers.Contract(
                coordinator || coorArtifact.networks[this.networkId].address,
                coorArtifact.abi,
                this.provider
            );

            this.contract = undefined;

            if (address) {

                this.address = address;
            }
            else {
                this.address = this.artifact.networks[this.networkId].address;
            }
            if (coordinator) {
                this.getContract()
                    .catch(console.error);
            }
            else {

                this.contract = new ethers.Contract(this.address, this.artifact.abi);
            }
        } catch (err) {
            throw err;
        }
    }

    async getContract() {
        let hardhatHttpProvider = new ethers.providers.JsonRpcProvider('http://localhost:8545');

        let hardhatAccounts = await hardhatHttpProvider.listAccounts();

        let signer = await hardhatHttpProvider.getSigner(hardhatAccounts[0]);

        let contractAddress = this.artifact.networks[this.networkId];

        this.contract = new ethers.Contract(contractAddress.address, this.artifact.abi, signer);

        return contractAddress.address;
    }

    /**
     * Gets the address of the owner of this contract.
     * @returns {Promise<string>} Returns a Promise that will eventually resolve into the address of this contract's owner.
     */
    async getContractOwner(): Promise<string> {

        let hardhatHttpProvider = new ethers.providers.JsonRpcProvider('http://localhost:8545');

        let hardhatAccounts = await hardhatHttpProvider.listAccounts();

        let signer = await hardhatHttpProvider.getSigner(hardhatAccounts[0]);

        this.contract = this.contract.connect(signer)

        return await this.contract.owner();

        // return await this.contract.owner();
    }
}