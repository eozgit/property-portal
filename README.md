# property-portal

#### Demo
https://youtu.be/pw-mvUUZXb8

### Features

- Search for properties
- Get property details
- Get property images

#### Pull container image
```sh
docker pull eozgit/property-portal
```

#### Run container
```sh
docker run --rm --tty --interactive \
    --publish 10000:10000 \
    --name property-portal \
    eozgit/property-portal
```

#### Start interactive shell
```sh
docker exec --interactive --tty property-portal bash
```

#### Run client
```sh
go run ./client/
```

---

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

---

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

#### Build image
```sh
docker build --tag eozgit/property-portal:latest --tag eozgit/property-portal:YYMMDD .
```

#### Push image
```sh
docker image push eozgit/property-portal --all-tags
```

---


#### Docker repo

https://hub.docker.com/r/eozgit/property-portal
