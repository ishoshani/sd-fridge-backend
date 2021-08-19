package main

import (
	"sdfridgeproject.com/m/v2/db"
	"sdfridgeproject.com/m/v2/routes"
)

func main() {
	db.Initialize()
	routes.App()
}
