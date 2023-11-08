FROM golang:1.21.1-alpine3.18

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && \
    apk add --no-cache make

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Environment variables
ENV DB_HOST=course_db
ENV DB_PORT=3306
ENV DB_NAME=courseservice_db
ENV DB_USER=root
ENV DB_PASSWORD=password
ENV SERVER_PORT=8080

# Create .env file
RUN echo "DB_CONNECTION=${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?charset=utf8mb4&parseTime=True&loc=Local" >> .env
RUN echo "SERVER_PORT=${SERVER_PORT}" >> .env

# make generate_grpc_code and Build the Go app
RUN make build

# Expose port 8080 to the outside world
EXPOSE 8080

ENTRYPOINT ["/app/course-management-service.exe"]


