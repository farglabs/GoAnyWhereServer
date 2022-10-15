package main

import (
	"freshmanual.com/db"
	"freshmanual.com/server"
)

func main() {
	db.InitDB()
	server.StartSSL()
}
