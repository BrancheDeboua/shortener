# Shortener

Small shortener written in go. The goal of this project is to better understand how to deploy simple web service using docker compose to keep multiple containers up.

## Dev

```bash
# For dev application, you'll need to have a postgresql database up and running. With docker you can use the following: 
$ docker run --name container-name -e POSTGRES_PASSWORD=password -p 5432:5432 -d database

# Then you need to start the app locally by running the main.go
$ go build main.go
$ ./main

# OR
$ go run main.go
``` 

## TODO

- Add tests
- Add gh actions
- Dockerfile and docker compose
