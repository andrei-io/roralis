# syntax=docker/dockerfile:1
FROM golang:1.17-alpine AS build

# Install tools required for project

WORKDIR /usr/src/roralis/
COPY . .
# Install library dependencies
RUN go get

# Copy the entire project and build it
# This layer is rebuilt when a file changes in the project directory
RUN 	go build -o /bin/server main.go

# This results in a single layer image
FROM alpine:3.14
COPY --from=build /bin/server /bin/server
CMD ["/bin/server"]
