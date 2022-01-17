# property-portal

### Features

- Search for properties
- Get property details
- Get Property images

#### Generate client & server stubs
```sh
cd propertyportal
protoc \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    property_portal.proto
```

#### Start server
```sh
go run server/server.go
```

#### Run client
```sh
go run client/client.go
```
