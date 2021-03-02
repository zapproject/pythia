// Imports the contract JSON files
import { Artifacts } from '../Artifacts/artifacts';

import { BaseContractType } from '../Types/types';

import { Utils } from './utils';

// import { ethers } from 'ethers';
const ethers = require('ethers')

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

            console.log(this.coordinator.db())

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
                this.contract = new this.provider.Contract(this.artifact.abi, this.address);
            }
        } catch (err) {
            throw err;
        }
    }

    async getContract() {
        const contractAddress = await this.coordinator.getContract(this.name.toUpperCase()).call().valueOf();
        this.contract = new this.provider.Contract(this.artifact.abi, contractAddress);
        return contractAddress;
    }

    /**
     * Gets the address of the owner of this contract.
     * @returns {Promise<string>} Returns a Promise that will eventually resolve into the address of this contract's owner.
     */
    async getContractOwner(): Promise<string> {
        return await this.contract.owner().call().valueOf();
    }
}

