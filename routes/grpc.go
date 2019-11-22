package routes

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"
	ctr "project-service/controllers"
	rpc "project-service/grpc"
	"strconv"

	"google.golang.org/grpc"
)

type (
	// ProjectServer struct for class this server
	ProjectServer struct{}
)

// Run function is first time load
func (ProjectServer) Run() {
	srv := grpc.NewServer()
	var garageSrv ProjectServer
	rpc.RegisterProjectsServer(srv, garageSrv)
	port := os.Getenv("APPS_PORT")
	log.Println("Starting RPC server at", port)

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", port, err)
	}
	log.Fatal(srv.Serve(l))
}

// List function for get all data
func (ProjectServer) List(ctx context.Context, paging *rpc.Pagination) (*rpc.Reponse, error) {
	grpc := ctr.GrpcRoute{}
	result := grpc.GetProjects(paging)
	var resp rpc.Reponse

	resp.Data = string(result)
	return &resp, nil
}

// Create function for get all data
func (ProjectServer) Create(ctx context.Context, pro *rpc.Project) (*rpc.Reponse, error) {
	grpc := ctr.GrpcRoute{}
	data, _ := json.Marshal(pro)
	result := grpc.Create(string(data))
	var resp rpc.Reponse

	resp.Data = string(result)
	return &resp, nil
}

// Detail function for get detail data project
func (ProjectServer) Detail(ctx context.Context, pro *rpc.Project) (*rpc.Reponse, error) {
	grpc := ctr.GrpcRoute{}
	idProject := strconv.Itoa(int(pro.Id))
	result := grpc.GetProject(idProject)
	var resp rpc.Reponse

	resp.Data = string(result)
	return &resp, nil
}

// Edit function for update data project
func (ProjectServer) Edit(ctx context.Context, pro *rpc.Project) (*rpc.Reponse, error) {
	var resp rpc.Reponse
	grpc := ctr.GrpcRoute{}

	id := strconv.Itoa(int(pro.Id))
	pro.Id = 0

	data, _ := json.Marshal(pro)
	err := grpc.Edit(id, string(data))
	if err != nil {
		log.Println(err)
		return &resp, err
	}

	resp.Data = string(data)
	return &resp, nil
}
