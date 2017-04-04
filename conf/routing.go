package conf

import (
  "github.com/labstack/echo"
  "okr_go/controller"
)

func Routing(e *echo.Echo) {
  e.GET("/user/:id", controller.GetUser())
}