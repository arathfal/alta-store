package controllers

import (
	"AltaStore/configs"

	"github.com/labstack/echo"
)

func SetupEchoDB() *echo.Echo {
	configs.InitDBTest()
	e := echo.New()
	return e
}
