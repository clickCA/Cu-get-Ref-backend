FROM golang:1.21.1-alpine3.18

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o course-management-service.exe main.go

# Expose port 8080 to the outside world
EXPOSE 8080

ENTRYPOINT ["/app/course-management-service.exe"]

