package tracker

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	zapCommon "github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/config"
	zap "github.com/zapproject/zap-miner/contracts"
	"github.com/zapproject/zap-miner/db"
	"github.com/zapproject/zap-miner/rpc"
	"github.com/zapproject/zap-miner/util"
)

var runnerLog = util.NewLogger("tracker", "Runner")

//Runner will execute all configured trackers
type Runner struct {
	client       rpc.ETHClient
	db           db.DB
	readyChannel chan bool
}

//NewRunner will create a new runner instance
func NewRunner(client rpc.ETHClient, db db.DB) (*Runner, error) {
	return &Runner{client: client, db: db, readyChannel: make(chan bool)}, nil
}

//Start will kick off the runner until the given exit channel selects.
func (r *Runner) Start(ctx context.Context, exitCh chan int) error {
	cfg := config.GetConfig()
	trackerNames := cfg.Trackers
	defaultTrackers := []string{"newCurrentVariables", "timeOut"}
	var trackers []Tracker
	for _, name := range trackerNames {
		t, err := createTracker(name)
		if err != nil {
			runnerLog.Error("Problem creating tracker: %s\n", err.Error())
			continue
		}
		trackers = append(trackers, t...)
	}
	for _, name := range defaultTrackers {
		t, err := createTracker(name)
		if err != nil {
			runnerLog.Error("Problem creating tracker: %s\n", err.Error())
			continue
		}
		trackers = append(trackers, t...)
	}
	if len(trackers) == 0 {
		return nil
	}
	runnerLog.Info("Created %d trackers", len(trackers))

	var err error
	masterInstance := ctx.Value(zapCommon.MasterContractContextKey)
	if masterInstance == nil {
		contractAddress := common.HexToAddress(cfg.ContractAddress)
		masterInstance, err = zap.NewZapMaster(contractAddress, r.client)
		if err != nil {
			runnerLog.Error("Problem creating zap master instance: %v\n", err)
			return err
		}
		ctx = context.WithValue(ctx, zapCommon.MasterContractContextKey, masterInstance)
	}

	runnerLog.Info("Trackers will run every %v\n", cfg.TrackerSleepCycle)

	// trackerSleepCycle, err := time.ParseDuration(cfg.TrackerSleepCycle)

	ticker := time.NewTicker(cfg.TrackerSleepCycle.Duration / time.Duration(len(trackers)))

	if ctx.Value(zapCommon.ClientContextKey) == nil {
		ctx = context.WithValue(ctx, zapCommon.ClientContextKey, r.client)
	}
	if ctx.Value(zapCommon.DBContextKey) == nil {
		ctx = context.WithValue(ctx, zapCommon.DBContextKey, r.db)
	}

	//after first run, let others know that tracker output data is ready for use
	doneFirstExec := make(chan bool, len(trackers))
	go func(n int) {
		for i := 0; i < n; i++ {
			<-doneFirstExec
		}
		r.readyChannel <- true
	}(len(trackers))
	runnerLog.Info("Waiting for trackers to complete initial requests")

	//run the trackers until we quit
	go func() {
		i := 0
		for {
			select {
			case _ = <-exitCh:
				{
					runnerLog.Info("Exiting run loop")
					ticker.Stop()
					return
				}
			case _ = <-ticker.C:
				{
					//runnerLog.Info("Running trackers...")
					go func(count int) {
						idx := count % len(trackers)
						err := trackers[idx].Exec(ctx)
						if err != nil {
							runnerLog.Error("Problem in tracker %s: %v\n", trackers[idx].String(), err)
						}
						//only increment this the first time a tracker is run
						if count < len(trackers) {
							doneFirstExec <- true
						}
					}(i)
					i++
				}
			}
		}
	}()

	return nil
}

//Ready provides notification channel to know that the tracker data output is ready for use
func (r *Runner) Ready() chan bool {
	return r.readyChannel
}
