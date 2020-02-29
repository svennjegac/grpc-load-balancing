package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/svennjegac/grpc-load-balancing/client/ip"
	"google.golang.org/grpc"
)

const (
	Target                 = "IP_SERVICE_TARGET"
	DefaultServiceConfig   = "DEFAULT_SERVICE_CONFIG"
	PodIP                  = "POD_IP"
	DialTimeoutMs          = "DIAL_TIMEOUT_MS"
	RequestTimeoutMs       = "REQUEST_TIMEOUT_MS"
	InBetweenRequestWaitMs = "IN_BETWEEN_REQUEST_WAIT_MS"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), duration(DialTimeoutMs))
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		os.Getenv(Target),
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(os.Getenv(DefaultServiceConfig)),
	)
	if err != nil {
		log.Fatalln("error dialing target;", err)
	}

	client := ip.NewIPServiceClient(conn)
	podIP := os.Getenv(PodIP)
	inBetweenRequestDuration := duration(InBetweenRequestWaitMs)
	for {
		ctx, cancel := context.WithTimeout(context.Background(), duration(RequestTimeoutMs))

		request := &ip.TellMeYourIPRequest{
			ClientIp: podIP,
		}
		response, err := client.TellMeYourIP(ctx, request)
		cancel()
		if err != nil {
			log.Fatalf("error making tell me your ip request; err=%+v, response=%+v\n", err, response)
		}
		if response == nil {
			log.Fatalln("error, got nil response")
		}

		log.Println("got response from server:", response.ServerIp)
		time.Sleep(inBetweenRequestDuration)
	}
}

func duration(envVar string) time.Duration {
	val, err := strconv.Atoi(os.Getenv(envVar))
	if err != nil {
		log.Fatalf("error parsing duration parameter; param=%s\n", envVar)
	}
	return time.Millisecond * time.Duration(val)
}
