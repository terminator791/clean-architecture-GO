# Clean Architecture Go - Product Management API

A production-ready Go application implementing Clean Architecture principles for a simple Product Management API with CRUD operations.

## 🏗️ Architecture

This project follows the Clean Architecture pattern with clear separation of concerns:

```
cmd/
└── server/           # Application entry point
internal/
├── domain/           # Enterprise Business Rules
│   ├── entity/       # Entities
│   └── repository/   # Repository interfaces
├── usecase/          # Application Business Rules
├── adapter/          # Interface Adapters
│   ├── controller/   # HTTP handlers
│   └── repository/   # Repository implementations
└── infrastructure/   # Frameworks & Drivers
    ├── database/     # Database connection
    └── router/       # HTTP routing
```

## 🚀 Features

- **Clean Architecture**: Strict adherence to Clean Architecture principles
- **CRUD Operations**: Complete Create, Read, Update, Delete functionality for products
- **PostgreSQL**: Production-ready database with connection pooling
- **Database Migrations**: Structured SQL migrations for database schema
- **Configuration Management**: Environment-based configuration with Viper
- **UUID Support**: Google UUID for entity identifiers
- **Containerization**: Docker and Docker Compose for easy deployment
- **Health Checks**: Built-in health check endpoints
- **Graceful Shutdown**: Proper server shutdown handling
- **CORS Support**: Cross-Origin Resource Sharing enabled

## 🛠️ Technology Stack

- **Language**: Go 1.21+
- **Database**: PostgreSQL
- **Database Access**: `jmoiron/sqlx`
- **Configuration**: `spf13/viper`
- **UUIDs**: `google/uuid`
- **Migrations**: `golang-migrate/migrate` convention
- **Containerization**: Docker & Docker Compose

## 📋 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| POST | `/products` | Create a new product |
| GET | `/products` | Get all products (with pagination) |
| GET | `/products/{id}` | Get product by ID |
| PUT | `/products/{id}` | Update product |
| DELETE | `/products/{id}` | Delete product |

## 🚀 Quick Start

### Using Docker Compose (Recommended)

1. **Clone the repository**
   ```bash
   git clone https://github.com/terminator791/clean-architecture-GO.git
   cd clean-architecture-GO
   ```

2. **Start the application**
   ```bash
   docker-compose up --build
   ```

3. **The API will be available at**: `http://localhost:8080`

### Manual Setup

1. **Prerequisites**
   - Go 1.21+
   - PostgreSQL 15+

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up PostgreSQL database**
   ```sql
   CREATE DATABASE products;
   ```

4. **Run migrations**
   ```bash
   # Using golang-migrate tool
   migrate -path migrations -database "postgres://postgres:password@localhost:5432/products?sslmode=disable" up
   ```

5. **Set environment variables**
   ```bash
   export APP_DB_HOST=localhost
   export APP_DB_PORT=5432
   export APP_DB_USER=postgres
   export APP_DB_PASSWORD=password
   export APP_DB_NAME=products
   export APP_DB_SSLMODE=disable
   export APP_SERVER_PORT=8080
   ```

6. **Run the application**
   ```bash
   go run cmd/server/main.go
   ```

## 📝 API Usage Examples

### Create a Product
```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "iPhone 15",
    "description": "Latest iPhone model",
    "price": 999.99
  }'
```

### Get All Products
```bash
curl http://localhost:8080/products
```

### Get All Products with Pagination
```bash
curl "http://localhost:8080/products?limit=10&offset=0"
```

### Get Product by ID
```bash
curl http://localhost:8080/products/{product-id}
```

### Update Product
```bash
curl -X PUT http://localhost:8080/products/{product-id} \
  -H "Content-Type: application/json" \
  -d '{
    "name": "iPhone 15 Pro",
    "description": "Pro version of iPhone 15",
    "price": 1199.99
  }'
```

### Delete Product
```bash
curl -X DELETE http://localhost:8080/products/{product-id}
```

### Health Check
```bash
curl http://localhost:8080/health
```

## ⚙️ Configuration

The application supports configuration through environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `APP_SERVER_PORT` | `8080` | Server port |
| `APP_DB_HOST` | `localhost` | Database host |
| `APP_DB_PORT` | `5432` | Database port |
| `APP_DB_USER` | `postgres` | Database user |
| `APP_DB_PASSWORD` | `password` | Database password |
| `APP_DB_NAME` | `products` | Database name |
| `APP_DB_SSLMODE` | `disable` | Database SSL mode |

## 🗄️ Database Schema

```sql
CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL CHECK (price >= 0),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
```

## 🧪 Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

## 🏗️ Build

```bash
# Build the application
go build -o bin/server cmd/server/main.go

# Run the built binary
./bin/server
```

## 🐳 Docker

### Build Docker Image
```bash
docker build -t clean-architecture-go .
```

### Run with Docker
```bash
docker run -p 8080:8080 \
  -e APP_DB_HOST=your-db-host \
  -e APP_DB_PASSWORD=your-password \
  clean-architecture-go
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Clean Architecture Resources

- [The Clean Architecture by Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Clean Architecture in Go](https://github.com/bxcodec/go-clean-arch)