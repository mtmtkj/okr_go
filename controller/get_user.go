package controller

import (
  "fmt"
  "strconv"
  "net/http"
  "github.com/labstack/echo"
  "github.com/gocraft/dbr"
  "okr_go/model"
)

func GetUser() echo.HandlerFunc {
  return func(c echo.Context) error {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
      return echo.NewHTTPError(http.StatusBadRequest, "invalid argument.")
    }
    fmt.Println(id)
    tx := c.Get("Tx").(*dbr.Tx)
    user := new(model.User)
    if err := user.Load(tx, id); err != nil {
      return echo.NewHTTPError(http.StatusNotFound, "user does not exists.")
    }
    return c.JSON(http.StatusOK, user)
  }
}