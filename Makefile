SHELL := /bin/bash

.PHONY: test

test:
	@set -euo pipefail; \
	cleanup() { \
		echo "==> Limpiando contenedores y volúmenes..."; \
		docker compose down -v --remove-orphans >/dev/null 2>&1 || true; \
	}; \
	trap cleanup EXIT; \
	echo "==> Generando código con sqlc..."; \
	go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.31.1 generate; \
	echo "==> Compilando proyecto..."; \
	go build ./...; \
	echo "==> Eliminando entorno de pruebas anterior..."; \
	docker compose down -v --remove-orphans; \
	echo "==> Levantando PostgreSQL..."; \
	docker compose up -d; \
	echo "==> Esperando a que PostgreSQL esté disponible..."; \
	ready=0; \
	for i in $$(seq 1 30); do \
		if docker compose exec -T db pg_isready -U postgres -d tpe_pw_test >/dev/null 2>&1; then \
			ready=1; \
			break; \
		fi; \
		sleep 1; \
	done; \
	if [ "$$ready" -ne 1 ]; then \
		echo "ERROR: PostgreSQL no estuvo disponible a tiempo."; \
		docker compose logs db; \
		exit 1; \
	fi; \
	echo "==> Ejecutando tests..."; \
	TEST_DATABASE_URL="postgres://postgres:postgres@localhost:5433/tpe_pw_test?sslmode=disable" \
		go test -v ./...
