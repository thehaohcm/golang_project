FROM golang:alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Add Maintainer Info
LABEL maintainer="Hao Nguyen <hao.nguyen@s3corp.com.vn>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

WORKDIR /app/api/cmd/golang_project/

# Build the Go app
RUN go build -o /app/api/cmd/golang_project/main

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]