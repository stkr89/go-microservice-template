# go-microservice-template

## Steps

- Rename `pb/math.proto` to `pb/<myservice>svc.proto`
- Update `pb/<myservice>svc.proto`
- Update proto file name in `pb/compile.sh`
- Run `pb/compile.sh` to generate `pb/<my_service>.pb.go`
- Add new methods to `Service` interface in `service/api.go` and implementations in `service` struct
- Add new endpoint in `endpoints/endpoints.go`
- Add new transport handler in `transports/grpc.go`