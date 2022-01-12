package test

import (
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/util"
)

var configJSON = `{
	"zapTokenAddress": "0x5fbdb2315678afecb367f032d93f642f64180aa3",
    "contractAddress": "0x0165878a594ca255338adfa4d48449f69242eb8f",
    "vaultAddress": "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853",
    "nodeURL": "http://127.0.0.1:8545/",
    "publicAddress": "70997970C51812dc3A010C7d01b50e0d17dc79C8",
    "privateKey": "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
    "serverHost": "localhost",
    "serverPort": 5001,
    "ethClientTimeout": 3000,
    "trackerCycle": 100,
    "gasMultiplier": 1,
    "gasMax":30,
    "serverWhitelist": ["0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", "0xcd3B766CCDd6AE721141F452C550Ca635964ce71", "0x2546BcD3c84621e976D8185a91A922aE77ECEc30", "0xbDA5747bFD65F08deb54cb465eB87D40e51B197E", "0xdD2FD4581271e230360230F9337D5c0430Bf44C0", "0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199"],
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
    "indexFolder": "..",
    "dbFile": "zapDB"
}
`
var configJSON0 = `{
	"zapTokenAddress": "0x5fbdb2315678afecb367f032d93f642f64180aa3",
    "contractAddress": "0x0165878a594ca255338adfa4d48449f69242eb8f",
    "vaultAddress": "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853",
    "nodeURL": "http://127.0.0.1:8545/",
    "publicAddress": "f39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
    "privateKey": "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
    "serverHost": "localhost",
    "serverPort": 5001,
    "ethClientTimeout": 3000,
    "trackerCycle": 100,
    "gasMultiplier": 1,
    "gasMax":30,
    "serverWhitelist": ["0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", "0xcd3B766CCDd6AE721141F452C550Ca635964ce71", "0x2546BcD3c84621e976D8185a91A922aE77ECEc30", "0xbDA5747bFD65F08deb54cb465eB87D40e51B197E", "0xdD2FD4581271e230360230F9337D5c0430Bf44C0", "0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199"],
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
    "indexFolder": "..",
    "dbFile": "zapDB"
}
`

var configJSON1 = `{
    "zapTokenAddress": "0x5fbdb2315678afecb367f032d93f642f64180aa3",
    "contractAddress": "0x0165878a594ca255338adfa4d48449f69242eb8f",
    "vaultAddress": "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853",
    "nodeURL": "http://127.0.0.1:8545/",
    "publicAddress": "70997970C51812dc3A010C7d01b50e0d17dc79C8",
    "privateKey": "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
    "serverHost": "localhost",
    "serverPort": 5001,
    "ethClientTimeout": 3000,
    "trackerCycle": 100,
    "gasMultiplier": 1,
    "gasMax":30,
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
    "indexFolder": "indexes",
    "dbFile": "zapDB"
}`

var configJSON2 = `{
    "zapTokenAddress": "0x5fbdb2315678afecb367f032d93f642f64180aa3",
    "contractAddress": "0x0165878a594ca255338adfa4d48449f69242eb8f",
    "vaultAddress": "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853",
    "nodeURL": "http://127.0.0.1:8545/",
    "publicAddress": "3C44CdDdB6a900fa2b585dd299e03d12FA4293BC",
    "privateKey": "5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a",
    "serverHost": "localhost",
    "serverPort": 5001,
    "ethClientTimeout": 3000,
    "trackerCycle": 100,
    "gasMultiplier": 1,
    "gasMax":30,
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
    "indexFolder": "indexes",
    "dbFile": "zapDB"
}`

var configJSON3 = `{
    "zapTokenAddress": "0x5fbdb2315678afecb367f032d93f642f64180aa3",
    "contractAddress": "0x0165878a594ca255338adfa4d48449f69242eb8f",
    "vaultAddress": "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853",
    "nodeURL": "http://127.0.0.1:8545/",
    "publicAddress": "90F79bf6EB2c4f870365E785982E1f101E93b906",
    "privateKey": "7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6",
    "serverHost": "localhost",
    "serverPort": 5001,
    "ethClientTimeout": 3000,
    "trackerCycle": 100,
    "gasMultiplier": 1,
    "gasMax":30,
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
    "indexFolder": "indexes",
    "dbFile": "zapDB"
}`

var configJSON4 = `{
    "zapTokenAddress": "0x5fbdb2315678afecb367f032d93f642f64180aa3",
    "contractAddress": "0x0165878a594ca255338adfa4d48449f69242eb8f",
    "vaultAddress": "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853",
    "nodeURL": "http://127.0.0.1:8545/",
    "publicAddress": "15d34AAf54267DB7D7c367839AAf71A00a2C6A65",
    "privateKey": "47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a",
    "serverHost": "localhost",
    "serverPort": 5001,
    "ethClientTimeout": 3000,
    "trackerCycle": 100,
    "gasMultiplier": 1,
    "gasMax":30,
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
    "indexFolder": "indexes",
    "dbFile": "zapDB"
}`

var configJSON5 = `{
    "zapTokenAddress": "0x5fbdb2315678afecb367f032d93f642f64180aa3",
    "contractAddress": "0x0165878a594ca255338adfa4d48449f69242eb8f",
    "vaultAddress": "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853",
    "nodeURL": "http://127.0.0.1:8545/",
    "publicAddress": "9965507D1a55bcC2695C58ba16FB37d819B0A4dc",
    "privateKey": "8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba",
    "serverHost": "localhost",
    "serverPort": 5001,
    "ethClientTimeout": 3000,
    "trackerCycle": 100,
    "gasMultiplier": 1,
    "gasMax":30,
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
    "indexFolder": "indexes",
    "dbFile": "zapDB"
}`

func setup() {
	config.ParseConfigBytes([]byte(configJSON))

	util.ParseLoggingConfig("./testConfig.json")

	buildContext()
}

func setupOwner() {
	config.ParseConfigBytes([]byte(configJSON0))

	util.ParseLoggingConfig("../testConfig.json")

	buildMinerContext(5)
}

func setupMiners() {
	config.ParseConfigBytes([]byte(configJSON1))
	util.ParseLoggingConfig("../testConfig.json")
	buildMinerContext(0)

	config.ParseConfigBytes([]byte(configJSON2))
	util.ParseLoggingConfig("../testConfig.json")
	buildMinerContext(1)

	config.ParseConfigBytes([]byte(configJSON3))
	util.ParseLoggingConfig("../testConfig.json")
	buildMinerContext(2)

	config.ParseConfigBytes([]byte(configJSON4))
	util.ParseLoggingConfig("../testConfig.json")
	buildMinerContext(3)

	config.ParseConfigBytes([]byte(configJSON5))
	util.ParseLoggingConfig("../testConfig.json")
	buildMinerContext(4)
}
