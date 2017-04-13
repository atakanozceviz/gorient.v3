package main

import (
	"log"
	"os"

	"github.com/atakanozceviz/gorient.v3/controller"
)

func main() {
	port := os.Getenv("PORT")
	log.Fatal(controller.StartServer(":" + port))
}
