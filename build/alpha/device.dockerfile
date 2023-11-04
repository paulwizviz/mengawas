
ARG GO_VER

FROM golang:${GO_VER} as builder 

ARG APP_NAME

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./go.sum ./go.sum
COPY ./go.mod ./go.mod

RUN go mod download && \
    env CGO_ENABLED=0 env GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./build/${APP_NAME} ./cmd/ex1/device

FROM ubuntu:18.04

ARG APP_NAME

COPY --from=builder /opt/build/${APP_NAME} /${APP_NAME}

CMD /${APP_NAME}