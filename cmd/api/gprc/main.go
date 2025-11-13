package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Evensee/user-service/internal"
	api "github.com/Evensee/user-service/internal/delivery/grpc"
	"github.com/Evensee/user-service/internal/dependency"
	"github.com/Evensee/user-service/internal/infrastructure/database"
	"google.golang.org/grpc"
)

func main() {
	const op = "UserServiceGrpc.Run"
	dbConfig, err := internal.LoadDatabaseConfig()

	if err != nil {
		panic(err)
	}

	appConfig := internal.MustLoadConfig[internal.AppConfig]()

	db := database.Connect(dbConfig)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(
		api.LoggingInterceptor,
	))

	appResolver := dependency.NewResolver(db)

	userGrpcService := api.New(appResolver)

	api.Register(grpcServer, userGrpcService)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", appConfig.GrpcApiPort))

	if err != nil {
		fmt.Printf("%s: %v", op, err)
	}

	fmt.Println("Starting gRPC Server")

	go func() {
		err = grpcServer.Serve(l)

		if err != nil {
			fmt.Printf("%s: %v", op, err)
		}
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	fmt.Println("Stopping gRPC Server...")
	grpcServer.GracefulStop()
	fmt.Println("gRPC Server Has Been Stopped")
}
