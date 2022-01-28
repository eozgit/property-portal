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
go run ./server/
```

#### Run client
```sh
go run ./client/
```

#### Build image
```sh
docker build --tag eozgit/property-portal .
```

#### Run container
```sh
docker run --tty --interactive --rm \
    --publish 10000:10000 \
    --name property-portal \
    eozgit/property-portal
```

#### Start interactive shell
```sh
docker exec --interactive --tty property-portal bash
```

#### Select properties
```sh
sqlite3 -header -column -echo \
    /tmp/property.db < ./server/sql/properties.sql
```

#### Sample properties
```sh
sqlite3 -header -column -echo \
    /tmp/property.db < ./server/sql/sample_properties.sql
```

#### Sample property-images
```sh
sqlite3 -header -column -echo \
    /tmp/property.db < ./server/sql/sample_property_images.sql
```
