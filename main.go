package main

import (
	"go-movies-crud/routes"
)

func main() {
	routes.InitializeRouter()
	routes.StartServer()
}
