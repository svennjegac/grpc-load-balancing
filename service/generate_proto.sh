#!/bin/bash

# set script directory as current directory (for relative paths)
cd $(dirname $0)

# generate proto file
protoc --proto_path=. --go_out=plugins=grpc:.  ip_service.proto

# cleanup client protobuf generated file
if [ -d "../client/ip" ]; then
  rm -Rf ../client/ip;
fi

# create client protobuf generated file
mkdir ../client/ip
cp ip_service.pb.go ../client/ip/ip_service.pb.go

# cleanup server protobuf generated file
if [ -d "../server/ip" ]; then
  rm -Rf ../server/ip;
fi

# create server protobuf generated file
mkdir ../server/ip
cp ip_service.pb.go ../server/ip/ip_service.pb.go

# cleanup current directory unnecessary protobuf generated file
rm -Rf ip_service.pb.go