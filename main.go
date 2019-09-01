package main

import (
	_ "go-fast-ws-mq/conf"
	"go-fast-ws-mq/protocol/http"
	"go-fast-ws-mq/routers"
	"log"
	"os"
)

func main() {
	router := routers.NewRouter()
	server := http.NewHttpProtocol(router)

	err := server.Start(os.Getenv("ADDR"))

	if err != nil {
		log.Fatal(err)
	}
}
