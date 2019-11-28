# Project Service

this service about project, all about project, create project, publish project, update project and the others.

**Collection Postman**
you can import collection in 
> `test/postman/API Golang.postman_collection.json`

**How to run**
- you can configuration config (**Database**) in (`env.dev`) to `.env` and compare your local configs and configs BNI
- if you development you can set variable `APP_ENV` in file `.env` to **STAGING**, if you already testing and ready publish you can set variable to **PRODUCTION** 

- install dependencies go get
	-  `go get github.com/go-sql-driver/mysql`
	-  `go get github.com/jinzhu/gorm`
    -  `github.com/streadway/amqp`
	-  `go get github.com/joho/godotenv`
	-  `github.com/satori/go.uuid`
	-  `github.com/dgrijalva/jwt-go`
	-  `github.com/gin-gonic/gin`
    -  `github.com/local/testify/assert`

## Migration 
run migration with this command 
> go run configs/migrate/migration.go

## Fitur
about all this fitur you can read this repo in WIKI or click this [link]()

## Unit Testing
run Testing with this command
> go test ./test

## Run Local
- Run `go run main.go`
  

## Run Docker
- build docker `docker build -t project-service .`
- Run `docker run -it --rm --name 002-project-service project-service`
