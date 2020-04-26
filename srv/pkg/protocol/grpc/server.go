package grpc

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"srv/pkg/api/v1"
)

func RunServer(ctx context.Context, studentAPI v1.StudentServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register services
	server := grpc.NewServer()
	v1.RegisterStudentServiceServer(server,studentAPI)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	return server.Serve(listen)
}