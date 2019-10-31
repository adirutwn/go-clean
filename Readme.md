## Go Clean

A simple clean architecture implemented in Golang.  
### Folder Structure
```
|-- app
|   |-- All Go code will be keep under this folder
|   |-- entities
|   |   |-- `entities` is where we keep business object
|   |   |-- most of this are struct that we pass around each layer
|   |-- environments
|   |   |-- `environments` is where we keep env variables
|   |-- extensions
|   |   |-- `extensions` is where initiate the connection to 3rd party services
|   |-- modules
|   |   |-- `modules` is where we separate domain
|   |   |-- users
|   |   |   |-- deliveries
|   |   |   |   |-- `deliveries` is where all end-user of this service will interact with
|   |   |   |   |-- can be http/grpc/cmd and etc.
|   |   |   |   `-- http
|   |   |   |   `-- ...
|   |   |   |-- repositories
|   |   |   |   |-- `repositories` is where we keep the code that connects
|   |   |   |   |-- other services not in the app (i.e. API, DB, etc.)
|   |   |   |`-- usecases
|   |   |   |   |-- `usecases` is where business logic at.
|   |   |-- repository.go -> Interface of repository
|   |   |-- usecase.go -> Interface of usecase
|   |-- test_helpers
|   |   |-- a package for storing helper functions to do unit-test
|   |-- utils
|   |   |-- `utils` keeps all util functions
|   `-- main.go -> this is where we glue each layer together
|-- resources
|   `-- docker
    |   `-- docker-compose.yaml
|-- Dockerfile
`-- Makefile
|-- README.md
|-- go.mod
`-- go.sum
```
### Folder Structure
```
// spinning the service up locally
$ make docker/local/up

// stop the service
$ make docker/local/down

// running all unit-test
$ make test/run

// generate mock
$ make mockgen module=user interface=Repository

// create migration file
$ make migration/create name=create_some_table
```
