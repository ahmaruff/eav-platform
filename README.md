# EAV Platform

A progressive Entity-Attribute-Value CRUD application built with Go,
focusing on dynamic data modeling and multi-tenant architecture.

## Current Status

- ✅ **V1**: User authentication with session management
- 📋 **V2**: Dynamic EAV system (planned)
- 📋 **V3**: Multi-tenant with RBAC (planned)

## Tech Stack

- **HTTP Router**: Chi
- **Database**: SQLx + database/sql (SQLite → PostgreSQL/MySQL)
- **Query Builder**: Squirrel
- **Migrations**: golang-migrate
- **Frontend**: Templ + HTMX + Tailwind CSS (standalone)
- **Session**: SCS (server-side sessions)
- **Logging**: slog (structured logging)

## Quick Start

### Prerequisites

- Go 1.24+
- Make (optional but recommended)

### Development Setup

1. **Clone and setup**:
   ```bash
   git clone <your-repo>
   cd eav-platform
   make setup  # Install dev tools
   ```

2. **Run in development**:
   ```bash
   make dev    # Auto-reload on changes
   ```

3. **Or run manually**:
   ```bash
   make build && make run
   ```

4. **Visit**: http://localhost:8080

## Available Commands

```bash
make help        # Show all commands
make build       # Build application
make run         # Run application  
make dev         # Development mode (auto-reload)
make test        # Run tests
make clean       # Clean build artifacts
make migrate-up  # Run database migrations
make generate    # Generate templ templates
```

## Project Structure (Hybrid DDD)

```
eav-platform/
├── cmd/server/           # Main application entry point
├── internal/
│   ├── user/            # User domain
│   │   ├── model.go     # User entity & value objects
│   │   ├── repository.go # Repository interface
│   │   ├── service.go   # Domain service
│   │   └── handler.go   # HTTP handlers
│   ├── auth/            # Authentication domain
│   │   ├── service.go   # Auth domain service
│   │   ├── session.go   # Session management
│   │   ├── middleware.go # Auth middleware
│   │   └── handler.go   # Auth handlers
│   ├── shared/          # Shared kernel
│   │   ├── errors.go    # Common error types
│   │   ├── validator.go # Input validation
│   │   ├── logger.go    # Logging utilities
│   │   └── database.go  # DB connection & utilities
│   └── infrastructure/  # Infrastructure layer
│       └── repository/  # Concrete repository implementations
├── migrations/          # Database migrations
├── templates/           # Templ templates
├── static/              # Tailwind CSS, assets
├── config/              # Configuration
└── docs/               # Documentation (ROADMAP.md, REQUIREMENTS.md)
```

## V1 Features

### Authentication Flow
1. **Register**: Create account with email/password
2. **Login**: Authenticate with credentials  
3. **Session**: Server-side session with secure cookies
4. **Protected Routes**: Dashboard requires authentication
5. **Logout**: Destroy session and redirect

### Security Features
- ✅ bcrypt password hashing
- ✅ Secure session cookies (httpOnly, sameSite)
- ✅ Input validation and sanitization
- ✅ SQL injection prevention (parameterized queries)
- ✅ XSS protection (proper templating)
- ✅ Structured error handling

### UI/UX
- ✅ Tailwind CSS for styling (no Node.js required)
- ✅ HTMX for dynamic interactions
- ✅ Progressive enhancement
- ✅ Mobile-responsive design
- ✅ Loading states and error messages

## Development

### Database
- SQLite for development (easy setup)
- Migrations with golang-migrate
- SQLx for type-safe database operations
- Prepared statements prevent SQL injection

### Frontend Architecture
- **Templ**: Type-safe HTML templates in Go
- **HTMX**: Dynamic interactions without JavaScript frameworks
- **Tailwind**: Utility-first CSS

### Error Handling
- Consistent error types (validation, auth, database, internal)
- Structured logging with correlation IDs
- User-friendly error messages
- Graceful degradation

## Roadmap

See `ROADMAP.md` for detailed development plan:
- **V1**: Foundation & Authentication ✅
- **V2**: EAV CRUD System 📋
- **V3**: Multi-tenant + RBAC 📋

## Requirements

See `REQUIREMENTS_V1.md` for V1 detailed specifications including:
- Functional requirements
- Non-functional requirements
- API endpoints
- Database schema
- Testing requirements

## Contributing

1. Follow the roadmap and requirements documents
2. Maintain hybrid DDD structure
3. Write tests for business logic
4. Use conventional commits
5. Update documentation

## License

MIT License - see [LICENSE](./LICENSE) file for details.
