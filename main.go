package main

import (
	"github.com/dibrinsofor/urlplaylists/server"
)

func main() {
	r := server.SetupServer()

	r.Run(":8080")
}
