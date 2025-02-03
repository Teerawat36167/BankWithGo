package gapi

import (
	"fmt"

	db "github.com/Teerawat36167/BankWithGo/db/sqlc"
	"github.com/Teerawat36167/BankWithGo/pb"
	"github.com/Teerawat36167/BankWithGo/token"
	"github.com/Teerawat36167/BankWithGo/util"
	"github.com/Teerawat36167/BankWithGo/worker"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

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
