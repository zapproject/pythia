import { Curve } from '../Curve/curve';
const Web3 = require('web3');
import { join, dirname } from 'path';
Web3.providers.HttpProvider.prototype.sendAsync = Web3.providers.HttpProvider.prototype.send;
Web3.providers.WebsocketProvider.prototype.sendAsync = Web3.providers.WebsocketProvider.prototype.send;
const zapContractDirName = dirname(require.resolve('zap_contracts/truffle-config.js'));
const migrationDir = join(zapContractDirName, 'migrations');
const contractsDir = join(zapContractDirName, 'contracts');
const workingDir = zapContractDirName;

// y = 2x + x^2 from [1, 100]
export const TEST_CURVE = [3, 0, 2, 1, 1000000000000000000];

export const migrate = require('@truffle/core/lib/commands/migrate.js');
/**
 * @ignore
 * Local ganache server options
 * @type {{hostname: string; network_id: number; port: number; total_accounts: number; ws: boolean; gas: number; gasPrice: number; network: string}}
 */
export const ganacheServerOptions = {
    hostname: 'localhost',
    network_id: 5777,
    port: 7545,
    total_accounts: 10,
    ws: true,
    quiet: true,
    gas: 6700000,
    gasPrice: 20000000,
    network: 'ganache-gui',
    keepAliveTimeout: 1000,
    mnemonic: 'solid giraffe crowd become skin deliver screen receive balcony ask manual current',
    solc: '0.4.25',
    compilers: {
        solc: {
            version: '0.4.25'
        }
    }
};
/**
 *@ignore
 * @type {{logger: Console; contracts_build_directory: string; contracts_directory: string; working_directory: string; migrations_directory: string}}
 */
export const buildOptions = {
    logger: console,
    contracts_build_directory: __dirname,
    contracts_directory: contractsDir,
    working_directory: workingDir,
    migrations_directory: migrationDir

};

/**
 * @ignore
 * Local test Zap provider information
 * @type {{pubkey: number; title: string; endpoint_params: string[]; endpoint: string; query: string; curve: Curve, broker: string}}
 */
export const testZapProvider: any = {
    pubkey: 111,
    title: 'testProvider',
    endpoint_params: ['p1', 'p2'],
    endpoint: 'testEndpoint',
    query: 'btcPrice',
    curve: new Curve(TEST_CURVE),
    broker: '0x0000000000000000000000000000000000000000'
};

export const EndpointBroker = 'EndpointBroker';
// export const ganacheProvider = new Web3.providers.HttpProvider('http://127.0.0.1:7545');
//export const ganacheProvider = new Web3.providers.WebsocketProvider('ws://127.0.0.1:7545');
export const ganacheProvider = 'http://127.0.0.1:7545';

/**
 *
 * @type {number}
 */
export const DEFAULT_GAS = 500000;
/**
 *
 * @type {number}
 */
export const GAS_PRICE = 40000000;
/**
 *
 * @type {string}
 */
export const NETWORK = 'ganache-gui';

export const NULL_ADDRESS = '0x0000000000000000000000000000000000000000';
