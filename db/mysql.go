package db

import (
  _ "github.com/go-sql-driver/mysql"
  "github.com/gocraft/dbr"
  "okr_go/conf"
)

func Init() *dbr.Session {
  session, err := getSession()
  if err != nil {
    return nil
  }
  return session
}

func getSession() (*dbr.Session, error) {
  var conn string
  conn = getSessionByConf(conf.Default)
  db, err := dbr.Open("mysql", conn, nil)
  if err != nil {
    return nil, err
  }
  session := db.NewSession(nil)
  return session, nil
}

func getSessionByConf(conf *conf.Database) string {
  return conf.User + ":" + conf.Password + "@tcp(" + conf.Host + ":" + conf.Port + ")/" + conf.Name
}
