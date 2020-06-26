# Basic Golang App

Run the project using the following command
```
go run . --config-file=./config.yml
```

## Unit Tests

For testing we use https://github.com/golang/mock   
```
GO111MODULE=on go get github.com/golang/mock/mockgen@latest
```

Generate mocks for unit test by running `go generate ./...`

Run unit tests `go test ./...`

## Docker Build

We use docker to build this service

Build the image
```
docker build -t charlires/go-app .
```

Push image to docket hub
```
docker push charlires/go-app
```

Run container 
```
docker run --rm -it -v `pwd`/config.yml:/etc/config/config.yml charlires/go-app 
```

