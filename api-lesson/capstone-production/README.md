# Capstone: Production Deployment

## What You'll Build

A production-ready deployment setup for the Task Manager API.

## Features

- Docker containerization
- Environment configuration
- Structured logging
- Health checks
- Rate limiting
- HTTPS/TLS setup
- Monitoring and metrics

## Project Structure

```
capstone-production/
├── Dockerfile              # Container definition
├── docker-compose.yml      # Multi-container setup
├── .env.example          # Environment variables
├── nginx/
│   └── nginx.conf        # Reverse proxy config
├── scripts/
│   ├── migrate.sh        # Database migration
│   └── seed.sh           # Seed data
├── monitoring/
│   ├── prometheus.yml    # Prometheus config
│   └── grafana.json      # Dashboard
└── deployment/
    ├── kubernetes/       # K8s manifests
    └── systemd/          # Systemd service
```

## Docker Setup

### Dockerfile

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o taskmanager .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/taskmanager .
COPY --from=builder /app/.env.example .env
EXPOSE 8080
CMD ["./taskmanager"]
```

### docker-compose.yml

```yaml
version: '3.8'
services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DB_PATH=/data/app.db
      - JWT_SECRET=${JWT_SECRET}
    volumes:
      - data:/data
    depends_on:
      - redis
  
  redis:
    image: redis:alpine
    ports:
      - "6379:6379"
  
  nginx:
    image: nginx:alpine
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx/nginx.conf:/etc/nginx/nginx.conf
    depends_on:
      - api

volumes:
  data:
```

## Environment Configuration

### .env

```
# Server
PORT=8080
ENV=production

# Database
DB_PATH=/data/app.db

# JWT
JWT_SECRET=your-secret-key-here
JWT_EXPIRATION=24h

# Redis (for rate limiting)
REDIS_ADDR=redis:6379

# Logging
LOG_LEVEL=info
LOG_FORMAT=json
```

## Health Checks

### HTTP Health Check

```go
// GET /health
// Returns: {"status": "ok"}
```

### Database Health Check

```go
// Check database connection
// Return 503 if database is unavailable
```

## Rate Limiting

### Redis-based Rate Limiter

```go
// 100 requests per minute per IP
// Return 429 if limit exceeded
```

## HTTPS/TLS

### Using Let's Encrypt

```bash
# Get certificate
certbot certonly --nginx -d yourdomain.com

# Auto-renewal
certbot renew --dry-run
```

### Self-signed Certificate (Development)

```bash
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365
```

## Monitoring

### Prometheus Metrics

```go
// /metrics endpoint
// Expose:
// - Request count
// - Request duration
// - Error count
// - Active connections
```

### Grafana Dashboard

- Request rate
- Error rate
- Response time
- Database connections

## Deployment Options

### 1. Docker Compose (Simple)

```bash
docker-compose up -d
```

### 2. Kubernetes

```bash
kubectl apply -f deployment/kubernetes/
```

### 3. Systemd

```bash
sudo cp deployment/systemd/taskmanager.service /etc/systemd/system/
sudo systemctl enable taskmanager
sudo systemctl start taskmanager
```

## Security Considerations

- Use HTTPS in production
- Set secure cookie flags
- Implement rate limiting
- Use strong JWT secrets
- Keep dependencies updated
- Run as non-root user

## Skills Applied

- Docker and containerization
- Environment configuration
- Security best practices
- Monitoring and observability
- Production deployment patterns