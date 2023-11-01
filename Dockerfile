# Defining App builder image
FROM golang:1.17-alpine3.14 AS builder


# Define current working directory
WORKDIR /app

RUN apk update; \
    apk add --no-cache \
    git \
    make
# turn on go module
ENV GO111MODULE on

# Setup build environments
# CGO_ENABLED=0
# for cross-compilation is because we should use the go
# built-in support of the target platform cross-compilation
# and there should be no reason not to do so.
# GOOS=linux
# use the linux operation system for build
# - GOARCH=arm64
# use the arm64 as the CPU architecture
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=arm64

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
COPY go.mod .
COPY go.sum .

COPY . ./

# Build App
RUN make ci && make build

# Defining App image
FROM alpine:3.14 as release

RUN apk add --no-cache --update ca-certificates \
    make

# Copy App binary to image
COPY --from=builder /app/go-hex /app/cmd/

RUN chmod +x /app/cmd/go-hex

ARG NONROOT_GROUP=nonroot-group
ARG NONROOT_USER=nonroot-user
ARG USER_ID=20000

RUN addgroup -S $NONROOT_GROUP && adduser -S -u $USER_ID $NONROOT_USER -G $NONROOT_GROUP

USER $NONROOT_USER:$NONROOT_GROUP

WORKDIR /app

CMD ["cmd/go-hex", "serve-rest"]