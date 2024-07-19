## Golang RestAPI with GIN 
Golang Boilerplate with Gin framework following The Clean Architecture.

This project has 4 Domain layer to follow the clean architecture principles. 
  - Models layer
  - Business layer
  - Data layer
  - Prensentation layer

The models layer will contains all database models. In the business layer, the business logic will be implemented and it have interface to all functions that will be used in Presentation layer. The data layer it have interface to all function that need to be used to interact with the database. The presentasion layer it will contain all teh function that will be exposed the feature of application in form of REST API. 


## How To Use

### Prerequisites
- Go 1.22+
- AIR (https://github.com/cosmtrek/air)
- PostgreSQL 

### Installation
 
Clone the repository
```bash
git clone https://github.com/ansxy/golang-boilerplate-gin.git
```

Go to the project folder
```bash
cd golang-boilerplate-gin
```

Install the dependencies
```bash
go mod tidy
```

Build the project
```bash
go build -o app cmd/main.go
```

Run the project
```bash
./app
```

### Run the project

To run the project, you can use the following command:

```bash
go run cmd/main.go
```

For Development purpose, you can use the following command on the project folder:

```bash
air
```

> Note if you are using Windows, you need to change the .air.toml file to use the correct path for the executable file.
> Note gracefully shutdown doesnt work on hot reload AIR 


<!-- ## Architecture

![Architecture](https://github.com/joaopaulovieira/golang-gin-boilerplate/blob/master/images/architecture.png) -->

## Folders

### `/cmd`
 This folder contains the application entry point. main.go is the entry point of the application. App folder is the main package of the application and contains the configuration of the application.

### `/config`
 This folder contains the configuration of the application. The configuration is loaded from the environment variables. The configuration will be used in the application to configure the application such as the database connection, the port, etc.

### `/internal`
 This folder will contain all the packages in application such as the business layer, data layer and presentation layer. 
 
 - `/internal/handler`
 This folder contains all the logic that will be used for expose the feature of the application in form of REST API. Handler will handle the request such as validation, authentication, authorization and so on and it will return the response to the client.

 - `/internal/model`
 This folder contains all the models that will be used in the application. 

 - `/internal/middleware`
 This folder contains all the middleware that will be used in the application such as authentication, authorization, logging, etc.

 - `/internal/repository`
 This folder contains all the repositories that will be used in the application.

 - `/internal/usecase`
 This folder contains all the usecases/business logic that will be used in the application.

 - `/internal/transport`
 This folder contains what kind of transport will be used in the application. For example, HTTP, gRPC, etc. In this case, it will be HTTP.

 - `/internal/request`
 This folder contains an interface that will be used to validate the request.

 - `/internal/response`
 This folder contains an interface that will be used to return the response.

### `/pkg`
 This folder contains all the packages that are used in the application. For any framework/library that is used in the application, it will be in this folder.



## References
- [go-clean-arch](https://github.com/bxcodec/go-clean-arch?tab=readme-ov-file)
- [go-clean-architecture](https://github.com/khannedy/golang-clean-architecture) 
