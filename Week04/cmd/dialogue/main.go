package main

import (
	"log"

	"dialogue/api/dialogue/service/v1"
	"dialogue/internal/service"
	"dialogue/pkg/apps"
	"dialogue/pkg/transport/grpc"
)

func main() {
	// transport server
	grpcSrv := grpc.NewServer(":9000")

	// register service
	gs, closeFunc, err := service.InitDialogueService()
	if err != nil {
		panic(err)
	}
	defer closeFunc()

	v1.RegisterDialogueServiceServer(grpcSrv.Server, gs)

	// application lifecycle
	app := apps.New()
	app.Append(apps.Hook{OnStart: grpcSrv.Start, OnStop: grpcSrv.Stop})

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		log.Printf("startup failed: %v\n", err)
	}
}
