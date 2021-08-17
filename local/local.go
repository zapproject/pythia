package local

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/db"
	"github.com/zapproject/pythia/rest/routes"
	"github.com/zapproject/pythia/util"
)

var serverLog = util.NewLogger("local", "LocalServer")

type LocalServer struct {
	Server *http.Server
	Router *routes.Router
}

func CreateLocal(ctx context.Context) *LocalServer {
	cfg := config.GetConfig()
	DB := ctx.Value(common.DBContextKey).(db.DB)

	router := routes.NewRouter(DB)
	srv := &http.Server{Addr: fmt.Sprintf("%s:%d", cfg.ServerHost, (cfg.ServerPort + 1))}
	srv.Handler = router

	serverLog.Info("Defined server on %s:%v", cfg.ServerHost, cfg.ServerPort)
	return &LocalServer{Server: srv, Router: router}
}

//Start the Server listening for incoming requests
func (s *LocalServer) Start() {
	h := &routes.IndexHandler{}
	routes.BuildEndpoints(s.Router, h)
	go func() {
		serverLog.Info("Starting Local Server on %+v\n", s.Server.Addr)
		// returns ErrServerClosed on graceful close
		if err := s.Server.ListenAndServe(); err != http.ErrServerClosed {
			s.Stop()
			serverLog.Error("ListenAndServe(): %s", err)
		}
	}()
}

//Stop stops the Server listening
func (s *LocalServer) Stop() error {
	serverLog.Info("Winding down local Server")
	return s.Server.Close()
}
