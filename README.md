# grpc-load-balancing
Example of gRPC load balancing on Kubernetes.

- server is deployed as headless service to expose all endpoints (pods) via dns
- client utilizes client side round robin load balancing over all reachable endpoints
