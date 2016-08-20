package main

import (
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web/middleware"
)

var MODE_DRY bool = false

func main() {
	goji.Get("/api/gh/:username", UserHandler)
	goji.Use(middleware.EnvInit)
	goji.Use(SuperSecure)
	goji.Serve()
}
