# grpc-load-balancing
Example of gRPC load balancing on Kubernetes.

- The server is deployed as a headless service to expose all endpoints (pods) via DNS
- The client utilizes client-side round-robin load balancing over all reachable endpoints
