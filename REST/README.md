# demo-golang-api-with-mongodb

## Overview

This application is built using Go and MongoDB and utilizes the Gin web framework. It allows you to perform CRUD operations on books.

## Requirements

- Go (Version 1.15 or higher)
- MongoDB

## Setup

### Environment Variables

This application uses environment variables to handle configurations. To set these up:

1. Create a .env file in the root directory of the application.
2. Add your MongoDB connection string as follows:

```bash
MONGO_URI=your_actual_mongo_uri_here
```

## Running the application

After setting up the environment variables:

1. Open a terminal and navigate to the root directory of the application.
2. Run the following commands:

```bash
go mod download
go run main.go
```

The application should now be running, and you should be connected to your MongoDB database.

### **Using Swagger to Test API**

To test the API using Swagger:

**Navigate to Swagger UI**: Open your web browser and go to **`http://localhost:8080/swagger/index.html`**.
