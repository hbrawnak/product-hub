## Overview
ProductHub is a RESTful API for managing product information, built using Go (Golang). It leverages Docker for containerized deployment and includes support for MySQL as the database backend. The API exposes endpoints for creating, reading, updating, and deleting (CRUD) product records, and organizes routes for better scalability and maintainability.

## Features
* RESTful API implementation with dedicated routes for products.
* Structured and modular codebase.
* Dockerized for consistent and portable deployments.
* Includes a Makefile for streamlined development and testing tasks.
* MySQL integration for persistent data storage.

## Project Structure
```
.
├── main.go            # Application entry point
├── routes.go          # Handles routing and route grouping
├── Makefile           # Automates build, run, and test tasks
├── go.mod             # Go module dependencies
├── docker-compose.yml # Docker configuration for services
└── README.md          # Project documentation
```

```
git clone https://github.com/your-username/producthub.git
cd producthub
```

## Start the services
```
make up         # Use Docker Compose to start MySQL and other required services
make run        # Run the application Build and run the application
make build      # Build and run locally 
make run-build  # Build and run locally 

```

## API Endpoints
```
Method  Endpoint	        Description
------  --------                -----------
GET	/	                Get API home page
GET	/api/products	        List all products
GET	/api/products/{id}	Get product by ID
POST    /api/products	        Create a new product
PUT	/api/products/{id}	Update an existing product
DELETE	/api/products/{id}	Delete a product
```

## Author
Developed by [Md Habibur Rahman](https://habib.im).