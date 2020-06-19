# Start from golang base image (using alpine linux)
FROM golang:alpine as build

# Add Info
LABEL maintainer="João Marcos <joaomarcoslopes@id.uff.br>"

# System updates and Git installation
# Git is required for fetching the dependecies
RUN apk update && apk add --no-cache git

# Change the working directory
WORKDIR /src

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependecies and cache it if files are not changed 
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .