# Bsdex Code challenge

A rest server which provides CURD APIs for page data.
  
  ## Frameworks, database and ORM used
  - [Echo](https://echo.labstack.com)
  - [Gorm](https://gorm.io/docs/)
  - [Sqlite](https://www.sqlite.org/index.html)

## Docker
First build docker image with `docker-compose` command
```
docker-compose build
```
 
 then fire up docker containers using following command
```
docker-compose up -d
```
  Above commands will run a http server on port 3000. http://localhost:3000 which will redirect to swagger API specs page. 
 
## Local Setup

### Build 
Run make command it will download all libraries and builds binary files.
```
make build
```

### Run server 
Use following command in order to run http server.
```
make run
```

### Run Unit Tests 
Use following command for unit tests.
```
make test-unit
```

## Useful commands
Use help argument to list all supported commands
```
./bin/challenge --help
```

#### Serve api
`serve_api` command starts server on `3000` port.

```
./bin/challenge serve_api
```

