package main

import (
	"log"
	"textQuest/Server"
)

func main() {
	log.Fatal(Server.RunHTTPServer(":8080"))
}
