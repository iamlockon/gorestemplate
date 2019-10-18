# Go Restful API Template

Template for golang restful API servers. Help you go up and running with Restful Go API server in no time.

## Usage
1. Clone the repository.
2. Install dependencies.
```
govendor init
govendor add +external
```
3. Format & Check syntax
```
$gofmt -w .
$go vet .
```
3. Change the configs in "conf/app.ini"
4. Build project
```
$go build .
```
## Dockerfile Usage
If you want to deploy your API in Docker environment, you can use the Dockerfile given. Since it built from scratch image, you will need to compile executable manually first, then build the image:
```
//Run database container, use mysql for example (you might need to create external volume directory first.)
$docker run --name <your-db-name> -p <db-exposed-port>:3306 -e MYSQL_ROOT_PASSWORD -v <external-volume-path>:/var/lib/mysql -d mysql
//Compile executable for Dockerfile build
$CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o <your-excutable-name> .
//Build docker image from executable
$docker build -t <your-executable-name>
//Run container & link with database container
$docker run --link <your-db-container>:<your-db-container-alias-name> -p 8001:8001 <your-executable-name>
```
## Features
    User Auth (jwt-go)
    Logging
    Pagination
    Configuration Management (ini)
    Cache (redis)
    ORM (gorm)
    Hot reload (endless)
    API Docs (Swagger)

For More, see:
https://github.com/EDDYCJY/go-gin-example

## Disclaimer
This repo is adapted from :
https://github.com/EDDYCJY/go-gin-example
