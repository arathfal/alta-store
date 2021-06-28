package main

import (
	"AltaStore/config"
	"AltaStore/routes"
)

func main() {
	configs.InitDB()
	e := routes.New()
	e.Start(":8080")
}
