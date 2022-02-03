package main

import "github.com/JustForCodin/simplewebserver/server"

func main() {
	server.SetupRoutes()
	server.Listen()
}
