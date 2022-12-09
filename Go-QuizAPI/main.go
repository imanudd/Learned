package main

import (
	"GO-QuizAPI/db"
	"GO-QuizAPI/routes"
)

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1234"))
	db.Init()
}
