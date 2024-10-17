# Microservices Project

This project is a collection of microservices designed for handling **Orders**, **Products**, and **User Management**. Each service has its own database migrations, gRPC and REST APIs, and distinct configurations. The services communicate via gRPC for internal interactions.

## Project Structure

```
├── docker-compose.yaml         # Docker configuration for running services
├── microservices.collection.postman_collection.json  # Postman collection for testing APIs
├── orders                      # Order service handling order-related logic
├── product                     # Product service for managing products
├── user                        # User service managing authentication and user profiles
```

### Services Overview

1. **Orders Service** (`/orders`)
   - Manages order creation and status.
   - Exposes both REST and gRPC APIs.
   - Handles database migrations and data models related to orders.

2. **Product Service** (`/product`)
   - Handles product CRUD operations.
   - Supports both REST and gRPC.
   - Includes database migrations for product management.

3. **User Service** (`/user`)
   - Manages user authentication, profiles, and JWT generation.
   - Provides REST and gRPC APIs for user registration and login.
   - Implements utility functions like password hashing and JWT validation.

## Technologies Used

- **Go**: Main programming language for each microservice.
- **gRPC**: For inter-service communication.
- **REST**: Exposed APIs for external consumption.
- **Docker**: Containerization of microservices.
- **MySQL**: Database used for persistence.
- **Migration Tools**: SQL migrations for setting up databases.

## How to Run

1. **Build and Run Services**:
   - Use Docker Compose to spin up the services:
     ```bash
     docker-compose up --build
     ```

2. **API Documentation**:
   - Use the Postman collection `microservices.collection.postman_collection.json` for testing the APIs.

3. **Migrations**:
   - Each service has its own migrations located in the `migrations/` folder. Run the migrations using the provided migration scripts.
