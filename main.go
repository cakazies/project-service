package main

import "github.com/cakazies/project-service/routes"

func main() {
	api := routes.ProjectServer{}
	api.Run()
}
