package main

import (
	"wenxinhexuan/database"

	"github.com/gin-gonic/gin"
)

// @title Wenxinhexuan
// @version 1.0
// @description Wenxinhexuan Backend

func main() {
	database.InitDatabase()
	defer database.CloseDatabase()
	app := InitApp()
	app.Run()
}

type App struct {
	e *gin.Engine
}

func NewApp(e *gin.Engine) App {
	return App{
		e: e,
	}
}
func (a *App) Run() {
	a.e.Run(":8080")
}
