package main

import (
	"fileserver/db"
	"fileserver/server"
)

func main() {
	db.InitDB()
	server.StartSSL()
}
