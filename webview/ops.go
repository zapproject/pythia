package  webview
import (
	// "time"
	"fmt"
	"github.com/zapproject/pythia/config"

)



type Config struct {
	TokenAddress                 string                `json:"zapTokenAddress"`
	ContractAddress              string                `json:"contractAddress"`
	VaultAddress                 string                `json:"vaultAddress"`
	NodeURL                      string                `json:"nodeURL"`
	PublicAddress                string                `json:"publicAddress"`
	EthClientTimeout             uint                  `json:"ethClientTimeout"`
	TrackerSleepCycle            string              `json:"trackerCycle"`
	Trackers                     []string              `json:"trackers"`
	DBFile                       string                `json:"dbFile"`
	ServerHost                   string                `json:"serverHost"`
	ServerPort                   uint                  `json:"serverPort"`
	GasMultiplier                float32               `json:"gasMultiplier"`
	GasMax                       uint                  `json:"gasMax"`
	ServerWhitelist              []string              `json:"serverWhitelist"`
	DisputeTimeDelta             string              `json:"disputeTimeDelta"` //ignore data further than this away from the value we are checking
	UseGPU                       bool                  `json:"useGPU"`
	//config parameters excluded from the json config file
	PrivateKey string `json:"privateKey"`
}


func applyConfig(cfg string) (bool, error){
	// fmt.Println(req.TokenAddress )
	err := config.ParseConfigBytes([]byte(cfg))
	if err != nil{
		fmt.Println(err.Error())
	} 

	return true,nil
}




