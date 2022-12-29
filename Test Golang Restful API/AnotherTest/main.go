package main

import (
	connection "AnotherTest/Connection"
	routes "AnotherTest/Routes"
)

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start(":2021"))
	connection.GetDB()
}
