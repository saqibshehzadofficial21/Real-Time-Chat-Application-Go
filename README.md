<div align="center">

# Chat Application Backend

**A production-grade real-time messaging platform built with Go and PostgreSQL.**

</div>

---

## Overview

This project is a fully functional chat application backend designed to support real-time one-to-one and group messaging. It implements a clean three-layer architecture with robust authentication, WebSocket-based live communication, and comprehensive database design for scalable conversations.

The system handles user registration, friend request workflows, conversation management, message persistence, and real-time message delivery through persistent WebSocket connections with automatic heartbeat management.

---

## Features

### Authentication & Authorization
- Secure user registration with bcrypt password hashing
- JWT-based stateless authentication with configurable expiry
- Token validation middleware protecting all private endpoints
- Participant authorization checks preventing unauthorized message access

### Messaging
- One-to-one conversations created automatically upon friend request acceptance
- Group conversations with multiple participants
- Persistent message storage with PostgreSQL
- Pagination support for message history retrieval
- Read status tracking per message

### Real-Time Communication
- WebSocket connections for live message delivery
- Room-based message broadcasting (conversation-scoped)
- Automatic ping/pong heartbeat to maintain long-lived connections
- Concurrent client handling via goroutines

### Friend Management
- Send and receive friend requests
- Pending request inbox for receivers
- Automatic conversation creation upon acceptance
- Duplicate request prevention

---

## Tech Stack

| Category | Technology | Purpose |
|----------|-----------|---------|
| Language | Go 1.21+ | High-performance backend runtime |
| Framework | Gin | HTTP web framework and routing |
| ORM | GORM | Database abstraction and migrations |
| Database | PostgreSQL | Relational data persistence |
| Authentication | JWT (golang-jwt/jwt/v5) | Stateless token-based auth |
| Password Security | bcrypt | Adaptive password hashing |
| Real-Time | Gorilla WebSocket | Bidirectional persistent connections |
| Configuration | godotenv | Environment variable management |

---

## System Architecture

The application follows a strict three-layer architecture to ensure separation of concerns, testability, and maintainability.

```
    Client (Postman / Web / Mobile)
              |
              v
    +---------------------------+
    |   Handler Layer           |
    |   (Presentation)          |
    |   - Parse HTTP requests   |
    |   - Validate input        |
    |   - Return JSON responses |
    +------------+--------------+
                 |
                 v
    +---------------------------+
    |   Service Layer           |
    |   (Business Logic)        |
    |   - Enforce rules         |
    |   - Authorization checks  |
    |   - Orchestrate data flow |
    +------------+--------------+
                 |
                 v
    +---------------------------+
    |   Repository Layer        |
    |   (Data Access)           |
    |   - GORM operations       |
    |   - Query optimization    |
    |   - Connection pooling    |
    +------------+--------------+
                 |
                 v
    +---------------------------+
    |   PostgreSQL Database     |
    |   - Users                 |
    |   - Conversations         |
    |   - Messages              |
    |   - Friend Requests       |
    +---------------------------+
```

### Layer Responsibilities

**Handler Layer**
- Receives HTTP requests and extracts parameters
- Performs input validation using struct tags
- Delegates processing to the Service layer
- Formats and returns HTTP responses with appropriate status codes

**Service Layer**
- Contains all business logic and domain rules
- Validates authorization (e.g., participant checks before message access)
- Handles password hashing and JWT generation
- Coordinates multiple repository operations for complex workflows

**Repository Layer**
- Executes all database operations via GORM
- Defines interfaces for each domain entity
- Manages query construction and relationship preloading
- Provides clean data access abstraction for the Service layer

---

## Installation & Setup

### Prerequisites

- Go 1.21 or later installed
- PostgreSQL 14+ running locally or via Docker
- Git for version control

### Step 1: Clone the Repository

```bash
git clone https://github.com/yourusername/chat-app.git
cd chat-app
```

### Step 2: Install Dependencies

```bash
go mod download
```

### Step 3: Start PostgreSQL

Using Docker:

```bash
docker run --name chat-db \
  -e POSTGRES_PASSWORD=yourpassword \
  -e POSTGRES_DB=chatapp \
  -p 5432:5432 \
  -d postgres
```

Or use an existing PostgreSQL instance and create the `chatapp` database manually.

### Step 4: Configure Environment Variables

Copy the example file and update the values:

```bash
cp .env.example .env
```

Edit `.env` with your credentials:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=chatapp
DB_SSLMODE=disable

JWT_SECRET=your-super-secret-jwt-key-minimum-thirty-two-characters
JWT_EXPIRY_HOURS=24

SERVER_PORT=8080
SERVER_MODE=debug

WS_READ_BUFFER_SIZE=1024
WS_WRITE_BUFFER_SIZE=1024
WS_PING_INTERVAL=30
WS_PONG_WAIT=60
```

### Step 5: Run the Application

```bash
go run cmd/server/main.go
```

The server will start on `http://localhost:8080` with automatic database migration.

---

## API Endpoints

### Public Endpoints

| Method | Route | Description |
|--------|-------|-------------|
| POST | `/api/register` | Create a new user account |
| POST | `/api/login` | Authenticate and receive JWT token |

### Protected Endpoints (Bearer Token Required)

| Method | Route | Description |
|--------|-------|-------------|
| POST | `/api/friend-requests` | Send a friend request |
| GET | `/api/friend-requests/pending` | List pending friend requests |
| POST | `/api/friend-requests/:id/accept` | Accept a friend request |
| GET | `/api/conversations` | List all conversations for the user |
| GET | `/api/conversations/:id` | Get details of a specific conversation |
| POST | `/api/groups` | Create a new group conversation |
| POST | `/api/messages` | Send a message to a conversation |
| GET | `/api/conversations/:id/messages` | Retrieve messages with pagination |
| GET | `/api/ws` | Establish WebSocket connection |

### Authentication Header Format

```
Authorization: Bearer <jwt_token>
```

---

## Project Structure

```
chat-app/
|
|-- cmd/
|   -- server/
|       -- main.go                    Application entry point
|
|-- internal/
|   |
|   |-- config/
|   |   -- config.go                  Environment configuration loader
|   |
|   |-- db/
|   |   -- postgres.go                Database connection and auto-migration
|   |
|   |-- models/
|   |   -- user.go                    User entity definition
|   |   -- conversation.go            Conversation and participant entities
|   |   -- message.go                 Message entity definition
|   |   -- friend_request.go          Friend request entity definition
|   |
|   |-- repository/
|   |   -- user_repository.go         User data access interface
|   |   -- friend_repository.go       Friend request data access interface
|   |   -- message_repository.go      Message data access interface
|   |   -- conversation_repository.go Conversation data access interface
|   |
|   |-- service/
|   |   -- auth_service.go            Authentication business logic
|   |   -- friend_service.go          Friend request business logic
|   |   -- message_service.go         Message business logic
|   |   -- conversation_service.go    Conversation business logic
|   |
|   |-- handlers/
|   |   -- auth_handler.go            HTTP handlers for authentication
|   |   -- friend_handler.go          HTTP handlers for friend requests
|   |   -- message_handler.go         HTTP handlers for messages
|   |   -- conversation_handler.go    HTTP handlers for conversations
|   |
|   |-- middleware/
|   |   -- auth_middleware.go         JWT validation middleware
|   |
|   |-- routes/
|   |   -- routes.go                  Route registration and grouping
|   |
|   -- websocket/
|       -- hub.go                     WebSocket room and client management
|       -- handler.go                 WebSocket connection upgrade handler
|
|-- pkg/
|   -- utils/
|       -- jwt.go                     JWT utility functions
|       -- response.go                Standardized HTTP response helpers
|
|-- .env.example                      Environment variable template
|-- .gitignore
|-- go.mod
|-- go.sum
```

### Directory Descriptions

- `cmd/server/` — Contains the main application entry point. This is the only package that should produce an executable.
- `internal/` — Private application code. Subpackages are organized by architectural layer.
- `internal/models/` — Pure data structures representing database tables. No business logic.
- `internal/repository/` — Data access layer with interface definitions for each entity.
- `internal/service/` — Business logic layer enforcing rules and authorization.
- `internal/handlers/` — HTTP presentation layer handling request/response cycles.
- `internal/middleware/` — Cross-cutting concerns such as authentication.
- `internal/websocket/` — Real-time communication infrastructure.
- `pkg/utils/` — Shared utility functions used across the application.

---

## Database Schema

### Users Table

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Conversations Table

```sql
CREATE TABLE conversations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    is_group BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Conversation Participants (Junction Table)

```sql
CREATE TABLE conversation_participants (
    conversation_id INT REFERENCES conversations(id) ON DELETE CASCADE,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (conversation_id, user_id)
);
```

### Messages Table

```sql
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    conversation_id INT NOT NULL REFERENCES conversations(id) ON DELETE CASCADE,
    sender_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_messages_conversation_id ON messages(conversation_id);
CREATE INDEX idx_messages_created_at ON messages(created_at);
```

### Friend Requests Table

```sql
CREATE TABLE friend_requests (
    id SERIAL PRIMARY KEY,
    sender_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    receiver_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status VARCHAR(20) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## UI/UX & Design Approach

While this repository contains the backend implementation, the API design and response structure are crafted with frontend consumption in mind:

- **Consistent Response Format**: All API responses follow a uniform JSON structure with clear `success`, `data`, and `error` fields.
- **Semantic Status Codes**: HTTP status codes are used precisely — `201 Created` for new resources, `403 Forbidden` for authorization failures, `401 Unauthorized` for authentication issues.
- **Error Clarity**: Error messages are descriptive yet concise, enabling frontend developers to display meaningful feedback to users.
- **Pagination Support**: Message endpoints support `limit` and `offset` parameters for efficient data loading in chat interfaces.
- **WebSocket Integration**: The real-time layer is designed for seamless frontend integration with standard JSON message formats and automatic reconnection support via heartbeat.
- **Relationship Preloading**: API responses include related user data (sender info, participant lists) to minimize frontend request chaining.

---

## Deployment

### Docker

A `Dockerfile` can be added to containerize the Go application:

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
```

Build and run:

```bash
docker build -t chat-app .
docker run -p 8080:8080 --env-file .env chat-app
```

### Kubernetes

For production deployments, the application can be deployed to a Kubernetes cluster with:

- A Deployment managing the Go application pods
- A Service exposing the application internally
- An Ingress handling external traffic and TLS termination
- A PostgreSQL StatefulSet or managed cloud database for persistence
- ConfigMaps and Secrets for environment configuration

---

## Future Improvements

- **Message Reactions**: Add emoji reactions to individual messages
- **Message Editing & Deletion**: Allow users to modify or remove sent messages within a time window
- **File Attachments**: Support image, video, and document uploads via object storage (S3/MinIO)
- **Push Notifications**: Integrate Firebase Cloud Messaging or APNs for offline notifications
- **Message Search**: Implement full-text search across conversation history
- **Typing Indicators**: Real-time typing status via WebSocket broadcasts
- **Read Receipts**: Per-user read tracking with timestamps
- **End-to-End Encryption**: Implement signal-protocol style encryption for sensitive conversations
- **Rate Limiting**: Add per-user request throttling to prevent abuse
- **Structured Logging**: Migrate to zap or logrus for production-grade observability
- **Metrics & Monitoring**: Integrate Prometheus metrics and health check endpoints
- **Horizontal Scaling**: Support multiple server instances with Redis-backed WebSocket broadcasting

---

## Author

**Your Name**

Full-Stack Developer specializing in backend systems, real-time applications, and cloud-native architectures.

- GitHub: [github.com/yourusername](https://github.com/yourusername)
- LinkedIn: [linkedin.com/in/yourprofile](https://linkedin.com/in/yourprofile)
- Email: your.email@example.com

---

<div align="center">

*Built with precision. Designed for scale.*

</div>
