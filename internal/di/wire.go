//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/hifat/sodium-api/internal/adapter"
	"github.com/hifat/sodium-api/internal/handler"
	"github.com/hifat/sodium-api/internal/handler/authHandler"
	"github.com/hifat/sodium-api/internal/middleware"
	"github.com/hifat/sodium-api/internal/repository"
	"github.com/hifat/sodium-api/internal/repository/authRepo"
	"github.com/hifat/sodium-api/internal/repository/userRepo"
	"github.com/hifat/sodium-api/internal/service/authService"
	"github.com/hifat/sodium-api/internal/service/middlewareService"
)

var AdapterSet = wire.NewSet(
	adapter.AdapterSet,
)

var MiddlewareSet = wire.NewSet(
	middleware.MiddlewareSet,
	middleware.AuthMiddlewareSet,
)

var RepoSet = wire.NewSet(
	repository.GormDBSet,
	authRepo.AuthRepoSet,
	userRepo.UserRepoSet,
)

var ServiceSet = wire.NewSet(
	authService.AuthServiceSet,
	middlewareService.AuthMiddlewareServiceSet,
)

var HandlerSet = wire.NewSet(
	authHandler.AuthHandlerSet,
	handler.HandlerSet,
)

func InitializeAPI() (adapter.Adapter, func()) {
	wire.Build(AdapterSet, MiddlewareSet, RepoSet, ServiceSet, HandlerSet)
	return adapter.Adapter{}, nil
}
