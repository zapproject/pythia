
# Pythia
<!-- Insert summary -->

Currently only supports Linux and OSX systems.

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
### Single Miner/Thread
(If running on local, have ZapHardhat running. https://github.com/zapproject/hardhat-bsc/)

Be sure to edit the `config.json` file to add your `publicAddress` and `privateKey` before running the following commands.

1) `release_build.sh`
2) `./pythia [cmd]`

### Remote Mining/Multiple Miners
(If running on local, have ZapHardhat running. https://github.com/zapproject/hardhat-bsc/)

Be sure to edit the configs found in `local_cfg` before running the follwing commands.

1) `release_build.sh`
2) `./start_local.sh`
3) Check the logs/ for failing miners. Manually run: `nohup ./pythia --config=local_cfgs/config{miner # 1-5}.json mine -r > logs/{miner # 1-5}.log &` for failing miners.

In order to run dispute commands:
### from cli
4) Locate `TimeStamp: %!(EXTRA *big.Int=XXX)` in the terminal running the dataserver.
5) Copy the big.Int value `XXX`. (ex: For "TimeStamp: %!(EXTRA *big.Int=168473848)", we only need 168473848.)
6) Then run: ```./pythia dispute new 1 {TimeStamp value `XXX` copied from step above} 4```

**Please reference CLI commands by using "./pythia" and the --help flag**

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


### Exmple config.json
```
{
    "zapTokenAddress": "0x4hj34jk53k4nk3434kh53k45hk3j45k3j4h5k43jl",
    "contractAddress": "0x3k4j53kj4h5k345kj345k3h45k3jh45k3h5k3b3k4",
    "vaultAddress": "0xk4j5nohgh43nio34hr4nr34ihjrl34jro34t34nl34t3lk",
    "nodeURL": "http://localhost:8545",
    "publicAddress": "0x123ABC123ABC123ABC123ABC123ABC123ABC",
    "privateKey": "123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC",
    "serverHost": "0.0.0.0",
    "serverPort": 5001,
    "ethClientTimeout": 3000,
    "trackerCycle": 100,
    "gasMultiplier": 1,
    "gasMax": 30,
    "serverWhitelist": [
        "0x123ABC123ABC123ABC123ABC123ABC123ABC",
    ],
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

### QT Build
In project directory run ```qt_build.sh```
In MINGW terminal run ```qtdeploy build desktop```