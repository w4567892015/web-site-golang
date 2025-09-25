FROM busybox AS box
ARG BUSYBOX_VERSION=1.31.0-i686-uclibc

FROM golang:latest AS base
ARG appName=""
WORKDIR /usr/src/
COPY go.work go.work.sum ./
COPY ./apps/${appName} ./apps/${appName}
COPY ./libs ./libs
RUN sed -i "/^\t\.\/apps\// { /${appName}/!d }" go.work

FROM base AS builder
RUN go mod download
RUN go build -ldflags="-linkmode external -extldflags -static" -o main ./apps/${appName}

FROM alpine:latest
COPY --from=box /bin/wget /usr/bin/wget
COPY --from=builder /usr/src/main /main

EXPOSE 3000
CMD ["./main"]
