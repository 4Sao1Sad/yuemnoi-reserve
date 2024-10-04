package main

import (
	"github.com/KKhimmoon/yuemnoi-reserve/db"
	"github.com/KKhimmoon/yuemnoi-reserve/config"
)

func main() {
	cfg := config.Load()
	database := db.InitPostgreSQL(cfg)
	err := db.ServerInit(cfg, database)
	if err != nil {
		panic(err)
	}
}