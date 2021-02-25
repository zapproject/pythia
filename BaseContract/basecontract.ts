import { Artifacts } from '../Artifacts/artifacts';
console.log(Artifacts)

// import { BaseContractType } from '@pythiajs/types';
// import { Utils } from './utils';
// import { ethers } from 'ethers';
// const Web3 = require('web3');

// /**
//  * Parent Class to Dispatch, Bondage, Arbiter, Token, Registry classes
//  * Provides access to contract instances and the Web3 provider instance
//  */
// export class BaseContract {
//     provider: any;
//     web3: any;
//     contract: any;
//     networkId: number;
//     coordinator: any;
//     artifact: any;
//     name: string;
//     address: string | undefined;

//     /**
//      * Creates a contract class wrapper for a given contract.
//      * @constructor
//      * @param {BaseContractType} b. {artifactsDir,artifactName,networkId,networkProvider}
//      * @param {string | null} b.artifactsDir - Directory where contract ABIs are located
//      * @param {string} b.artifactName - Contract name for this contract object
//      * @param {number | null} b.networkId - Select which network the contract is located on (mainnet, testnet, private)
//      * @param {any | null} b.networkProvider - Ethereum network provider (e.g. Infura)
//      */
//     constructor({ artifactsDir, artifactName, networkId, networkProvider, coordinator, address, web3 }: BaseContractType) {
//         let coorArtifact;
//         this.name = artifactName;
//         try {
//             if (!artifactsDir) {
//                 this.artifact = Artifacts[artifactName];
//                 coorArtifact = Artifacts['ZAPCOORDINATOR'];
//             }
//             else {
//                 const artifacts: any = Utils.getArtifacts(artifactsDir);
//                 this.artifact = artifacts[artifactName];
//                 coorArtifact = artifacts['ZAPCOORDINATOR'];
//             }
//             // let currentProvider = networkProvider || new Web3.providers.HttpProvider("http://localhost:8545");
//             this.provider = web3 || new Web3(networkProvider || new Web3.providers.HttpProvider("https://cloudflare-eth.com"));
//             //network id default to mainnet
//             this.networkId = networkId || 1;
//             this.coordinator = new this.provider.eth.Contract(coorArtifact.abi, coordinator || coorArtifact.networks[this.networkId].address);
//             this.contract = undefined;
//             if (address) {
//                 this.address = address;
//             }
//             else {
//                 this.address = this.artifact.networks[this.networkId].address;
//             }
//             if (coordinator) {
//                 this.getContract()
//                     .catch(console.error);
//             }
//             else {
//                 this.contract = new this.provider.eth.Contract(this.artifact.abi, this.address);
//             }
//         } catch (err) {
//             throw err;
//         }
//     }

//     async getContract() {
//         const contractAddress = await this.coordinator.methods.getContract(this.name.toUpperCase()).call().valueOf();
//         this.contract = new this.provider.eth.Contract(this.artifact.abi, contractAddress);
//         return contractAddress;
//     }

//     /**
//      * Gets the address of the owner of this contract.
//      * @returns {Promise<string>} Returns a Promise that will eventually resolve into the address of this contract's owner.
//      */
//     async getContractOwner(): Promise<string> {
//         return await this.contract.methods.owner().call().valueOf();
//     }
// }
