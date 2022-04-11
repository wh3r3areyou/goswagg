# Goswagg - generate template simple golang project with GIN based on swagger documentation

Goswagg -  this is a project that allows you to create an application infrastructure for working with REST API 
based on swagger documentation, and it remains only to implement the business logic of the application
in controllers, repositories and requests (requests to third-party applications)

It is based on [go-swagger][go-swagger] which is used for the code generation.

## Install  

You can simply run `go get` to install goswagg.

```
$ go get github.com/wh3r3areyou/goswagg
$ goswagg --help
usage: goswagg generate --file=SWAGGER.YAML

generate app from swagger.yaml

Usage:
  generate [flags]

Flags:
  -f, --file string   path to swagger.yaml
  -h, --help          help for generate

```

## How it works

Goswagg generates a [gin][gin] based service given a swagger spec.

Take example swagger docs: https://petstore.swagger.io/v2/swagger.yaml
```
You can generate the REST API using the following command:

$ goswagger -f=swagger.yaml
```


## Features
* [ ] Health checkers
* [ ] Middlewares
* [ ] Test controllers
* [ ] Refactoring service code and package :)

[gin]: https://github.com/gin-gonic/gin
[go-swagger]: https://github.com/go-swagger/go-swagger