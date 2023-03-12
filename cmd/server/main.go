package main

import (
	"time"
	"yadp/server"
)

func main() {
	server := &server.Server{
		Host:         "127.0.0.1",
		Port:         1053,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	server.Run()
}
