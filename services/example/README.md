# Example Service

Directory structure
|Directory|Purpose|
|---------|-------|
|config|Runtime configuration parsing for the service|
|pkg|Contains interfaces (pkg/\*) and implementations (pkg/\*/\*) for (typically) network resources|
|service|Contains business logic for serving the API of the service. The service implementation should be agnostic about transport (HTTP/gRPC/TCP).|
|transport|Contains logic for providing the service over the implemented transport. For example, `transport/http` should serve the service over HTTP and `transport/gRPC` should serve the service over gRPC.|
|.|The `main` package loads runtime configuration, initializes clients for (typically) network resources (from pkg/\*)), initializes the service instance, and serves the service using one or more transport implementations.|
