package test

import (
	"context"
	"os"

	"github.com/zapproject/pythia/ops"
	"github.com/zapproject/pythia/pow"
	"github.com/zapproject/pythia/util"
)

type WorkSource interface {
	GetWork(input chan *pow.Work) (*pow.Work, bool)
}

type SolutionSink interface {
	Submit(context.Context, *pow.Result) bool
}

//MiningMgr holds items for mining and requesting data
type MiningMgr struct {
	//primary exit channel
	exitCh  chan os.Signal
	log     *util.Logger
	Running bool

	group      *pow.MiningGroup
	tasker     WorkSource
	solHandler SolutionSink
	solution   *pow.Result

	dataRequester *ops.DataRequester
	//data requester's exit channel
	requesterExit chan os.Signal
}
