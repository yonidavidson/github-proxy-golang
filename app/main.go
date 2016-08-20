package main

import (
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web/middleware"
	"os"
)

var MODE_DRY bool = false
var SECURITY_OPEN string = os.Getenv("SECURITY_OPEN")

func main() {
	goji.Get("/api/gh/:username", UserHandler)
	goji.Use(middleware.EnvInit)
	goji.Use(SuperSecure)
	goji.Serve()
}
