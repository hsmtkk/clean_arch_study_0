package router

import (
	"github.com/hsmtkk/clean_arch_study_0/pkg/infra/database"
	"github.com/hsmtkk/clean_arch_study_0/pkg/interface/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	sqlHandler, err := database.New()
	if err != nil {
		e.Logger.Fatal(err)
	}
	ctrl := user.New(sqlHandler)
	e.POST("/users", ctrl.Create)
	e.GET("/users", ctrl.Index)
	e.GET("/users/:id", ctrl.Show)
	e.Logger.Fatal(e.Start(":8000"))
}
