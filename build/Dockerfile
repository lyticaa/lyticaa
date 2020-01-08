#
# Build
#
FROM golang:1.13 AS build

WORKDIR /app

COPY . .

ENV GOBIN=/go/bin
ENV GOPATH=/go
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go mod verify
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o $GOBIN/dashd ./cmd/dashd
RUN go build -a -installsuffix cgo -o $GOBIN/workerd ./cmd/workerd

#
# Main
#
FROM alpine

# Copy the compiled binaires over.
COPY --from=build /go/bin/dashd /usr/bin/
COPY --from=build /go/bin/workerd /usr/bin/

EXPOSE 3000
