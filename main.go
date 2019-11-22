package main

import "project-service/routes"

func main() {
	api := routes.ProjectServer{}
	api.Run()
}
