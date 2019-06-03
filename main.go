package main

import (
	"line"
	"sql"
)

func main() {
	chanSecret := "xxxxxxxx"
	chanToken := "xxxxxxxx"
	lineConf := line.NewLineConfig(chanToken, chanSecret)
	sqlDB := sql.NewSQLdb("localhost", "root", "password", "subscriber", 3306)
	server := line.NewServer(sqlDB, lineConf)
	server.Run()
}
