package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/db"
)

//BalanceHandler handles balance requests
type TokenBalanceHandler struct {
}

//Incoming implementation for  handler
func (h *TokenBalanceHandler) Incoming(ctx context.Context, req *http.Request) (int, string) {
	DB := ctx.Value(common.DBContextKey).(db.DB)
	v, err := DB.Get(db.TokenBalanceKey)
	if err != nil {
		log.Printf("Problem reading TokenBalance from DB: %v\n", err)
		return 500, `{"error": "Could not read TokenBalance from DB"}`
	}
	return 200, fmt.Sprintf(`{ "TokenBalance": "%s" }`, string(v))
}
