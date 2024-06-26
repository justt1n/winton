FROM golang:1.12.7-alpine3.10 AS build
# Support CGO and SSL
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/test ./main.go

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin/test /go/bin/test
EXPOSE 8080
ENTRYPOINT ["/go/bin/test", "--port", "8080"]