package main

import "Belajar-Go-Echo/infrastructure/http/server"

func main() {
	app := server.Server()

	app.Start(":8080")
}