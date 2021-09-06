package main

import (
	"test/router"
)

func main() {
	r := router.NewRouter()
	r.Setup()
	r.Engine.Run(":8080")

}
