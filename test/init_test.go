package test

import (
	"log"

	project "github.com/cakazies/project-service/grpc"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file error : ", err)
	}
}

func serviceProject() project.ProjectsClient {
	port := ":6001"
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return project.NewProjectsClient(conn)
	// s, addr := serverSetUp(t, true, nil, math.MaxUint32, grpc.NewGZIPCompressor(), grpc.NewGZIPDecompressor(), e)
	// cc := clientSetUp(t, addr, grpc.NewGZIPCompressor(), grpc.NewGZIPDecompressor(), "", e)
	// // Unary call
	// tc := testpb.NewTestServiceClient(cc)
	// defer tearDown(s, cc)

	// return tc
}
