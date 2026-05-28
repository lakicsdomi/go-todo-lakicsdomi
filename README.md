# Go Todo Application

![Coverage](https://img.shields.io/badge/Coverage-79.8%25-brightgreen)
![Go Version](https://img.shields.io/badge/Go-1.26-blue)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED)

A lightweight Todo application built with Go, focused on clean project structure, containerized development, and reliable testing workflows.

The project was originally inspired by [ichtrojan/go-todo](https://github.com/ichtrojan/go-todo), but redesigned using my software architecture knowledge, integration testing, and automated CI/CD practices and it uses my own logger, Argus.

---

## Features

### Application

* MVC-style project structure
* Dependency injection for database access
* Bootstrap 5 frontend embedded directly into the binary using Go's `embed`
* CRUD task management
* Integrated Argus telemetry logging dashboard

### Infrastructure & DevOps

* Multi-stage Docker builds
* Docker Compose setup for app + MySQL
* Integration tests against a real database container
* GitHub Actions CI pipeline
* Automated coverage badge updates through Makefile scripts

---

## Prerequisites

* Go 1.26+
* Docker & Docker Compose
* Make

---

## Quick Start (Docker)

Clone the repository and start the full stack:

```bash id="7h2kq1"
git clone https://github.com/lakicsdomi/go-todo.git
cd go-todo

make docker-up
```

Services:

* App: `http://localhost:8080`
* Argus Dashboard: `http://localhost:9090`

Stop and clean up containers:

```bash id="m91x2d"
make docker-down
```

---

## Local Development

Create a `.env` file in the project root:

```env id="r8pn4v"
PORT=8080
DB_HOST=127.0.0.1
DB_USER=db_user
DB_PASS=db_password
```

Start your local MySQL server, then run:

```bash id="k0dz5t"
make run
```

---

## Testing & Code Quality

The project includes integration tests covering the full CRUD lifecycle against a real MySQL instance.

### Available Commands

| Command             | Description                                              |
| ------------------- | -------------------------------------------------------- |
| `make test`         | Run unit and integration tests with coverage             |
| `make docker-test`  | Run tests inside Docker with a temporary MySQL container |
| `make lint`         | Run `golangci-lint` locally                              |
| `make docker-lint`  | Run linting inside a container                           |
| `make update-badge` | Update the README coverage badge automatically           |

---

## CI/CD

The GitHub Actions workflow handles testing, linting, and coverage updates.

You can also simulate the pipeline locally using:

```bash id="c4u8y6"
act
```

This runs the workflow in a local Docker-based GitHub Actions environment before pushing changes.
