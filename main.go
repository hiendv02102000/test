package main

import (
	"test/router"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	r := router.NewRouter()
	r.Engine.Run(":8080")
}
