package  webview
import (
	// "time"
	"fmt"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/setup"
	// "github.com/zapproject/pythia/ops"

)


func configWallet(cfg string) (bool, error){
	// fmt.Println(req.TokenAddress )
	err := config.ParseConfigBytes([]byte(cfg))
	if err != nil{
		fmt.Println(err.Error())
	} 

	return setup.App()
}