package main

import (
	"log"

	"github.com/atakanozceviz/gorient.v3/controller"
)

func main() {
	port := "8080" //os.Getenv("PORT")
	log.Fatal(controller.StartServer("192.168.0.11:" + port))
}
