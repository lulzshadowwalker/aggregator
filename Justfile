set dotenv-load := true

DB_URL       := env_var("DATABASE_URL")
MIGRATIONS_DIR := "internal/database/migrations"

default:
    @just --list

migrate-up:
    goose --dir {{MIGRATIONS_DIR}} postgres "{{DB_URL}}" up

migrate-down:
    goose --dir {{MIGRATIONS_DIR}} postgres "{{DB_URL}}" down

migrate-status:
    goose --dir {{MIGRATIONS_DIR}} postgres "{{DB_URL}}" status

migrate-reset:
    goose --dir {{MIGRATIONS_DIR}} postgres "{{DB_URL}}" reset

migrate-fresh:
    goose --dir {{MIGRATIONS_DIR}} postgres "{{DB_URL}}" reset
    goose --dir {{MIGRATIONS_DIR}} postgres "{{DB_URL}}" up

migrate-create name:
    goose --dir {{MIGRATIONS_DIR}} create {{name}} sql
