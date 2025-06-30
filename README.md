# URL Shortener

![Recording](./docs/recording.gif)

## Project Description

This project is a web application that allows users to shorten long URLs into compact, easy-to-share links. The system stores the mapping between the original and shortened URLs in a PostgreSQL database and automatically redirects users to the original URL when they access the short link.

The project is fully dockerized for easy local development and deployment. It also includes a CLI tool for managing database migrations.

## Goal

The main goal is to provide a simple and efficient tool for shortening URLs, making them more manageable for sharing on social media, messaging platforms, and other contexts where long links are inconvenient. The project also serves as an educational example of building a modern web application with robust validation and persistent storage.

## Technologies Used

- **Go (Golang):** Main programming language for the backend.
- **Gin:** Web framework for routing and HTTP controller management.
- **GORM:** ORM for database management using SQLite.
- **SQLite:** Lightweight embedded database for storing URLs.
- **HTML + Tailwind CSS:** For a modern, responsive user interface.
- **HTMX:** Enhances user experience with dynamic interactions without full page reloads.
- **SHA-256:** Hashing algorithm used to generate unique short codes from original URLs.

## Project Structure

- `/app`: Business logic, models, controllers, and routes.
- `/templates`: HTML files for the user interface.
- `/tools`: Tools for project management and development.
- `/migrations`: Database migration definitions.
- `/static`: Static assets (JS, CSS).
- `main.go`: Application entry point.
- `Dockerfile`: Multi-stage Docker build for development and production.
- `compose.yml`: Docker Compose file for orchestrating the app and PostgreSQL database.
- `go.mod` and `go.sum`: Dependency management files.

## Getting Started

### Clone the repository

```bash
git clone git@github.com:HectorZR/url-shortener.git
cd url-shortener
```

---

## Running with Docker

The recommended way to run the project is using Docker and Docker Compose. This will start both the application and a PostgreSQL database.

### 1. Copy and configure environment variables

```bash
cp .env.example .env
# Edit .env as needed (defaults work for local development)
```

### 2. Build and start the services

```bash
docker compose up
```

This will start:
- The URL shortener app on [http://localhost:8000](http://localhost:8000)
- A PostgreSQL database

### 3. Run database migrations

In a new terminal, run the migration CLI inside the app container:

```bash
docker compose exec app go run cli/cli.go migrate up
```

You can also run `migrate down` to rollback.

---

## Running Locally (without Docker)

Make sure you have [Go](https://go.dev/dl/) installed (version 1.24 or higher recommended) and a running PostgreSQL instance matching your `.env` configuration.

### 1. Install dependencies

```bash
go mod download
```

### 2. Start the project

```bash
go install github.com/air-verse/air@latest
air
```

The server will start on [http://localhost:8000](http://localhost:8000).

### 3. Run migrations

```bash
go run tools/cli.go migrate up
```

---

## CLI Tool for Database Migrations

A simple CLI tool is provided for managing database migrations.

**Usage:**

```bash
go run cli/cli.go migrate up    # Run all migrations
go run cli/cli.go migrate down  # Rollback all migrations
```

If using Docker, run these commands inside the app container:

```bash
docker compose exec app go run tools/cli.go migrate up
```

---


## Improvement Opportunities

- **User management and authentication:** Allow registered users to manage their own links.
- **Usage statistics:** Display how many times each short URL has been accessed.
- **Link expiration:** Enable short links to expire after a certain time or number of uses.
- **Custom short codes:** Allow users to choose their own short code.
- **Support for other databases:** Facilitate migration to systems like MySQL or SQLite for production environments.
- **Advanced collision handling:** Implement more sophisticated strategies to avoid short code collisions.
- **Cloud deployment:** Automate deployment to platforms like Heroku, AWS, or GCP.
- **Internationalization:** Support multiple languages in the user interface.

---

This project is an excellent starting point for learning about web development in Go, Docker-based deployment, database integration, and best practices for validation and security in web applications.
