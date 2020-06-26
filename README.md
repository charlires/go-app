# Basic Golang App

Run the project using the following command
```
go run . --config-file=./config.yml
```

For testing we use
gomock
```
github.com/golang/mock/mockgen@latest
```

Generate mocks for unit test run `go generate ./...`

Run unit tests `go test ./...`