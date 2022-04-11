package main

import (
	app "myapp/scripts"
	cg "myapp/scripts/configs"
)

func main() {
	cg.ConnectToDatabase()
	app.Initialize()
	app.Run()
}
