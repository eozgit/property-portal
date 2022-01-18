# property-portal

### Features

- Search for properties
- Get property details
- Get property images

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

#### Build image
```sh
docker build --tag eozgit/property-portal .
```

### Run container
```sh
docker run --tty --interactive --rm --publish 10000:10000 --name property-portal eozgit/property-portal
```

### Start interactive shell
````sh
docker exec --interactive --tty property-portal bash
```
