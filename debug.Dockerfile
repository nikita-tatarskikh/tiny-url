# Compile stage
FROM golang:1.22 AS build-env

# Build Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

ADD . /dockerdev
WORKDIR /dockerdev

RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o /tiny

# Final stage
FROM debian:buster

EXPOSE 8080 40000

WORKDIR /
COPY --from=build-env /go/bin/dlv /
COPY --from=build-env /tiny /

CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/tiny"]