# Architecture Documentation

This project follows a Hexagonal Architecture (also known as Ports and Adapters) pattern, which promotes separation of concerns and maintainability through clear layer boundaries.

## Project Structure

The project is organized into three main layers:

### 1. Application Layer (`api/application`)

The application layer acts as an orchestrator between the domain and infrastructure layers. It:
- Contains use cases and application services
- Orchestrates the flow of data between the outer and inner layers
- Implements application-specific business rules
- Handles the transformation of data between layers

### 2. Domain Layer (`api/domain`)

The domain layer is the heart of the application, containing:
- Business entities and value objects
- Domain interfaces (ports)
- Business rules and logic
- Domain events and event handlers
- Repository interfaces

This layer is completely independent of external concerns and frameworks.

### 3. Infrastructure Layer (`api/infrastructure`)

The infrastructure layer contains all external-facing components and implementations:
- Database implementations
- External service integrations
- API controllers
- Repository implementations

## Key Principles

1. **Dependency Rule**: Dependencies only point inwards. The domain layer has no dependencies on outer layers.
2. **Isolation**: The domain logic is isolated from external concerns.
3. **Adaptability**: External components can be easily replaced without affecting the core business logic.
4. **Testability**: Business logic can be tested without external dependencies.

## Flow of Control

1. External requests come through the infrastructure layer
2. The application layer orchestrates the use case
3. Domain logic is executed in the domain layer
4. Results flow back through the application layer
5. Infrastructure layer handles the final response

## Interface Adapters

- **Primary (Driving) Adapters**: REST controllers, gRPC handlers, CLI commands
- **Secondary (Driven) Adapters**: Database repositories, external service clients, cache implementations

This architecture ensures that our business logic remains clean and independent of external concerns while providing flexibility to change external implementations without affecting the core business rules.