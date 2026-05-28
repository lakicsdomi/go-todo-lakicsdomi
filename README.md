![Coverage](https://img.shields.io/badge/Coverage-0%25-brightgreen)

# Go Todo App

Inspired by [https://github.com/ichtrojan/go-todo](https://github.com/ichtrojan/go-todo)

A lightweight, fully functional ToDo application built with Go, implementing Clean Architecture principles and a Bootstrap 5 responsive UI.

## Features
- **MVC Architecture**: Clean separation between routing, controllers, and models.
- **MySQL Integration**: Persistent storage using the standard `database/sql` library.
- **Dependency Injection**: Database connections are injected to prevent global state, ensuring perfect testability.
- **Dynamic Theming**: Built-in Light/Dark mode toggling via Bootstrap 5.
- **Custom Logging**: Integrated with the modular `argus` logging system and its standalone telemetry dashboard.
- **Embedded Frontend**: Uses Go's `embed` package to compile HTML views directly into the binary.

## Prerequisites
- Go 1.26 or higher
- Docker & Docker Compose (Recommended)
- Make (Optional, for build automation)
- MySQL server (if running bare-metal)

## Installation & Setup (Docker - Recommended)

The easiest way to run the application is via Docker. This will spin up a MySQL database and a minimal, multi-stage built Go container.

```bash
git clone [https://github.com/lakicsdomi/go-todo.git](https://github.com/lakicsdomi/go-todo.git)
cd go-todo
make docker-up
```

The app will be available at http://localhost:8080, and the Argus telemetry dashboard at http://localhost:9090.
To stop the containers, run `make docker-down`.

## Installation & Setup (Local/Bare-metal)
1. Clone the repository.

2. Create a .env file in the root directory based on .env.example:
    ```ini, toml
    PORT=8080
    DB_HOST=127.0.0.1
    DB_USER=db_user
    DB_PASS=db_password
    ```

3. Run the Application
    ```bash
    make run
    ```

## Testing & CI/CD

This repository utilizes **GitHub Actions** for Continuous Integration. The pipeline is designed to be agnostic, relying purely on the `Makefile` targets.
- **Run tests locally**: `make test`
- **Local CI Validation**: To test the `.github/workflows/ci.yml` file locally before pushing, it is highly recommended to use [act](https://github.com/nektos/act). Simply run `act` in the root directory to spin up an isolated Docker container that simulates the GitHub Actions runner.