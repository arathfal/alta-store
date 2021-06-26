package main

import (
	"AltaStore/configs"
	"AltaStore/routes"
)


func main() {
	configs.InitDB()
	e := routes.New()
	e.Start(":8080")

}