import BN from 'bn.js';
import { ContractAbi } from 'ethereum-types';
export type address = string;
export type txid = string;
export type NumType = string | number | BN;

export interface defaultTx {
    from?: address,
    gas?: string | number | BN,
    gasPrice?: string | number | BN
}

export interface listenEvent {
    filter?: Filter;
    callback: () => void;
}

export interface Filter {
    fromBlock?: number | BN,
    toBlock?: number | BN,
    provider?: address,
    subscriber?: address,
    terminator?: address,
    endpoint?: string,
    id?: number | string | BN
}

export interface Artifact {
    contract_name: string,
    abi: ContractAbi,
    networks: {
        [networkId: string]: {
            address: string
        }
    }
}



//== Base contract
export interface BaseContractType {
    artifactsDir?: string | undefined,
    artifactName: string,
    networkId?: number | undefined,
    networkProvider?: any | undefined,
    contract?: any,
    coordinator?: string,
    address?: string,
    ethersjs?: any
}

export interface NetworkProviderOptions {
    artifactsDir?: string | undefined,
    networkId?: number | undefined,
    networkProvider: any,
    coordinator?: string,
    address?: string,
    ethers?: any
}

export interface TransferType extends defaultTx {
    to: address,
    amount: BN | string | number
}


// CONSTANTS

export const DEFAULT_GAS = new BN(400000);
export const NULL_ADDRESS = '0x0000000000000000000000000000000000000000';


//######### ARBITER ################
export interface SubscriptionInit extends defaultTx {
    provider: address,
    endpoint: string,
    endpoint_params: Array<string>,
    blocks: NumType,
    pubkey: NumType
}

export interface SubscriptionEnd extends defaultTx {
    provider?: address,
    subscriber?: address,
    endpoint: string
}

export interface SubscriptionType {
    provider: address,
    subscriber: address,
    endpoint: string
}

export interface SubscriptionParams extends defaultTx {
    receiver: address,
    endpoint: string,
    params: Array<string>
}

export interface DataPurchaseEvent extends Filter {
    publicKey?: number | string | BN,
    amount?: number | string | BN,
    endpoint?: string,
    endpointParams?: string[]
}

export interface SubscriptionEndEvent extends Filter {
    terminator?: address
}

export interface ParamsPassedEvent {
    sender?: address,
    receiver?: address,
    endpoint?: string,
    params?: string
}


//############### BONDAGE ###############

export interface TokenBondType extends defaultTx {
    endpoint: string,
    dots: NumType
}
export interface BondType extends defaultTx {
    subscriber?: address,
    provider: address,
    endpoint: string,
    dots: NumType
}

export interface DelegateBondType extends BondType {
    subscriber: address
}
export interface UnbondType extends defaultTx {
    provider: address,
    endpoint: string,
    dots: NumType
}

export interface SubscribeType extends defaultTx {
    provider: address,
    endpoint: string,
    dots: NumType,
    endpointParams: string[]
}


export interface SubscriberHandler {
    handleResponse: () => void,
    handleUnsubscription?: () => void,
    handleSubscription?: () => void
}


export interface ApproveType extends defaultTx {
    provider: address,
    zapNum: string | number | BN
}


export interface BondArgs extends defaultTx {
    provider: address;
    endpoint: string;
    dots: NumType;
}

export interface UnbondArgs extends defaultTx {
    provider: address;
    endpoint: string;
    dots: NumType;
}

export interface DelegateBondArgs extends defaultTx {
    provider: address;
    endpoint: string;
    dots: NumType,
    subscriber: address;
}

export interface BondageArgs {
    subscriber?: address;
    provider: address;
    endpoint: string;
    dots?: NumType;
    zapNum?: NumType;
}

export interface CalcBondRateType {
    provider: address;
    endpoint: string;
    zapNum: NumType;
}

export interface BondFilter extends Filter {
    numDots?: NumType,
    numZap?: NumType
}

//#################### DISPATCH ##################



export interface ResponseArgs extends defaultTx {
    queryId: string,
    responseParams: Array<string | number>,
    dynamic: boolean
}

export interface cancelQuery extends defaultTx {
    queryId: NumType
}

export interface QueryArgs extends defaultTx {
    provider: address,
    endpoint: string,
    query: string,
    endpointParams: Array<string>,
    onchainProvider?: boolean,
    onchainSubscriber?: boolean
}


export interface ResponseArgs extends defaultTx {
    queryId: string,
    responseParams: Array<string | number>,
    dynamic: boolean
}
//=============Dispatch
export interface OffchainResponse {
    id?: number | string,
    subscriber?: address,
    provider?: address,
    response?: string[] | number[],
    response1?: string,
    response2?: string,
    response3?: string,
    response4?: string
}


//############################### PROVIDER #########################333


export interface InitProvider extends defaultTx {
    public_key: string,
    title: string
}

export interface InitCurve extends defaultTx {
    endpoint: string,
    term: number[],
    broker?: address
}

export interface InitDotTokenCurve extends InitCurve {
    symbol: string
}

export type UnsubscribeListen = {
    subscriber: address,
    terminator: address,
    fromBlock: number
};

export type ListenQuery = {
    queryId: string,
    subscriber: address,
    fromBlock: number
};

export interface Respond extends defaultTx {
    queryId: string,
    responseParams: Array<string | number>,
    dynamic: boolean
}



export interface SetProviderParams extends defaultTx {
    key: string,
    value: string
}

export interface SetProviderTitle extends defaultTx {
    title: string
}

export interface EndpointMethods extends defaultTx {
    endpoint: string
}


//################################## REGISTER #########################


export interface InitProvider extends defaultTx {
    public_key: string,
    title: string
}


export interface InitCurve extends defaultTx {
    endpoint: string,
    term: CurveType,
    broker?: address | undefined
}

export interface EndpointParams extends defaultTx {
    endpoint: string,
    endpoint_params: string[]
}

export interface SetProviderParams extends defaultTx {
    key: string,
    value: string
}

export interface SetProviderTitle extends defaultTx {
    from: address,
    title: string
}

export interface Endpoint extends defaultTx {
    endpoint: string
}


//#################################### SUBSCRIBER #########################33


export interface BondType extends defaultTx {
    subscriber?: address,
    provider: address,
    endpoint: string,
    dots: NumType
}

export interface DelegateBondType extends BondType {
    subscriber: address
}
export interface UnbondType extends defaultTx {
    provider: address,
    endpoint: string,
    dots: NumType
}

export interface SubscribeType extends defaultTx {
    provider: address,
    endpoint: string,
    dots: NumType,
    endpointParams: string[]
}


export interface SubscriberHandler {
    handleResponse: () => void,
    handleUnsubscription?: () => void,
    handleSubscription?: () => void
}



//########################## CURVE ######################

export type CurveType = number[];

export interface CurveTerm {
    fn: number;
    power: number;
    coef: number;
}

//###################### TOKEN DOT FACTORY #####################
export interface InitProvider extends defaultTx {
    public_key: string,
    title: string
}

export interface InitTokenCurve extends defaultTx {
    specifier: string,
    ticker: string,
    term: CurveType
}

export interface InitCurve extends defaultTx {
    endpoint: string,
    term: CurveType,
    broker?: address | undefined
}


export type NextEndpoint = {
    provider: address,
    endpoint: string
};

export interface EndpointParams extends defaultTx {
    endpoint: string,
    endpoint_params: string[]
}

export interface SetProviderParams extends defaultTx {
    key: string,
    value: string
}

export type TransactionCallback =
    (error: string | null,
        hash?: string | null) => void;
