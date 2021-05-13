package main

import (
	"net/http"
	"test/router"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := router.NewRouter()
	http.ListenAndServe(":8080", r.R)
}
