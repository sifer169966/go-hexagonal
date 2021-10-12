# Defining App builder image
FROM golang:alpine AS builder

ARG VERSION=unversioned
ARG PRIVATE_KEY

# Add git to determine build git version
RUN apk add --no-cache --update git openssh

# Set GOPATH to build Go app
ENV GOPATH=/go

# Set apps source directory
ENV SRC_DIR=${GOPATH}/src/github.com/hex-gonal-template

# Define current working directory
WORKDIR ${SRC_DIR}

# Copy apps scource code to the image
COPY . ${SRC_DIR}

# Setup environment
RUN ["chmod", "+x","./setup.sh"]

# Build App
RUN go mod verify && \
go mod vendor && \
go fmt ./... && \
go install -v -ldflags="-X ${REPO_PATH}/version.Version=${VERSION}" ./cmd/...

# Defining App image
FROM alpine:latest

RUN apk add --no-cache --update ca-certificates

# Copy App binary to image
COPY --from=builder /go/src/github.com/hex-gonal-template/config/*.yaml /app/config/
COPY --from=builder /go/bin/cmd /usr/local/bin/cmd

ENTRYPOINT ["cmd"]