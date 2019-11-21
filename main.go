package main

import "project-service/routes"

func main() {
	api := routes.ProjectService{}
	api.Run()
}
