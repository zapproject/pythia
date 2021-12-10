# Pythia

This is a miner system that run on POW (Proof-Of-Work) built using Go. It is part of the Zap Protocol providing Oracle services for cryptocurrency prices and/or Forex exchange rates (against the US dollar).

To take part in mining, first configure your account in `config.json` with the deployed mining contracts and your wallet info. You can then stake 500,000 Zap Tokens and immediately start mining.

Pythia CLI is available on Linux/OSX systems and also available for Windows through WSL2. Please check out *[Prerequisites For Windows](#prerequisites-for-windows)*.


## Get Started
1. [Install Go](#install-go)
2. [Clone/Download Pythia Repository from Github](#repo)
2. [Configure Your Miner](#configuring-your-miner)
3. [Build Pythia](#build-pythia)
    - [Build for Window](#for-windows)
    - [Build for Linux/OSX](#for-linux-osx)
4. [Intro to CLI](#intro-to-cli)
5. [Pythia Commands](#supported-commands)
6. [Windows Prerequisites](#prerequisites-for-windows)


## Misc
1. [BSC Testnet Miner Contract Addresses](#bsc-testnet-miner-contract-addresses)
2. [Config.json Layout and definitions](#config.json)
3. [ENV Variables](#env-variables)
4. [Testnet BNB](#test-bnb)
2. [Testnet Zap](#test-zap)
3. [Running Multiple Miners](#running-multiple-miners)
4. [Running the miner on a local testnet](#local-testnet)
4. [Debug](#debug)

***
***

### Install Go

Download and install Go from Go's official site: https://golang.org/doc/install

*If you're using Windows, please check [Prerequisites For Windows](#prerequisites-for-windows), which has instructions for installing Go for WSL.*

***

### Repo

Clone Pythia by running the git command like below or navigate to the url and download to repository.

```bash
git clone https://github.com/zapproject/pythia
```

Once cloned, navigate into the pythia folder to start configuration.

***

### Configuring Your Miner
In order to start running your miner, you need to first create and configure a `config.json` file with your user information.


In the root folder, edit your `config.json` file by adding your publicAddress and privateKey from your wallet, and also the miner **Contract Addresses**. The fields already filled out are it's default values.

1) in the root folder, create a new file and call it `config.json`.
2) paste this template into `config.json`:
```json
    {
        "zapTokenAddress": "0x09d8af358636d9bcc9a3e177b66eb30381a4b1a8",
        "contractAddress": "",
        "vaultAddress": "",
        "nodeURL": "https://data-seed-prebsc-1-s2.binance.org:8545",
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

3) Pay special attention to the `contractAddress` and `vaultAddress` fields, update them with the ones in [this file](contracts.md).

4) Add your `publicAddress` and `privateKey` as well. **This public address is your wallet address that should have some Testnet BNB and BSC ZAP balance.**.

5) Also add your `publicAddress` to the `serverWhitelist` field like this
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

6) Save the file.

You can find an explanation of each field [here](#configjson).


**Note**  
*You have the choice of using environment variables instead of editing the `config.json`.
[You can see the available environment variables here](.env.example)*

***

### Build Pythia 
Running the build command will take a brief moment and once completed, there should be a Pythia executable in the root folder.

- #### For Windows: 
**Run the following command**

```bash
    ./release_build_win.sh`
```

**Note**
When trying to build Pythia on Windows and you get the error below:
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
3) save the file, and then run `release_build_win.sh` 

- #### For Linux OSX: 
**Run the following command**  
```bash
    `./release_build_linux.sh`
```

***

## Intro to CLI
After making your config file and building Pythia on your machine you will have all [pythia commands](#supported-commands) available to you.
Before you are even able to mine you will need to stake 500K BSC ZAP first per our requirements.
Aside from that you can begin getting used to the CLI by running a few commands like the following.
```bash
    ./pythia balance
```
This command will show the available balances from the account you set up your miner with.

You can also check your stake status at any time with this command.
```bash
    ./pythia stake status
```

After confirming that you have the correct amount of Zap needed to stake you can run the following command.
```bash
    ./pythia stake deposit
```
Any other tips and help can be found by reading through the readme. If you cannot find an answer to your question please join the Zap Community discord server and ask your questions in the validator-support channel!

***
## Supported Commands
The commands below are available to you once you have configured and built Pythia.
Besides `mine` and `dataserver` you can also run the following commands based on your needs:

**Pythia Commands**
   - `balance` (shows your ZAP and BNB balance)
   - `approve` (approve BSC ZAP tokens)
   - `transfer` (transfer BSC ZAP tokens)
   - `stake status` (shows whether or not your are staked)
   - `stake deposit` (this command will stake 500,000 BSC ZAP to the vault contract; **requirement for mining**)
   - `stake request` (you must request to withdraw your stake before withdrawing. **Once requested, there is a 7 day waiting period.**)
   - `stake withdraw` (withdraws your stake. can only run this command after running the stake request command)
   - `dispute vote` (vote on an active dispute)
   - `dispute new` (start a new dispute)
   - `dispute show` (show existing disputes)   
   - `mine` (start mining, will run the dataserver and miner together)
   - `mine -r` (indicates to mine utilizing a remote/independent dataserver)
   - `dataserver` (runs the remote/independent dataserver, it does not do any mining)

**Please reference CLI commands by running `./pythia --help`**.

**Pythia Flags**
   - `--config` (path to your config file, default is `config.json` in same Pythia root directory)
   - `--logConfig` (location of logging config file; default path is Pythia root directory)

***

## Prerequisites for Windows
Click here and [Install WSL2](https://www.omgubuntu.co.uk/how-to-install-wsl2-on-windows-10).

*If you already have WSL installed, upgrade to WSL2 by following the instructions near the bottom of that page.*
 
Then [install Go through WSL](https://sal.as/post/install-golan-on-wsl/).

Finally using WSL, download OpenCL using the commands below:
```bash
sudo apt update
sudo apt install ocl-icd-opencl-dev
```

***

## BSC Testnet Miner Contract Addresses
[See here to see all of the Pythia contract addresses](contracts.md)

***

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

***
 
## Test Tokens
### Test BNB
    BSC Testnet faucet: https://testnet.binance.org/faucet-smart

### Test Zap
    You can get test Zap from here http://faucet.zap.org/

***

### Running Multiple Miners
**This setup will allow you to run multiple miners on a single host address.**

The Zap oracle network enables users to have a single data server provide data for multiple miner clients.  
**Clients in this case refers to your Command Window or Terminal.**

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

Next, go on Client A's `config.json` and update the `serverWhitelist` field with the public addresses of Clients B to F. **Remember, these public addresses are the BSC wallet addresses that should contain Testnet BNB and BSC ZAP**.

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

1) Go on Client A and build Pythia based on your machine (Windows or Linux/OSX) then run `./pythia dataserver`
2) Go on Clients B to F and run `./release_build.sh` then `./pythia mine -r`

Your miners should now be running.

***

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

***

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
 
***

## Running Disputes
### from cli
1) Locate `TimeStamp: %!(EXTRA *big.Int=XXX)` in the terminal running the dataserver.
2) Copy the big.Int value `XXX`. (ex: For "TimeStamp: %!(EXTRA *big.Int=168473848)", we only need 168473848.)
3) Then run: ```./pythia dispute new 1 {TimeStamp value `XXX` copied from step above} 4```

***

## Local testnet
**If you would like to test on a localhost BSC Testnet node**, be sure have ZapHardhat running. https://github.com/zapproject/hardhat-bsc/.

Then, replace the `nodeURL` in the config.json with `http://localhost:8545`.

 ***

## **Debug the Miner**

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
