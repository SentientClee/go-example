# Example Go service providing gRPC and JSON

Install dependencies
```sh
make install
```

Run the example service
```sh
docker-compose up example
```

Perform a request
```sh
curl "http://localhost:8080/v1/echo?message=hello" | json_pp
```
