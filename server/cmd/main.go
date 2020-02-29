package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/svennjegac/grpc-load-balancing/server/ip"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
)

const (
	MaxConnectionIdleMs     = "MAX_CONNECTION_IDLE_MS"
	MaxConnectionAgeMs      = "MAX_CONNECTION_AGE_MS"
	MaxConnectionAgeGraceMs = "MAX_CONNECTION_AGE_GRACE_MS"
	TimeMs                  = "TIME_MS"
	TimeoutMs               = "TIMEOUT_MS"
	PodIP                   = "POD_IP"
	ServerPort              = "SERVER_PORT"
)

func main() {
	// create grpc server and register ip service handler
	grpcServer := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     duration(MaxConnectionIdleMs),
		MaxConnectionAge:      duration(MaxConnectionAgeMs),
		MaxConnectionAgeGrace: duration(MaxConnectionAgeGraceMs),
		Time:                  duration(TimeMs),
		Timeout:               duration(TimeoutMs),
	}))
	ipService := &ipService{ip: os.Getenv(PodIP)}
	ip.RegisterIPServiceServer(grpcServer, ipService)

	// initialize listener for incoming tcp connections
	port := os.Getenv(ServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("error listening for tcp connections; port=%s, err=%s\n", port, err)
	}
	defer lis.Close()

	// start listening for grpc requests
	log.Printf("ip service started; port=%s\n", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("error serving grpc requests; err=%s\n", err)
	}
}

type ipService struct {
	ip string
}

func (i *ipService) TellMeYourIP(ctx context.Context, req *ip.TellMeYourIPRequest) (*ip.TellMeYourIPResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "nil request received")
	}

	log.Println("handling request;", req.ClientIp)

	return &ip.TellMeYourIPResponse{
		ServerIp: i.ip,
	}, nil
}

func duration(envVar string) time.Duration {
	val, err := strconv.Atoi(os.Getenv(envVar))
	if err != nil {
		log.Fatalf("error parsing duration parameter; param=%s\n", envVar)
	}
	return time.Millisecond * time.Duration(val)
}
