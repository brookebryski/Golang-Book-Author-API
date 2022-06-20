package main

import (
	"Golang-Book-Author-API/app"
	"Golang-Book-Author-API/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
