package  webview
import (
	// "time"
	"fmt"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/setup"
	// "github.com/zapproject/pythia/ops"
	ZapCommon "github.com/zapproject/pythia/common"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zapproject/pythia/contracts"
)


func configWallet(cfg string) (bool, error){
	// fmt.Println(req.TokenAddress )
	err := config.ParseConfigBytes([]byte(cfg))
	if err != nil{
		fmt.Println(err.Error())
	} 

	return setup.App()
}


func getBalance() (string, error) {
	// setup.App()
	addr := setup.CTX.Value(ZapCommon.PublicAddress).(common.Address)
	instance := setup.CTX.Value(ZapCommon.MasterContractContextKey).(*contracts.ZapMaster)
	zapBalance, err := instance.BalanceOf(nil, addr)

	if err != nil{
		fmt.Print(err.Error())
		return "",err
	}

	return zapBalance.String(),nil

}

func stakeStatus() (int64, error) {
	  
	tmaster := setup.CTX.Value(ZapCommon.MasterContractContextKey).(*contracts.ZapMaster)

	publicAddress := setup.CTX.Value(ZapCommon.PublicAddress).(common.Address)
	status, _, err := tmaster.GetStakerInfo(nil, publicAddress)
	if err != nil {
		return 0,fmt.Errorf("failed to get stake status: %s", err.Error())
	}

	fmt.Println(status)

	return status.Int64(),nil

}