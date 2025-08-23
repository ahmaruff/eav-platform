# EAV Platform - Development Roadmap

A progressive Entity-Attribute-Value CRUD application built with Go,
focusing on dynamic data modeling and multi-tenant architecture.

## Tech Stack

- **HTTP Router**: Chi
- **Database**: SQLx + database/sql (SQLite -> PostgreSQL/MySQL)
- **Query Builder**: Squirrel
- **Migrations**: golang-migrate
- **Frontend**: Templ + HTMX + Tailwind CSS
- **Session**: SCS (server-side sessions)
- **Logging**: slog (structured logging)

## Development Philosophy

Each version builds incrementally on the previous one, ensuring solid foundations before adding complexity.
Focus on learning core patterns while building something production-ready.

---

## 🚀 Version 1: Foundation & Authentication

**Goal**: Establish robust authentication system with proper error handling and middleware patterns.

### Core Features
- User registration and login system
- Password hashing with bcrypt
- Server-side session management
- Protected route middleware
- Comprehensive error handling
- Structured logging

### Database Schema
```sql
users (
  id, email, password_hash, 
  created_at, updated_at
)
```

### Deliverables
- ✅ User registration/login forms (Templ + HTMX)
- ✅ Session-based authentication
- ✅ Auth middleware for protected routes
- ✅ Error handling patterns
- ✅ Structured logging with slog
- ✅ Database migrations setup
- ✅ Basic dashboard after login

**Success Criteria**: Secure user auth flow with proper error handling and logging.

---

## 🔧 Version 2: EAV CRUD System

**Goal**: Implement dynamic Entity-Attribute-Value system for flexible data modeling.

### Core Features
- Dynamic entity creation
- Field management with data types
- Record CRUD operations
- Auto-generated UI forms
- Query optimization for EAV patterns

### Database Schema
```sql
entities (
  id, name, created_by, 
  created_at, updated_at
)

fields (
  id, entity_id, name, field_type, 
  required, default_value, 
  created_at, updated_at
)

values (
  id, entity_id, record_id, field_id, 
  value_text, value_number, value_date, value_boolean,
  created_at, updated_at
)
```

### Deliverables
- ✅ Entity management interface
- ✅ Dynamic field creation (text, number, date, boolean)
- ✅ Record CRUD with auto-generated forms
- ✅ Data validation and type conversion
- ✅ Basic search and filtering
- ✅ Performance optimization for EAV queries

**Success Criteria**: Admin can create entities/fields, users can manage records dynamically.

---

## 🏢 Version 3: Multi-Tenant & RBAC

**Goal**: Scale to multi-tenant SaaS with role-based access control.

### Core Features
- Tenant isolation
- Role and permission system
- Multi-user per tenant
- Tenant-scoped data access
- Advanced security middleware

### Database Schema
```sql
tenants (
  id, name, slug, settings,
  created_at, updated_at
)

-- Updated users table
users (
  id, tenant_id, email, password_hash,
  created_at, updated_at
)

roles (
  id, tenant_id, name, description,
  created_at, updated_at
)

permissions (
  id, code, description, resource, action
)

role_permissions (
  role_id, permission_id
)

user_roles (
  user_id, role_id
)
```

### Deliverables
- ✅ Tenant registration and management
- ✅ User invitation system
- ✅ Role/permission management UI
- ✅ Tenant isolation middleware
- ✅ Permission-based route protection
- ✅ Tenant switching interface
- ✅ Data export/import per tenant

**Success Criteria**: Multiple tenants can operate independently with proper access controls.

---

## 🔮 Future Versions (V4+)

### Potential Features
- **API & Webhooks**: REST/GraphQL API with webhook system
- **Advanced UI**: Rich text fields, file uploads, drag-drop forms
- **Analytics**: Usage metrics, custom dashboards
- **Integrations**: Third-party API connections
- **Mobile**: React Native or PWA
- **Enterprise**: SSO, audit logs, compliance features

---

## Development Guidelines

### Code Organization (Hybrid DDD Approach)
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
│           ├── user_sqlite.go
│           └── session_memory.go
├── migrations/          # Database migrations
├── templates/           # Templ templates
├── static/              # CSS, JS, assets
├── config/              # Configuration
└── docs/               # Documentation
```

### Quality Standards
- **Testing**: Unit tests for business logic, integration tests for handlers
- **Documentation**: Clear README, API docs, code comments
- **Security**: Input validation, SQL injection prevention, XSS protection
- **Performance**: Database indexing, query optimization, caching strategies
- **Logging**: Structured logs with correlation IDs
- **Error Handling**: Consistent error types and user-friendly messages

### Migration Strategy
- **V1→V2**: Add EAV tables, migrate user data
- **V2→V3**: Add tenant/RBAC tables, scope existing data to default tenant
- Each version should support graceful upgrades without data loss
