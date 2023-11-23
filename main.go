package main

import (
	"database/sql"
	"log"
	
	"github.com/QuangPham789/demo-bank/api"
	db "github.com/QuangPham789/demo-bank/db/sqlc"
	"github.com/QuangPham789/demo-bank/util"
	
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
