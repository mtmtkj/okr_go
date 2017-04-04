package model

import (
  "github.com/gocraft/dbr"
)

type User struct {
  Id int64 `json:"id"`
  Name string `json:"name"`
  Email string `json:"email"`
}

func (m *User) Load(tx *dbr.Tx, id int64) error {
  return tx.Select("*").
    From("users").
    Where("id = ?", id).
    LoadStruct(m)
}
