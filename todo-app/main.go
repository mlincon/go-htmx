package main

import (
	model "todo-app/model"
	routes "todo-app/routes"
)

func main() {
	model.Setup()
	routes.SetupAndRun()
}
