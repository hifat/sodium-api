//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/hifat/sodium-api/internal/handler"
	"github.com/hifat/sodium-api/internal/handler/authHandler"
	"github.com/hifat/sodium-api/internal/repository"
	"github.com/hifat/sodium-api/internal/repository/authRepo"
	"github.com/hifat/sodium-api/internal/repository/userRepo"
	"github.com/hifat/sodium-api/internal/service/authService"
	"github.com/hifat/sodium-api/internal/service/middlewareService"
)

var HandlerSet = wire.NewSet(
	authHandler.AuthHandlerSet,
	handler.HandlerSet,
)

var ServiceSet = wire.NewSet(
	authService.AuthServiceSet,
	middlewareService.AuthMiddlewareServiceSet,
)

var RepoSet = wire.NewSet(
	repository.GormDBSet,
	authRepo.AuthRepoSet,
	userRepo.UserRepoSet,
)

func InitializeAPI() (handler.Handler, func()) {
	wire.Build(RepoSet, ServiceSet, HandlerSet)
	return handler.Handler{}, nil
}
