# Basic Golang App

Run the project using the following command
```bash
go run . --config-file=./config.yml
```

## Unit Tests

For testing we use <https://github.com/golang/mock>
```bash
GO111MODULE=on go get github.com/golang/mock/mockgen@latest
```

Generate mocks for unit test by running
```bash
go generate ./...
```

Run unit tests
```bash
go test -v ./...
```

## Code Linting

TBD

## Docker Build

We use docker to build this service

Build the image
```bash
docker build -t charlires/go-app .
```

Push image to docket hub
```bash
docker push charlires/go-app
```

Run container
```bash
docker run --rm -it -v `pwd`/config.yml:/etc/config/config.yml charlires/go-app
```

## Deploy App

### Run app in docker-desktop kubernetes

Generate Kubernetes Deployment files
```bash
kustomize build ./k8s/base > ./k8s/out/base.yml
```

```bash
kubectl config use-context docker-desktop
```

```bash
kubectl apply -f ./k8s/out/base.yml
```
