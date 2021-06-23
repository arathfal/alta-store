package controllers

import (
	"strconv"

	"github.com/labstack/echo")


func DeleteCartControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	_ = id
	return err
}