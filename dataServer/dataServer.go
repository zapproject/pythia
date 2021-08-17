package dataServer

import (
	"context"
	"log"

	"github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/db"
	"github.com/zapproject/pythia/local"
	"github.com/zapproject/pythia/rest"
	"github.com/zapproject/pythia/rpc"
	"github.com/zapproject/pythia/tracker"
	"github.com/zapproject/pythia/util"
)

// This package starts the rest server, db and runner for handling
// requests sent to the miner

//DataServer holds refs to primary stack of utilities for data retrieval and serving
type DataServer struct {
	server       *rest.Server
	localserver  *local.LocalServer
	DB           db.DB
	runner       *tracker.Runner
	ethClient    rpc.ETHClient
	exitCh       chan int
	runnerExitCh chan int
	Stopped      bool
	readyChannel chan bool
	log          *util.Logger
}

//CreateServer creates a data server stack and kicks off all go routines to start retrieving and serving data
func CreateServer(ctx context.Context) (*DataServer, error) {
	cfg := config.GetConfig()

	DB := ctx.Value(common.DBContextKey).(db.DB)
	client := ctx.Value(common.ClientContextKey).(rpc.ETHClient)
	run, err := tracker.NewRunner(client, DB)
	if err != nil {
		log.Fatal(err)
	}

	serverPort := cfg.ServerPort

	srv, err := rest.Create(ctx, cfg.ServerHost, serverPort)
	localsrv := local.CreateLocal(ctx)

	if err != nil {
		log.Fatalf("Error in creating remote proxy: %s", err)
	}

	//make sure channel buffer size 1 since there is no guarantee that anyone
	//would be listening to the channel
	ready := make(chan bool, 1)
	return &DataServer{server: srv,
		localserver:  localsrv,
		DB:           DB,
		runner:       run,
		ethClient:    client,
		exitCh:       nil,
		Stopped:      true,
		runnerExitCh: nil,
		readyChannel: ready,
		log:          util.NewLogger("dataServer", "DataServer")}, nil
}

//Start the data server and all underlying resources
func (ds *DataServer) Start(ctx context.Context, exitCh chan int) error {
	ds.exitCh = exitCh
	ds.runnerExitCh = make(chan int)
	ds.Stopped = false
	err := ds.runner.Start(ctx, ds.runnerExitCh)
	if err != nil {
		return err
	}

	ds.localserver.Start()
	ds.server.Start()
	go func() {
		<-ds.runner.Ready()
		ds.log.Info("Runner signaled it is ready")
		ds.readyChannel <- true
		ds.log.Info("DataServer ready for use")
		<-ds.exitCh
		ds.log.Info("DataServer received signal to stop")
		err = ds.stop(ctx)
		if err != nil {
			ds.log.Error("Data server can't be stopped")
		}
	}()
	return nil
}

//Ready provides notification channel that data from trackers is ready for use
func (ds *DataServer) Ready() chan bool {
	return ds.readyChannel
}

func (ds *DataServer) stop(ctx context.Context) error {
	//stop tracker run loop
	ds.runnerExitCh <- 1

	//stop REST erver
	ds.localserver.Stop()
	ds.server.Stop()
	// ds.server.Shutdown(ctx)

	//stop the DB
	ds.DB.Close()

	//stop the eth RPC client
	ds.ethClient.Close()

	//all done
	ds.Stopped = true
	return nil
}
