package ops

import (
	"context"

	zapCommon "github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/db"
	"github.com/zapproject/pythia/rpc"
)

//TxnSubmitter just concrete type for txn submitter
type TxnSubmitter struct {
}

//NewSubmitter creates a new TxnSubmitter instance
func NewSubmitter() TxnSubmitter {
	return TxnSubmitter{}
}

//PrepareTransaction relies on rpc package to prepare and submit transactions
func (s TxnSubmitter) PrepareTransaction(ctx context.Context, proxy db.DataServerProxy, ctxName string, callback zapCommon.TransactionGeneratorFN) error {
	return rpc.PrepareContractTxn(ctx, proxy, ctxName, callback)
}
