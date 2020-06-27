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

Run unit tests `go test -v ./...`

## Code Linting

TBD

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

## Deploy App

### Run app in docker-desktop kubernetes

Generate Kubernetes Deployment files  
```
kubectl kustomize ./k8s/base > ./k8s/out/base.yml
```

```
kubectl config use-context docker-desktop
```

```
kubectl apply -f ./k8s/out/base.yml
```
