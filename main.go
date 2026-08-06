package main

import (
	"database/sql"
	"log"
	"simplebank/api"
	db "simplebank/db/sqlc"
	"simplebank/utils"

	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbDriver := config.DBDriver
	dbSource := config.DBSource
	serverAddress := config.ServerAddress

	log.Println("db driver:", dbDriver)
	log.Println("db source:", dbSource)
	log.Println("server address:", serverAddress)

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	if err = conn.Ping(); err != nil {
		log.Fatal("cannot ping db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	if err := server.Start(serverAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}
	log.Println("server started on", serverAddress)
}
