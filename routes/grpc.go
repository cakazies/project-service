package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	ctr "github.com/cakazies/project-service/controllers"
	rpc "github.com/cakazies/project-service/grpc"

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
		return &resp, fmt.Errorf("This ID project %s doesn't exist", id)
	}

	resp.Data = string(data)
	return &resp, nil
}

// ListGallery function for update data project
func (ProjectServer) ListGallery(ctx context.Context, paging *rpc.Pagination) (*rpc.Reponse, error) {
	var resp rpc.Reponse
	grpc := ctr.GrpcRoute{}

	idProject := paging.Params
	result, err := grpc.GetGallery(idProject)
	if err != nil {
		return &resp, fmt.Errorf("This ID project %s doesn't exist", idProject)
	}

	resp.Data = string(result)
	return &resp, nil
}

// CreateGallery function for get all data
func (ProjectServer) CreateGallery(ctx context.Context, pro *rpc.ProjectGallery) (*rpc.Reponse, error) {
	var resp rpc.Reponse

	grpc := ctr.GrpcRoute{}
	data, _ := json.Marshal(pro)
	result, err := grpc.CreateGallery(string(data))

	if err != nil {
		return &resp, fmt.Errorf("This ID project %v doesn't exist", pro.ProjectId)
	}

	resp.Data = string(result)
	return &resp, nil
}

// UpdateGallery function for update data project
func (ProjectServer) UpdateGallery(ctx context.Context, pro *rpc.ProjectGallery) (*rpc.Reponse, error) {
	var resp rpc.Reponse
	grpc := ctr.GrpcRoute{}

	id := int(pro.Id)
	pro.Id = 0

	data, _ := json.Marshal(pro)
	err := grpc.UpdateGallery(id, string(data))
	if err != nil {
		return &resp, fmt.Errorf("This ID project %v doesn't exist", id)
	}

	resp.Data = string(data)
	return &resp, nil
}
