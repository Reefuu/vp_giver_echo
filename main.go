package main

import (
	"vp_week11_echo/routes"
	"vp_week11_echo/db"
)
func main() {

	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":7070"))

}
