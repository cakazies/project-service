package main

import "github.com/crowdeco/project-service/routes"

func main() {
	api := routes.ProjectServer{}
	api.Run()
}
