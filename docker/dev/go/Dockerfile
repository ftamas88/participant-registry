FROM golang:1.14-alpine as build

COPY . /source
WORKDIR /source

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o ./bin/app -v ./cmd/app/app.go
FROM alpine:latest

COPY --from=build /source/bin/app /app

ENTRYPOINT ["/app"]
