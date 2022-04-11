package tools

// DockerIgnore .dockerignore
const DockerIgnore = "/vendor"

// DockerCompose docker-compose.yml
const DockerCompose = `version: '3.7'
services:
  postgresql:
    image: postgres:14-alpine
    healthcheck:
      test: [ "CMD-SHELL", "pg_isready -U postgres" ]
      interval: 5s
      timeout: 5s
      retries: 5
    environment:
      - "PGPASSWORD=${POSTGRES_PASS}"
      - "POSTGRES_PASSWORD=${POSTGRES_PASS}"
    ports:
      - "127.0.0.1:54321:5432"
    networks:
      - "api.network"
  ms-go:
    build:
      context: .
      dockerfile: ./build/dev/Dockerfile
    container_name: "ms-go"
    ports:
      - "127.0.0.1:8005:8005"
    volumes:
      - ./:/var/www
    links:
      - postgresql
    networks:
      - "api.network"
networks:
  api.network:`

// Env .env
const Env = `POSTGRES_PASS=default`

// GitIgnore .gitignore
const GitIgnore = `
/vendor
.idea
.env`

// Makefile Makefile
const Makefile = `#!make
include .env
export $(shell sed 's/=.*//' .env)

run: mod db app

mod:
	go mod vendor

db:
	docker-compose up -d --force-recreate --no-deps --build postgresql

app:
	docker-compose up -d --force-recreate --no-deps --build ms-go

swag:
	@swag init -g ./cmd/server/main.go  --parseDependency --parseInternal`

const Config = `apiserver_port: "8005"`

const DockerDev = `FROM golang:1.17.3-alpine

WORKDIR /app

COPY . .

RUN go build -o /app/cmd/server/apiserver -v ./cmd/server/main.go

RUN apk add --no-cache go

#ENTRYPOINT ["tail", "-f", "/dev/null"]

CMD ["/app/cmd/server/apiserver"]
`

const DockerProd = `FROM golang:1.17.3-alpine

WORKDIR /app

COPY . .

RUN go build -o /app/cmd/server/apiserver -v ./cmd/server/main.go

CMD ["/app/cmd/server/apiserver"]`
