package cmd

import (
	"context"
	"flag"
	"fmt"

	"srv/pkg/protocol/grpc"
	"srv/pkg/service/v1"
)

// Config is configuration for Server
type Config struct {
	GRPCPort string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "port", "", "gRPC port to bind")
	cfg.GRPCPort = "5000"
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	API := v1.NewStudentServiceServer()
	//return grpc.RunServer(ctx, API, cfg.GRPCPort)
	return grpc.RunServer(ctx, API, "5000")
}
