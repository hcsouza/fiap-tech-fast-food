# fiap-tech-fast-food


### Tech

This api was built using [Golang](https://golang.org/) and some tools:
 * [Gin](http://github.com/gin-gonic/gin) - Web framework 
 * [Mongo-Driver](http://go.mongodb.org/mongo-driver) - driver to deal with MongoDB
 * [viper](https://github.com/spf13/viper) - Config solution tool
 * [mockery](https://github.com/vektra/mockery) - Mock tool to use on unit tests

## Run

The app can be started using docker and you can use the actions pre-defineds on Makefile

* ***Build image***

To build an image from project to push to a registry you can use the command below:

```sh
make build-image
```
this command will generate an image with this tag: *fiap-tech-fast-food*

If you want run this image locally you can use this command: 

```sh
docker run -v /host/configs.yaml:/app/data/configs/config.yaml -p 8080:8080 -it fiap-tech-fast-food
```

An sample of configs.yaml could be found on: 

`internal/adapter/infra/config/configs.yaml.sample`


* ***Generate docs***

To generate the documentation to publish on project like an openApi you can use the command below:

```sh
make serve-swagger
```
this command will generate a directory called `docs` 


### Development

To run in development for debug or improvement you can use another command:

```sh
make start-local-development
``` 

this command will start a container with hot-reload to any modification on the code. Including a container with an instance of MongoDB.

To stop the container execute:

```sh
make stop-local-development
```

### Test

Locally you can use the command below:

```sh
go test ./...  -v
```

or use a make action: 

```sh
make test   
```