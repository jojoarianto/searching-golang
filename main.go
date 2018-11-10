package main

import "github.com/jojoarianto/searching-golang/app"

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":8000")
}
