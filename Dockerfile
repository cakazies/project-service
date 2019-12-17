FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
# Create appuser.
RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/github.com/crowdeco/project-service
COPY . .
# Using go get.
RUN go get

# buildin apps in project-service 
RUN go build -o project-service

# running project-service
ENTRYPOINT ./project-service

# running in port 
EXPOSE 8003
