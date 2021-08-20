package  webview
import (
	// "time"
	"fmt"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/setup"
	"github.com/zapproject/pythia/ops"
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

func stakeStatus() (string, error) {
	  
	err := ops.ShowStatus(setup.CTX)
	if err != nil{
		fmt.Print(err.Error())
		return "",err
	}

	return "",nil

}