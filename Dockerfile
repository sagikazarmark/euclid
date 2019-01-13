ARG GO_VERSION=1.11

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk add --update --no-cache ca-certificates make git curl mercurial

RUN mkdir /build
WORKDIR /build

COPY go.mod go.sum /build/
RUN go mod download

COPY . /build
RUN BUILD_DIR= BINARY_NAME=app make build-release


FROM alpine:3.8

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /app /app

EXPOSE 8000 10000
CMD ["/app", "--instrumentation.addr", ":10000", "--app.addr", ":8000"]
