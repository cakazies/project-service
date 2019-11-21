package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"strconv"

	rpc "project-service/grpc"
	"project-service/routes"

	"google.golang.org/grpc"
)

type (
	// ProjectServer struct for class this server
	ProjectServer struct{}
)

func main() {
	// api := routes.ProjectService{}
	// api.Run()

	srv := grpc.NewServer()
	var garageSrv ProjectServer
	rpc.RegisterProjectsServer(srv, garageSrv)
	port := ":6001"
	log.Println("Starting RPC server at", port)

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", port, err)
	}
	log.Fatal(srv.Serve(l))
}

// List function for get all data
func (ProjectServer) List(ctx context.Context, paging *rpc.Pagination) (*rpc.Reponse, error) {
	grpc := routes.GrpcRoute{}
	result := grpc.GetProjects(paging)
	var resp rpc.Reponse

	resp.Data = string(result)
	return &resp, nil
}

// Create function for get all data
func (ProjectServer) Create(ctx context.Context, pro *rpc.Project) (*rpc.Reponse, error) {
	grpc := routes.GrpcRoute{}
	data, _ := json.Marshal(pro)
	result := grpc.Create(string(data))
	var resp rpc.Reponse

	resp.Data = string(result)
	return &resp, nil
}

// Detail function for get detail data project
func (ProjectServer) Detail(ctx context.Context, pro *rpc.Project) (*rpc.Reponse, error) {
	grpc := routes.GrpcRoute{}
	idProject := strconv.Itoa(int(pro.Id))
	result := grpc.GetProject(idProject)
	var resp rpc.Reponse

	resp.Data = string(result)
	return &resp, nil
}
