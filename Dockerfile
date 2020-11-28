FROM golang:buster as build

WORKDIR /app

COPY . .

RUN ls

RUN CGO_ENABLED=0 GOOS=linux go build \
    -mod=vendor -ldflags "-s -w" -a -installsuffix cgo -o bin .

# Build image with binary
FROM scratch

WORKDIR /root/
COPY --from=build /app/bin .
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["./bin"]

EXPOSE 8081
