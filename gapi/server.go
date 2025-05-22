package gapi

import (
	"fmt"

	db "github.com/cluna2/simplebank/db/sqlc"
	"github.com/cluna2/simplebank/pb"
	"github.com/cluna2/simplebank/token"
	"github.com/cluna2/simplebank/util"
	"github.com/cluna2/simplebank/worker"
)

// serves  gRPC requests for our banking service
type Server struct {
	pb.UnimplementedSimplebankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// creates a new gRPC server
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
