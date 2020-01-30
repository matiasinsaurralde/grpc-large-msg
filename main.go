package main

import (
	"log"
	"net"

	"net/http"

	"os"
	"strconv"

	coprocess "github.com/TykTechnologies/tyk-protobuf/bindings/go"
	"google.golang.org/grpc"
)

const (
	defaultGrpcMaxSize = 10000000
)

var grpcMaxSize int

func init() {
	if sz := os.Getenv("GRPC_MAX_SIZE"); sz == "" {
		grpcMaxSize = defaultGrpcMaxSize
	} else {
		i, err := strconv.Atoi(sz)
		if err != nil {
			panic(err)
		}
		grpcMaxSize = i
	}
	log.Printf("Setting grpcMaxSize to %d", grpcMaxSize)
}

const (
	ListenAddress   = ":9111"
	ManifestAddress = ":8888"
)

func main() {
	lis, err := net.Listen("tcp", ListenAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Starting grpc server on %v", ListenAddress)
	s := grpc.NewServer(
		grpc.MaxRecvMsgSize(grpcMaxSize),
		grpc.MaxSendMsgSize(grpcMaxSize),
	)
	coprocess.RegisterDispatcherServer(s, &Dispatcher{})
	go s.Serve(lis)
	log.Fatal(http.ListenAndServe(ManifestAddress, nil))
}
