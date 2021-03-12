import { Curve } from '../Curve/curve';

export const TEST_CURVE = [3, 0, 2, 1, 10000];

export const testZapProvider: any = {
    pubkey: 111,
    title: 'testProvider',
    endpoint_params: ['p1', 'p2'],
    endpoints: ['testEndpoint', 'newEndpoint'],
    query: 'btcPrice',
    curve: new Curve(TEST_CURVE),
    broker: '0x0000000000000000000000000000000000000000',
    markdownUrl: 'https://raw.githubusercontent.com/mxstbr/markdown-test-file/master/TEST.md',
    jsonUrl: 'https://gateway.ipfs.io/ipfs/QmaWPP9HFvWZceV8en2kisWdwZtrTo8ZfamEzkTuFg3PFr'
};

export const HardhatServerOptions = {
    hostname: 'localhost',
    network_id: 31337,
    port: 8545,
};

export const delay = (ms: number) => new Promise(_ => setTimeout(_, ms));

export const HardhatProvider = 'http://127.0.0.1:8545';
