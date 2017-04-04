package handler

import (
  "github.com/gocraft/dbr"
  "github.com/labstack/echo"
)

const ContextKey = "Tx"

func TransactionHandler(db *dbr.Session) echo.MiddlewareFunc {
  return func(next echo.HandlerFunc) echo.HandlerFunc {
    return echo.HandlerFunc(func (c echo.Context) error {
      tx, _ := db.Begin()
      c.Set(ContextKey, tx)
      if err := next(c); err != nil {
        tx.Rollback()
        return err
      }
      tx.Commit()
      return nil
    })
  }
}
