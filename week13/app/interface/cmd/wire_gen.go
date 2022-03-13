// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"gbmerp/app/interface/internal/biz"
	"gbmerp/app/interface/internal/conf"
	"gbmerp/app/interface/internal/data"
	"gbmerp/app/interface/internal/server"
	"gbmerp/app/interface/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

func initApp(confServer *conf.Server, confData *conf.Data, registry *conf.Registry, auth *conf.Auth, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(registry)
	userClient := data.NewUserServiceClient(discovery)
	staffClient := data.NewStaffServiceClient(discovery)
	dutyClient := data.NewDutyServiceClient(discovery)
	departmentClient := data.NewDeptServiceClient(discovery)
	dataData := data.NewData(confData, userClient, staffClient, dutyClient, departmentClient)
	userRepo := data.NewUserRepo(dataData, auth)
	authUseCase := biz.NewAuthUseCase(auth)
	userUseCase := biz.NewUserUseCase(userRepo, authUseCase)
	interfaceService := service.NewInterfaceService(userUseCase)
	grpcServer := server.NewGRPCServer(confServer, tracerProvider, interfaceService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
	}, nil
}