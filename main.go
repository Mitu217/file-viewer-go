package main

import (
	"github.com/Mitu217/file-viewer/pkg/app"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		panic(err)
	}
	app.Run()
}
