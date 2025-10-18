# Social API

**Description:**  
Social API is a RESTful backend service built with Go, designed for practicing and learning backend development concepts. It provides endpoints for managing users, posts, and social interactions, serving as a foundation for experimenting with:

- Go web development
- Authentication
- Database integration

This project is ideal for developers looking to deepen their understanding of Go and modern API design.

## Modules

The project is organized into several key modules:

- **api**: Handles HTTP routing, request validation, and response formatting.
- **domain**: Contains core business entities and interfaces, defining the application's business rules.
- **repository**: Implements the Repository Pattern, providing abstractions and concrete implementations for data access (e.g., database operations).
- **service**: Encapsulates business logic, orchestrating interactions between repositories and other modules.
- **config**: Manages application configuration, such as environment variables and settings.
- **middleware**: Provides reusable middleware components for authentication, logging, and error handling.
- **db**: Handles database connection setup and migrations.
- **utils**: Contains utility functions and helpers used across the project.

This modular structure supports clean separation of concerns and makes the codebase easier to test and maintain.