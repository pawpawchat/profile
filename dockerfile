FROM golang:1.22 AS build

WORKDIR /profile

COPY . .

RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main cmd/main.go

FROM alpine:latest

COPY --from=build /profile/main /main

CMD ["/main"]
