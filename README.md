# Pythia

**Currently only supports Linux and OSX systems.**

On windows, you can use the cli-client using WSL2.

[This is a great guide for installing WSL2 on Windows 10+ by OMG!Ubuntu](https://www.omgubuntu.co.uk/how-to-install-wsl2-on-windows-10).

**If you're running on Linux or WSL2**, be sure that you have OpenCL installed using these commands before proceeding:
```bash
sudo apt update
sudo apt install ocl-icd-opencl-dev
```

## Installation

1) Install go-lang https://golang.org/doc/install
2) You should now be able to run the commands in the [execute section](#Execute) 


## Supported Commands
**Pythia Flags**
   - `--config` (path to your config file, default is `config.json` in same Pythia root directory)
   - `--logConfig` (location of logging config file; default path is Pythia root directory)

**Pythia Commands**
   - `mine` (indicates to run the miner)
   - `mine -r` (indicates to mine utilizing a remote server)
   - `dataserver` (indicates to run the dataServer (no mining))
   - `transfer` (AMOUNT) (TO ADDRESS) (indicates transfer, toAddress is BSC address and the amount is number of Tributes (eg. transfer 10 0xea... (this transfers 10 tokens)))
   - `approve` (AMOUNT) (TO ADDRESS) (amount to approve the to address to send this amount of tokens
   - `stake deposit` (indicates to deposit tokens in the contract)
   - `stake request` (indicates you wish to withdraw your stake)
   - `stake withdraw` (withdraws your stake, run 1 week after request)
   - `stake status` (shows your staking balance)
   - `balance` (shows your balance)

## Execute

### Configuring your miner
Before running the miner, edit your `config.json` file so that you can add your publicAddress, privateKey and **Contract Addresses**:

1) Open a text editor (nano, VSCode, vim, or your OS's **text editor**) and create a new document/file.
2) paste this template in the file:
```json
    {
        "zapTokenAddress": "0x09d8af358636d9bcc9a3e177b66eb30381a4b1a8",
        "contractAddress": "",
        "nodeURL": "https://data-seed-prebsc-1-s2.binance.org:8545",
        "vaultAddress": "",
        "publicAddress": "",
        "privateKey": "",
        "serverHost": "0.0.0.0",
        "serverPort": 5001,
        "ethClientTimeout": 3000,
        "trackerCycle": 100,
        "gasMultiplier": 1,
        "gasMax":30,
        "serverWhitelist": [
        ],
        "useGPU":false,
        "trackers": [
            "balance",
            "disputeStatus",
            "gas",
            "tokenBalance",
            "indexers",
            "newCurrentVariables",
            "currentVariables"
        ],
        "dbFile": "zapDB",
        "disputeTimeDelta": "600s"
    }
```

You can find an explanation of each field [here](#configjson).

3) Pay special attention to the `contractAddress` and `vaultAddress` fields, update them with the ones in [this file](contracts.md).

4) Be sure to also add your `publicAddress` and `privateKey` to the new file as well. **Remember, this public address is the BSC wallet address that should contain Testnet BNB and BSC ZAP**.

5) Add the same `publicAddress` to the `serverWhitelist` field like this
```json
{
    ...
    serverWhitelist: [
        YOUR_PUBLIC_ADDRESS_HERE,
    ],
    ...
}
```

6) Save this file in the **root of this project's folder** as **config.json**. Meaning, it should be in the same folder as [main.go](main.go).

**If you would like to test on a localhost BSC Testnet node**, be sure have ZapHardhat running. https://github.com/zapproject/hardhat-bsc/.

Then, replace the `nodeURL` in the config.json with `http://localhost:8545`.

### Single Miner/Thread

**Running the commands**
1) `./release_build.sh`
2) `./pythia [cmd]`

replace `[cmd]` with one of the [supported commands](#supported-commands) based on what you would like to do, for example:

```bash
./pythia mine
```

Will get your miner running.

In this setup, your client will act as a datasever, getting blockchain information, and a miner, solving problems and writing data to blockchain.

### Remote Mining/Multiple Miners
**This setup will allow you to run multiple miners on a single host address.**

The Zap oracle network enables users to have a single data server provide data for multiple miner clients.

This data server will get data on the latest challenge to mine and provide it to every miner client you want it to be connected to.

Here is an idea of how your setup can be (note that you can have more or less miner clients than this):

```
Client A: Data Server

Client B: Miner Client
Client C: Miner Client
Client D: Miner Client
Client E: Miner Client
Client F: Miner Client
```

Your data server, Client A, can have the same `config.json` [you configured in this step](#configuring-your-miner).

On Clients B to F, change the serverAddress in `config.json` to the IP address or host address of Client A. **This may not be the same as the `0.0.0.0` address set for Client A.**

On Clients B to F, change their serverPort in `config.json` to the serverPort set in Client A's `config.json`. So by default, you can set it to `5001`.

Next, go on Client A's `config.json` and update the `serverWhitelist` field with the `publicAddresses` of Clients B to F. **Remember, these public addresses are the BSC wallet addresses that should contain Testnet BNB and BSC ZAP**.

Here is an example of how your `serverWhitelist` should look like:
```json
{
    ...
    serverWhitelist: [
        PUBLIC_ADDRESS_B_HERE,
        PUBLIC_ADDRESS_C_HERE,
        PUBLIC_ADDRESS_D_HERE,
        PUBLIC_ADDRESS_E_HERE,
        PUBLIC_ADDRESS_F_HERE
    ],
    ...
}
```

Now, follow these steps to get your miners running.

1) Go on Client A and run the following commands `./release_build.sh` then `./pythia dataserver`
2) Go on Clients B to F and run `./release_build.sh` then `./pythia mine -r`

Your miners should now be running.

## Running Disputes
### from cli
4) Locate `TimeStamp: %!(EXTRA *big.Int=XXX)` in the terminal running the dataserver.
5) Copy the big.Int value `XXX`. (ex: For "TimeStamp: %!(EXTRA *big.Int=168473848)", we only need 168473848.)
6) Then run: ```./pythia dispute new 1 {TimeStamp value `XXX` copied from step above} 4```

**Please reference CLI commands by running `./pythia --help`**.

## Subgraph

https://thegraph.com/explorer/subgraph/acemasterjb/zapminer

This subgraph indexes events on the BSC Testnet.

## BSC Testnet Miner Contract Addresses
```
BSC TESTNET MINER CONTRACTS

ZapToken Address: 0x09d8AF358636D9BCC9a3e177B66EB30381a4b1a8

ZapGettersLibary Address: 0x57b23103867b11F64dB2Ade885B7c3655F6be3c8

ZapTransfer Address: 0xAd6f4151ef0c7D49fb40CDd6e7d3Ed8977543080

ZapDispute Address: 0xF23937335c0794ea9920aC06c0f7947407919718

ZapStake Address: 0x00cf02aB915A9026C2328B07E9E45A854CCb28fd

ZapLibrary Address: 0x5896c4024de2DA768433123d442d8Ff2C1Ec694B

Zap Address: 0x4646939E9336139b16213B06EaCbA0a53b999f94

ZapMaster Address: 0xe623305CC15792f4e3E4Cd8880B5B9D665976df5

Vault Address: 0xc91fE0A6546f599b2887F342F5074E33D4fA37eE
```

## Config.json
```
tokenAddress (required)     - address of the Zap Token Contract

contractAddress (required)  - address of Zap Contract

vaultAddress (required)     - address of Vault Contract

nodeURL (required)          - RPC URL to the network (most likely BSC or localhost).

publicAddress (required)    - public address for your miner (note, no 0x)

privateKey (required)       - private key for miner node to sign transaction, without prefix 0x

ethClientTimeout (required) - timeout for making requests from your node

trackerCycle (required)     - how often your database updates (in seconds)

trackers (required)         - which pieces of the database you update

dbFile (required)           - where you want to store your local database (if self-hosting)

serverHost (required)       - location to host server

serverWhitelist (required)  - whitelists which publicAddress can access the data server.
                                Please have at least one in the array - your own publicAddress. See example below.

fetchTimeout                - timeout for requesting data from an API

useGPU                      - turn GPU usage on or off (DEFAULT, PLEASE DO NOT UNSET)

requestDataInterval         - min frequency at which to request data at (in seconds, default 30)

gasMultiplier               - Multiplies the submitted gasPrice

gasMax                      - a max for the gas price in gwei (note: this max comes BEFORE the gas multiplier.
                                So a max gas cost of 10 gwei, can have gas prices up to 20 if gasMultiplier is 2)

heartbeat                   - an integer that controls how frequently the miner process should
                                report the hashrate (larger is less frequent, try 1000000 to start)

numProcessors               - an integer number of CPU cores/threads to use for mining.

disputeTimeDelta (required) - how far back to store values for min/max range - default 5 (in minutes)

disputeThreshold            - percentage of acceptable range outside min/max for dispute checking -
```

## ENV Variables
If you wish to use environmental variables instead of config.json for some or all of your config fields, use 

`SET`       - export VARIABLE_NAME=arg ie. export PRIVATE_KEY=123ABC123ABC123ABC123ABC123ABC  
`CHECK`     - echo $VARIABLE_NAME  
`DELETE`    - unset $VARIABLE_NAME  

## Testing
You have a few options when you want to test pythia
### Test Scripts
In the root of this project, run:

```bash
./runPkgTest.sh <PACKAGE_NAME>
```

This will run all tests, identified by `*_test.go`, in the given package. It will also give you coverage insights.
If instead you'd like to test a specific test in a package run:

```bash
./runTest.sh <TEST_NAME> <PACKAGE_NAME>
```
e.g. `./runTest.sh TestDataServer dataServer`

This will run the given `<TEST_NAME>` as long as it is inside a `*_test.go` file in the given `<PACKAGE_NAME>`.

### Go test from CLI
You could of course also `cd` into the package of choice and run the standard `go test [test flags...]` e.g.
```bash
cd pow
go test -v -cover  # Test with increased verbosity and give a coverage analysis
```

You can also run individual test just like when running the `./runTest.sh` script
```bash
cd pythia # while in the root of the project...
go test -v [test flags...] [PACKAGE_NAME] -run [TEST_NAME]
```

For more `go test` flags you can use, run this command in your terminal
```bash
go help testflag
```


## **Debug**

### Install
Delve is a debugger for the Go programming language. Follow the steps in this [repo](https://github.com/go-delve/delve) to install onto your machine.

*For debugging, we needed to take out the -s and -w flags in the ./release-build.sh script.*
1) Follow the steps above up until Execute - 1. Don't run the ```./pythia``` command.
2) Instead run ```dlv exec ./pythia [command]```. This allows the debugger to run the script.
3) Set breakpoint(s).
4) Enter ```continue``` command to allow the program to run till breakpoint.
5) Check variables using ```locals``` and/or ```print```.
6) Step In and Out of functions with ```step``` and ```stepout```.
7) Use ```restart``` to start over. Breakpoints will persist.
8) Use ```quit``` to exit debug mode.


### **Basic Debug Commands**
[continue](#continue) | Run until breakpoint or program termination.  
[break](#break) | Sets a breakpoint. (EX. ```break tracker/index.go:39``` That will set a breakpoint on line 39 in tracker.index.go file.)  
[breakpoints](#breakpoints) | Print out info for active breakpoints.  
[step](#step) | Single step through program.  
[stepout](#stepout) | Step out of the current function.  
[locals](#locals) | Print local variables.  
[print](#print) | Evaluate an expression.  
[restart](#restart) | Restart process.  


More commands here ```$GOPATH/src/github.com/go-delve/delve/tree/master/Documentation/cli/locspec.md``` or type ```help``` when in debug mode.