# ddb go api demo

This is a demo project using ddb to show how to setup a go htpp server with it.

DISCLAIMER : I am not using traefik in my ddb configuration. I bind ports with docker-compose.

## Introduction

At your disposal is a REST api and an exemple of SOAP api as well.

Regarding the db, you can connect to an existing one or use the docker one. Bear in mind that the migration system is not automatized, therefore you need to add your migrations directly into : the database.Migrate func.

To use a postgres db instead, use this package :

```
go get -u github.com/lib/pq
```

If you want to use an ORM which will handle migrations and drivers for you, you can get GORM : https://github.com/go-gorm/gorm

There are also web frameworks like GIN : https://github.com/gin-gonic/gin

## Installation

### Clone the project

```
https://github.com/inetum-orleans/ddb-go-api-demo.git
```

### Environment variables (Optionnal)

Modify the .env.jinja to your need. Do use the same API port as the one specified in ddb.yml : app.api.port. Or redo the binding in the docker-compose.yml accordingly.

### Configure ddb

```
ddb configure
```

### Build and launch the containers

```
docker compose --build -d
```

### Start the API in dev mode

Hot reloading is done with air-verse: https://github.com/air-verse/air. This feature will watch the files in your project and rebuild the binary every time a change is detected.

To start the Go server with air, run:

```
air
```

or

```
make watch
```

To build the binary :

```
make build
```
