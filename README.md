# Pythia

# Development

- These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.
- Currently have to run the Zap-Hardhat node separately from Pythia until integration

## Installing

### Pythia
1. Clone the Pythia repository to your local machine
```git clone https://github.com/zapproject/pythia.git```

2. After the Pythia repository is cloned 
```npm install``` to build the dependencies inside Pythia


### Zap-Hardhat
1. Clone the Zap-Hardhat repository to your local machine
```git clone https://github.com/zapproject/zap-hardhat.git```

2. Install dependencies
 ```npm install``` 
``` npm run build```

3. Build Contracts and Generate Typechain Typings
```npm run compile```

## Run Hardhat Node & Deploy Contracts
1. In a separate terminal or window of Zap-Hardhat run the shell command by preference ```sh start.sh``` or ```./start.sh```

## Running Tests
1. In a separate terminal or window of Pythia run ```npm test```

## Contracts

### Registry
- Enabling data providers (oracles) to register their endpoints and bonding curves. Furthermore, this contract enables data subscribers to discover oracles and receive configuration data (such as titles, endpoints, and parameters) from the Registry smart contract.

### ZapToken
- Enabling a Zap holder to check their balance, transfer Zap to other accounts, and delegate transfers (approvals) to other accounts and smart contracts.

### Subscriber
- Enabling queries to be sent to oracles and responses to be received by subscribers

### Dispatch
- Enabling queries to be sent to oracles and responses to be received by subscribers.

### Arbiter
- Enabling subscribers to listen to subscription events
