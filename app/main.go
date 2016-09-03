package main

import (
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web/middleware"
	"log"
	"os"
)

var MODE_DRY string = os.Getenv("MODE_DRY")
var SECURITY_OPEN string = os.Getenv("SECURITY_OPEN")

func main() {
	log.Println("MODE_DRY:", MODE_DRY, " SECURITY_OPEN:", SECURITY_OPEN)
	goji.Get("/api/gh/:username", UserHandler)
	goji.Use(middleware.EnvInit)
	goji.Use(SuperSecure)
	goji.Use(JsonText)
	goji.Serve()
}
