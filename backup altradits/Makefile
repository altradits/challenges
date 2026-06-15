.PHONY: run build dev setup-db

run:
	go run cmd/server/main.go

build:
	go build -o bin/altradits cmd/server/main.go

dev:
	DB_URL=postgres://altradits:password@localhost:5432/altradits?sslmode=disable go run cmd/server/main.go

setup-db:
	psql -c "CREATE USER altradits WITH PASSWORD 'password';" || true
	psql -c "CREATE DATABASE altradits OWNER altradits;" || true
	psql -c "GRANT ALL PRIVILEGES ON DATABASE altradits TO altradits;" || true

clean:
	rm -rf bin/
