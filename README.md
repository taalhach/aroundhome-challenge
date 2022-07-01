# Aroundhome Code challenge

Aroundhome code challenge is a REST API server which proposes the right
partner based on the details of a customer's flooring project.
  
  ## Frameworks, database and ORM used
  - [Echo](https://echo.labstack.com)
  - [Gorm](https://gorm.io/docs/)
  - [Postgis](https://postgis.net)
  - [swaggo/swag](https://github.com/swaggo/swag) (*for API specs*)
## How to build and run server
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
 

### Run Unit Tests 
Use following command for unit tests.
```
make test-unit
```

