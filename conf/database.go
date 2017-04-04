package conf

type Database struct {
  User string
  Password string
  Name string
  Host string
  Port string
}

var Default = &Database {
  User: "root",
  Password: "",
  Name: "okr",
  Host: "localhost",
  Port: "3306",
}
