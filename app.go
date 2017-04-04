package main

import (
	"github.com/labstack/echo"
  "okr_go/conf"
  "okr_go/db"
  "okr_go/handler"
)

func main() {
	e := echo.New()
  e.Use(handler.TransactionHandler(db.Init()))

  conf.Routing(e)

	e.Logger.Fatal(e.Start(":1323"))
}
