package main

import (
	"ddb-go-demo/internal/db"
	"ddb-go-demo/internal/router"
	"ddb-go-demo/internal/server"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := db.DataBaseConnect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	ctx := server.Context{
		Db: db,
	}

	r := router.ServeRoutes(ctx)
	server := server.NewServer(r)

	port := os.Getenv("API_PORT")
	fmt.Printf("Server  running on 0.0.0.0:%s\n", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting server : %v", err)
	}
}
