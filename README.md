# Pythia

**Currently only supports Linux and OSX systems.**

On windows, you can use the cli-client using WSL2.

[This is a great guide for installing WSL2 on Windows 10+ by OMG!Ubuntu](https://www.omgubuntu.co.uk/how-to-install-wsl2-on-windows-10).

**If you're running on Linux or WSL2**, be sure that you have OpenCL installed using these commands before proceeding:
```bash
sudo apt update
sudo apt install ocl-icd-opencl-dev
```

**Note**
If you ever get an error like this
```
# pkg-config --cflags  -- OpenCL
Package 'OpenCL' has no Version: field
pkg-config: exit status 1
```
Then follow these steps:

1) run this command to open `/opt/intel/opencl/OpenCL.pc`
```
sudo nano /opt/intel/opencl/OpenCL.pc
```

2) replace the content of the file with this
```
prefix=/opt/intel/opencl
libdir=${prefix}
includedir=${prefix}/include

Name: OpenCL
Description: Open Computing Language generic Installable Client Driver loader
Version: 2.1
Libs: -L${libdir} -lOpenCL
Cflags: -I${includedir}
```
3) save the file, and then run `release_build.sh`

## Installation

1) Install go-lang https://golang.org/doc/install
2) You should now be able to run the commands in the [execute section](#Execute) 

## Execute

First of all, be sure to clone this repo

```bash
git clone https://github.com/zapproject/pythia
```

Then, create a file in the root of the cloned project and call it `config.json`.

### Configuring your miner
Before running the miner, edit your `config.json` file so that you can add your publicAddress, privateKey and **Contract Addresses**:

1) paste this template in the file:
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
        "localPort": 6363,
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

**Note**
**You have the choice of using environment variables instead of editing the `config.json`**.
[You can see the available environment variables here](.env.example)

2) Pay special attention to the `contractAddress` and `vaultAddress` fields, update them with the ones in [this file](contracts.md).

3) Again, be sure to also add your `publicAddress` and `privateKey` to the new file as well. **Remember, this public address is the BSC wallet address that should contain Testnet BNB and BSC ZAP**.

You can get test Zap from here http://faucet.zap.org/, BSC Testnet faucet: https://testnet.binance.org/faucet-smart.

4) Add your `publicAddress` to the `serverWhitelist` field like this
```json
{
    // ...
    "gasMax":30,
    "serverWhitelist": [
        "YOUR_PUBLIC_ADDRESS_HERE",
    ],
    "useGPU":false,
    // ...
}
```

5) Save `config.json`.

**If you would like to test on a localhost BSC Testnet node**, be sure have ZapHardhat running. https://github.com/zapproject/hardhat-bsc/.

Then, replace the `nodeURL` in the config.json with `http://localhost:8545`.

You can find an explanation of each field [here](#configjson).

### Run a Single Miner Client

**Run the following commands**
1) `./release_build.sh`
2) `./pythia mine`

Will get your miner running.

In this setup, your client will also act as a datasever, and a miner.

The data server is set on port 5001. There is also a forex data server on port 6363.

These are default ports, but you can change if you so wish.

### Remote Mining/Multiple Miners
**This setup will allow you to run multiple miners on a single host address.**

The Zap oracle network enables users to have a single data server provide data for multiple miner clients.

Here is an idea of how your setup can look like with this option:

```
Client A: Data Server

Client B: Miner Client
Client C: Miner Client
Client D: Miner Client
Client E: Miner Client
Client F: Miner Client
```

Your data server, Client A, can have the same `config.json` [you configured earlier](#configuring-your-miner).

On Clients B to F, change the `serverAddress` in your `config.json` to the IP address or host address of Client A.

On Clients B to F, change their serverPort in `config.json` to the serverPort set in Client A's `config.json`. So by default, you can set it to `5001`. Do the same for localPort, which is 6363 by default.

Next, go on Client A's `config.json` and update the `serverWhitelist` field with the public addresses of Clients B to F.
**Remember, these public addresses are the BSC wallet addresses that should contain Testnet BNB and BSC ZAP**.

Here is an example of how your `serverWhitelist` should look like:
```json
{
    "gasMax":30,
    "serverWhitelist": [
        "PUBLIC_ADDRESS_B_HERE",
        "PUBLIC_ADDRESS_C_HERE",
        "PUBLIC_ADDRESS_D_HERE",
        "PUBLIC_ADDRESS_E_HERE",
        "PUBLIC_ADDRESS_F_HERE"
    ],
    "useGPU":false,
}
```

Now, follow these steps to get your miners running.

1) Go on Client A and run the following commands `./release_build.sh` then `./pythia dataserver`
2) Go on Clients B to F and run `./release_build.sh` then `./pythia mine -r`

Your miners should now be running.

## Supported Commands
Besides `mine` and `dataserver` you can also run the following commands based on your needs:

**Pythia Commands**
   - `mine` (start mining, will run the dataserver and miner together)
   - `mine -r` (indicates to mine utilizing a remote/independent dataserver)
   - `dataserver` (runs the remote/independent dataserver, it does not do any mining)
   - `transfer [AMOUNT] [TO ADDRESS]` (transfer BSC ZAP, `TO ADDRESS` is a BSC address and the `AMOUNT` is number of ZAP (eg. `transfer 10 0xea...` (this transfers 10 ZAP BSC tokens)))
   - `approve [AMOUNT] [TO ADDRESS]` (`AMOUNT` of BSC ZAP to approve the `TO ADDRESS` to send this amount of tokens
   - `stake deposit` (this command will stake 500,000 BSC ZAP to the vault contract; **requirement for mining**)
   - `stake request` (you must request to withdraw your stake before withdrawing)
   - `stake withdraw` (withdraws your stake, can only be ran 1 week after running `stake request`)
   - `stake status` (shows whether or not your are staked)
   - `balance` (shows your ZAP and BNB balance)

**Pythia Flags**
   - `--config` (path to your config file, default is `config.json` in same Pythia root directory)
   - `--logConfig` (location of logging config file; default path is Pythia root directory)

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
[See here to see all of the Pythia contract addresses](contracts.md)

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

serverHost (required)       - ip address of dataserver

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