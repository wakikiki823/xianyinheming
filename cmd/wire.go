//go:build wireinject
// +build wireinject

package main

import (
	"wenxinhexuan/controllers"
	"wenxinhexuan/dao"
	"wenxinhexuan/routes"
	"wenxinhexuan/service"

	"github.com/google/wire"
)

func InitApp() App {
	wire.Build(
		dao.NewUserDAO,
		dao.NewSongDAO,
		service.NewSongService,
		service.NewUserService,
		controllers.NewUserController,
		controllers.NewSongController,
		routes.RegisterRoutes,
		NewApp,
	)
	return App{}
}
