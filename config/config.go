package config

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"strings"
	"time"
)

//unfortunate hack to enable json parsing of human readable time strings
//see https://github.com/golang/go/issues/10275
//code from https://stackoverflow.com/questions/48050945/how-to-unmarshal-json-into-durations
type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value * float64(time.Second))
		return nil
	case string:
		dur, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		d.Duration = dur
		return nil
	default:
		return fmt.Errorf("invalid duration")
	}
}

type GPUConfig struct {
	//number of threads in a workgroup
	GroupSize int `json:"groupSize"`
	//total number of threads
	Groups int `json:"groups"`
	//number of iterations within a thread
	Count uint32 `json:"count"`

	Disabled bool `json:"disabled"`
}

//Config holds global config info derived from config.json
type Config struct {
	TokenAddress                 string                `json:"zapTokenAddress"`
	ContractAddress              string                `json:"contractAddress"`
	VaultAddress                 string                `json:"vaultAddress"`
	NodeURL                      string                `json:"nodeURL"`
	DatabaseURL                  string                `json:"databaseURL"`
	PublicAddress                string                `json:"publicAddress"`
	EthClientTimeout             uint                  `json:"ethClientTimeout"`
	TrackerSleepCycle            Duration              `json:"trackerCycle"`
	Trackers                     []string              `json:"trackers"`
	DBFile                       string                `json:"dbFile"`
	ServerHost                   string                `json:"serverHost"`
	ServerPort                   uint                  `json:"serverPort"`
	FetchTimeout                 Duration              `json:"fetchTimeout"`
	RequestData                  uint                  `json:"requestData"`
	MinConfidence                float64               `json:"minConfidence"`
	RequestDataInterval          Duration              `json:"requestDataInterval"`
	RequestTips                  int64                 `json:"requestTips"`
	MiningInterruptCheckInterval Duration              `json:"miningInterruptCheckInterval"`
	GasMultiplier                float32               `json:"gasMultiplier"`
	GasMax                       uint                  `json:"gasMax"`
	NumProcessors                int                   `json:"numProcessors"`
	Heartbeat                    Duration              `json:"heartbeat"`
	ServerWhitelist              []string              `json:"serverWhitelist"`
	GPUConfig                    map[string]*GPUConfig `json:"gpuConfig"`
	EnablePoolWorker             bool                  `json:"enablePoolWorker"`
	Worker                       string                `json:"worker"`
	Password                     string                `json:"password"`
	PoolURL                      string                `json:"poolURL"`
	IndexFolder                  string                `json:"indexFolder"`
	DisputeTimeDelta             Duration              `json:"disputeTimeDelta"` //ignore data further than this away from the value we are checking
	DisputeThreshold             float64               `json:"disputeThreshold"` //maximum allowed relative difference between observed and submitted value
	UseGPU                       bool                  `json:"useGPU"`

	//config parameters excluded from the json config file
	PrivateKey string `json:"privateKey"`
}

const defaultTimeout = 30 * time.Second //30 second fetch timeout

const defaultRequestInterval = 300 * time.Second //300 seconds between data requests (0-value tipping)
const defaultMiningInterrupt = 15 * time.Second  //every 15 seconds, check for new challenges that could interrupt current mining
const defaultCores = 2

const defaultHeartbeat = 15 * time.Second //check miner speed every 10 ^ 8 cycles
var (
	config *Config
)

const defaultMaxParallelPSR = 4

const defaultTrackerInterval = 30 * time.Second

const DefaultMaxCheckTimeDelta = 5 * time.Minute

//threshold, a percentage of the expected value
const DefaultDisputeThreshold = 0.01

const ZapTokenAddress = "ZAP_TOKEN_ADDRESS"

const ContractAddress = "CONTRACT_ADDRESS"

const VaultAddress = "VAULT_ADDRESS"

const NodeURLEnvName = "NODE_URL"

const PublicAddress = "PUBLIC_ADDRESS"

const PrivateKeyEnvName = "ETH_PRIVATE_KEY"

const ServerHost = "SERVER_HOST"

const ServerPort = "SERVER_PORT"

const EthClientTimeout = "ETH_CLIENT_TIMEOUT"

const TrackerCycle = "TRACKER_CYCLE"

const GasMultiplier = "GAS_MULTIPLIER"

const GasMax = "GAS_MAX"

const UseGPU = "USE_GPU"

const DBFile = "DB_FILE"

const DisputeTimeDelta = "DISPUTE_TIME_DELTA"

const ServerWhiteList = "SERVER_WHITELIST"

const Trackers = "TRACKERS"

const RequestData = "REQUEST_DATA"

const RequestTips = "REQUEST_TIPS"

const NumProcessors = "NUM_PROCESSORS"

//ParseConfig and set a shared config entry
func ParseConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to open config file %s: %v", path, err)
	}
	return ParseConfigBytes(data)
}

func ParseConfigBytes(data []byte) error {

	//parse the json
	err := json.Unmarshal(data, &config)
	if err != nil {
		return fmt.Errorf("failed to parse json: %s", err.Error())
	}

	parsedNilDuration, _ := time.ParseDuration("0")

	nilDuration := Duration{parsedNilDuration}

	ZapTokenEnv := os.Getenv(ZapTokenAddress)

	if ZapTokenEnv == "" {
		if config.TokenAddress == "" {

			return fmt.Errorf(
				"missing the zapToken contract address environment variable'%s'", ZapTokenAddress)

		}

	} else {

		config.TokenAddress = ZapTokenEnv
	}

	ContractAddressEnv := os.Getenv(ContractAddress)

	if ContractAddressEnv == "" {

		if config.ContractAddress == "" {

			return fmt.Errorf("missing the contract address environment variable '%s'",
				ContractAddress)

		}

	} else {

		config.ContractAddress = ContractAddressEnv
	}

	VaultAddressEnv := os.Getenv(VaultAddress)

	if VaultAddressEnv == "" {

		if config.VaultAddress == "" {

			return fmt.Errorf("missing the vault address enviornment variable '%s'", VaultAddress)
		}
	} else {

		config.VaultAddress = VaultAddressEnv
	}

	NodeUrlEnv := os.Getenv(NodeURLEnvName)

	if NodeUrlEnv == "" {

		if config.NodeURL == "" {

			return fmt.Errorf("missing the node url environment variable '%s'",
				NodeURLEnvName)

		}

	} else {

		config.NodeURL = NodeUrlEnv
	}

	PublicAddressEnv := os.Getenv(PublicAddress)

	if PublicAddressEnv == "" {

		if config.PublicAddress == "" {
			return fmt.Errorf(
				"missing the public address environment variable '%s'", PublicAddress)

		}
	} else {

		config.PublicAddress = PublicAddressEnv
	}

	PrivateKeyEnv := os.Getenv(PrivateKeyEnvName)

	if PrivateKeyEnv == "" {

		if config.PrivateKey == "" {
			return fmt.Errorf("missing the private key environment variable '%s'", PrivateKeyEnvName)
		}

	} else {

		config.PrivateKey = PrivateKeyEnv
	}

	ServerHostEnv := os.Getenv(ServerHost)

	if ServerHostEnv == "" {

		if config.ServerHost == "" {

			return fmt.Errorf(
				"missing the server host environment variable '%s'", ServerHost)
		}
	} else {

		config.ServerHost = ServerHostEnv
	}

	ServerPortEnv := os.Getenv(ServerPort)

	if ServerPortEnv == "" {

		if config.ServerPort == 0 {

			return fmt.Errorf(
				"missing the server port environment variable '%s'", ServerPort)
		}

	} else {
		parsedUint, err := strconv.ParseUint(ServerPortEnv, 10, 32)

		if err != nil {

			return fmt.Errorf("error in parsing ServerPort from os.env: %s", err)
		}
		config.ServerPort = uint(parsedUint)
	}

	EthClientTimeoutEnv := os.Getenv(EthClientTimeout)

	if EthClientTimeoutEnv == "" {

		if config.EthClientTimeout == 0 {

			return fmt.Errorf(
				"missing the eth client timeout environment variable '%s'", EthClientTimeout)
		}

	} else {
		parsedUint, err := strconv.ParseUint(EthClientTimeoutEnv, 10, 32)
		if err != nil {
			return fmt.Errorf("error in parsing EthClientTimeout from os.env: %s", err)
		}
		config.EthClientTimeout = uint(parsedUint)
	}

	GasMultiplierEnv := os.Getenv(GasMultiplier)

	if GasMultiplierEnv == "" {

		if config.GasMultiplier == 0 {

			return fmt.Errorf(
				"missing the gas multiplier environment variable '%s'", GasMultiplier)
		}
	} else {

		parsedFloat, err := strconv.ParseFloat(GasMultiplierEnv, 32)

		if err != nil {
			return fmt.Errorf("error in parsing GasMultiplier from os.env: %s", err)
		}
		config.GasMultiplier = float32(parsedFloat)
	}

	GasMaxEnv := os.Getenv(GasMax)

	if GasMaxEnv == "" {

		if config.GasMax == 0 {

			return fmt.Errorf(
				"missing the gas max environment variable '%s'", GasMax)
		}
	} else {

		parsedUint, err := strconv.ParseUint(GasMaxEnv, 10, 32)

		if err != nil {

			return fmt.Errorf("error in parsing GasMax from os.env: %s", err)
		}
		config.GasMax = uint(parsedUint)
	}

	ServerWhitelistEnv := os.Getenv(ServerWhiteList)

	if ServerWhitelistEnv == "" {

		if len(config.ServerWhitelist) == 0 {

			return fmt.Errorf(
				"missing the server whitelist environment variable '%s'", ServerWhiteList)
		}
	} else {

		config.ServerWhitelist = strings.Split(ServerWhitelistEnv, ",")

	}

	RequestDataEnv := os.Getenv(RequestData)

	if RequestDataEnv == "" {

		if config.RequestData == 0 {

			fmt.Println("RequestData flag not set, continuing without tipping for a set interval...")
		}

	} else {
		parsedUint, err := strconv.ParseUint(RequestDataEnv, 10, 32)
		if err != nil {
			return fmt.Errorf("error in parsing RequestData from os.env: %s", err)
		}
		config.RequestData = uint(parsedUint)
	}

	RequestTipsEnv := os.Getenv(RequestTips)

	if RequestTipsEnv == "" {
		if config.RequestTips == 0 && config.RequestData != 0 {
			fmt.Println("Tip amount was not set, defaulting to 1 Zap per tip")
			config.RequestTips = 1
		} else if config.RequestData == 0 {
			fmt.Println("Not going to set a tip amount because this miner is not configured to tip requests")
		}
	} else {
		parsedUint, err := strconv.ParseUint(RequestTipsEnv, 10, 32)
		if err != nil {
			return fmt.Errorf("error parsing RequestTips from os.env: %s", err)
		}
		config.RequestTips = int64(parsedUint)
	}

	UseGpuEnv := os.Getenv(UseGPU)

	if UseGpuEnv == "" {

		config.UseGPU = false

	} else {

		useGpu, _ := strconv.ParseBool(UseGpuEnv)

		config.UseGPU = useGpu
	}

	TrackersEnv := os.Getenv(Trackers)

	if TrackersEnv == "" {

		if len(config.Trackers) == 0 {

			return fmt.Errorf(
				"missing the trackers environment variable '%s'", TrackersEnv)
		}
	} else {

		config.Trackers = strings.Split(TrackersEnv, ",")

	}

	DBFileEnv := os.Getenv(DBFile)

	if DBFileEnv == "" {

		if config.DBFile == "" {

			return fmt.Errorf(
				"missing the db file environment variable '%s'", DBFileEnv)
		}
	} else {

		config.DBFile = DBFileEnv
	}

	DisputeTimeDeltaEnv := os.Getenv(DisputeTimeDelta)

	if DisputeTimeDeltaEnv == "" {

		if config.DisputeTimeDelta == nilDuration {

			return fmt.Errorf(
				"missing the dispute time delta environment variable '%s'", DisputeTimeDeltaEnv)

		}
	} else {

		parsedDuration, err := time.ParseDuration(DisputeTimeDeltaEnv)
		if err != nil {
			return fmt.Errorf("failed to parse DisputeTimeDelta from os.env: %s", err)
		}
		config.DisputeTimeDelta = Duration{parsedDuration}
	}

	TrackerCycleEnv := os.Getenv(TrackerCycle)

	if TrackerCycleEnv == "" {
		if config.TrackerSleepCycle == nilDuration {

			return fmt.Errorf(
				"missing the tracker cycle environment variable '%s'", TrackerCycleEnv)

		}
	} else {

		parsedDuration, err := time.ParseDuration(TrackerCycleEnv)
		if err != nil {
			return fmt.Errorf("failed to parse TrackerCycle from os.env: %s", err)
		}
		config.TrackerSleepCycle = Duration{parsedDuration}
	}

	if config.MinConfidence == 0 {
		config.MinConfidence = 0.2
	}

	if config.FetchTimeout.Seconds() == 0 {
		config.FetchTimeout.Duration = defaultTimeout
	}
	if config.RequestDataInterval.Duration == 0 {
		config.RequestDataInterval.Duration = defaultRequestInterval
	}
	if config.MiningInterruptCheckInterval.Seconds() == 0 {
		config.MiningInterruptCheckInterval.Duration = defaultMiningInterrupt
	}

	NumProcessorsEnv := os.Getenv(NumProcessors)

	if NumProcessorsEnv == "" {

		if config.NumProcessors == 0 {
			config.NumProcessors = defaultCores
		}

	} else {
		parsedInt, err := strconv.ParseInt(NumProcessorsEnv, 10, 32)
		if err != nil {
			return fmt.Errorf("error in parsing NumProcessors from os.env: %s", err)
		}
		config.NumProcessors = int(parsedInt)
	}
	// TrackerSleepCycle, err := time.ParseDuration(config.TrackerSleepCycle)

	if config.TrackerSleepCycle.Duration == 0 {
		config.TrackerSleepCycle.Duration = defaultTrackerInterval
	}

	if config.Heartbeat.Seconds() == 0 {
		config.Heartbeat.Duration = defaultHeartbeat
	}

	if len(config.ServerWhitelist) == 0 {
		if strings.Contains(config.PublicAddress, "0x") {
			config.ServerWhitelist = append(config.ServerWhitelist, config.PublicAddress)
		} else {
			config.ServerWhitelist = append(config.ServerWhitelist, "0x"+config.PublicAddress)
		}
	}

	config.PrivateKey = strings.ToLower(strings.ReplaceAll(config.PrivateKey, "0x", ""))

	config.PublicAddress = strings.ToLower(strings.ReplaceAll(config.PublicAddress, "0x", ""))

	if config.DisputeThreshold == 0 {
		config.DisputeThreshold = DefaultDisputeThreshold
	}

	if config.DisputeTimeDelta.Duration == 0 {
		config.DisputeTimeDelta.Duration = DefaultMaxCheckTimeDelta
	}

	err = validateConfig(config)
	if err != nil {
		return fmt.Errorf("validation failed: %s", err)
	}
	return nil
}

func validateConfig(cfg *Config) error {
	b, err := hex.DecodeString(cfg.PublicAddress)
	if err != nil || len(b) != 20 {
		return fmt.Errorf("expecting 40 hex character public address, got \"%s\"", cfg.PublicAddress)
	}
	if cfg.EnablePoolWorker {
		if len(cfg.Worker) == 0 {
			return fmt.Errorf("worker name required for pool")
		}
		if len(cfg.Password) == 0 {
			return fmt.Errorf("password name required for pool")
		}
	} else {
		b, err = hex.DecodeString(cfg.PrivateKey)
		if err != nil || len(b) != 32 {
			return fmt.Errorf("expecting 64 hex character private key, got \"%s\"", cfg.PrivateKey)
		}
		if len(cfg.ContractAddress) != 42 {
			return fmt.Errorf("expecting 40 hex character contract address, got \"%s\"", cfg.ContractAddress)
		}
		b, err = hex.DecodeString(cfg.ContractAddress[2:])
		if err != nil || len(b) != 20 {
			return fmt.Errorf("expecting 40 hex character contract address, got \"%s\"", cfg.ContractAddress)
		}

		gasMultiplier := cfg.GasMultiplier

		if gasMultiplier < 0 || gasMultiplier > 20 {
			return fmt.Errorf("gas multiplier out of range [0, 20] %f", cfg.GasMultiplier)
		}
	}

	for name, gpuConfig := range cfg.GPUConfig {
		if gpuConfig.Disabled {
			continue
		}
		if gpuConfig.Count == 0 {
			return fmt.Errorf("gpu '%s' requires 'count' > 0", name)
		}
		if gpuConfig.GroupSize == 0 {
			return fmt.Errorf("gpu '%s' requires 'groupSize' > 0", name)
		}
		if gpuConfig.Groups == 0 {
			return fmt.Errorf("gpu '%s' requires 'groups' > 0", name)
		}
	}

	return nil
}

//GetConfig returns a shared instance of config
func GetConfig() *Config {
	return config
}
