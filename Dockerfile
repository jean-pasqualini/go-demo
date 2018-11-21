FROM golang:1.8 AS build-env

WORKDIR /go/src/go-demo
RUN echo $GOPATH
RUN go get -u github.com/golang/dep/cmd/dep
COPY . .
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/app main.go


# final stage
FROM alpine
WORKDIR /
COPY --from=build-env /app/app /
ENTRYPOINT ./app