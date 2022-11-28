package main

import (
	"vp_giver_echo/routes"
	"vp_giver_echo/db"
)
func main() {

	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":7070"))

}
