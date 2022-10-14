package main

import (
	"freshmanual.com/server"
	"freshmanual.com/db"
)

func main() {
	_, err:= db.InitDB()
	if err != nil {
		return 
	}
	server.StartSSL()
}
