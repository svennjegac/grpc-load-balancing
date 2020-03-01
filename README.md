# grpc-load-balancing

Example of gRPC load balancing on Kubernetes.
Server is deployed as headless service to expose all endpoints (pods) via dns.
Client utilizes client side round robin load balancing over all reachable endpoints.
