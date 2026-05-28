# 📝 Go Todo Application

![Coverage](https://img.shields.io/badge/Coverage-79.8%25-brightgreen)
![Go Version](https://img.shields.io/badge/Go-1.26-blue)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED)
![Architecture](https://img.shields.io/badge/Architecture-MVC-orange)

A lightweight, fully functional ToDo application built with Go. This project serves as a showcase of **Clean Architecture**, **Dependency Injection**, and **Enterprise-grade DevOps practices**.

Inspired by [ichtrojan/go-todo](https://github.com/ichtrojan/go-todo), but completely re-engineered from the ground up with focus on testability, containerization, and automated CI/CD workflows.

---

## ✨ Key Features

### 🏗 Software Architecture
- **MVC Pattern**: Strict separation of concerns between Routing, Controllers, and Models.
- **Dependency Injection**: Database connections are injected into controllers, preventing global state mutation and allowing for flawless integration testing.
- **Embedded Frontend**: Uses Go's `embed` package to compile Bootstrap 5 HTML views directly into a single, portable binary.
- **Custom Telemetry**: Integrated with the modular `argus` logging system, providing a standalone HTTP telemetry dashboard.

### 🚀 DevOps & Infrastructure
- **Multi-Stage Docker Builds**: Segregated stages for base dependencies, testing, compiling, and building the final minimal production image (using Alpine).
- **Isolated Integration Testing**: A dedicated `docker-test` environment that spins up a throwaway MySQL instance, waits for healthchecks, and runs full lifecycle integration tests against a real database.
- **Resilient CI/CD Pipeline**: GitHub Actions workflow featuring automated coverage badge generation via `Makefile` `sed` scripting, and a self-healing linting mechanism (Action with Docker fallback).

---

## 🛠 Prerequisites

- **Go 1.26+** (For local development)
- **Docker & Docker Compose** (Recommended for isolated execution)
- **Make** (For build automation)

---

## 🚀 Quick Start (Docker - Recommended)

The easiest way to run the application is via Docker. This spins up a configured MySQL database and the minimal Go production container.

```bash
git clone [https://github.com/lakicsdomi/go-todo.git](https://github.com/lakicsdomi/go-todo.git)
cd go-todo

# Start the full stack (App + Database)
make docker-up
```

- **App Interface**: http://localhost:8080

- **Argus Telemetry Dashboard**: http://localhost:9090

To stop the containers and clean up volumes:
```bash
make docker-down
```

## 💻 Local Development
1. Create a `.env` file in the root directory (based on `.env.example`):

    ```ini,toml
    PORT=8080
    DB_HOST=127.0.0.1
    DB_USER=db_user
    DB_PASS=db_password
    ```
2. Start your local MySQL server.
3. Run the application:
   ```bash
    make run
   ```

## 🧪 Testing & Code Quality

This project takes testing seriously, featuring full CRUD integration tests without hardcoded endpoints. The `Makefile` serves as the central hub for all operations.

### Makefile Commands

| Command | Description |
| --- | --- |
| `make test` | Runs local unit and integration tests with coverage. |
| `make docker-test` | **[CI Standard]** Runs integration tests inside an isolated Docker network alongside a throwaway DB. Safely extracts `coverage.out` using `docker cp`. |
| `make lint` | Runs `golangci-lint` locally. |
| `make docker-lint` | Runs strict static analysis in an isolated container. |
| `make update-badge` | Dynamically parses `coverage.out` and updates this README's badge using `sed` (No 3rd party actions required!). |

### CI/CD Simulation

To validate the `.github/workflows/ci.yml` pipeline locally before pushing, this repository supports [act](https://github.com/nektos/act).
Run `act` in the root directory to simulate the exact GitHub Actions runner environment.

