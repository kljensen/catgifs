FROM golang:1.21-alpine

# create a working directory
WORKDIR /app
# add source code
COPY go.mod server.go ./
COPY images ./images

# build the source
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -installsuffix cgo -o main

# run the binary
ENTRYPOINT [ "./main" ]
