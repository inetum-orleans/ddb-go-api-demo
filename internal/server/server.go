package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Server struct {
	port int
}

func NewServer(router http.Handler) *http.Server {

	port, _ := strconv.Atoi(os.Getenv("API_PORT"))

	NewServer := &Server{
		port: port,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", NewServer.port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
