# Go Todo App

Inspired by [https://github.com/ichtrojan/go-todo](https://github.com/ichtrojan/go-todo)

A lightweight, fully functional ToDo application built with Go, implementing Clean Architecture principles and a Bootstrap 5 responsive UI.

## Features
- **MVC Architecture**: Clean separation between routing, controllers, and models.
- **MySQL Integration**: Persistent storage using the standard `database/sql` library.
- **Dynamic Theming**: Built-in Light/Dark mode toggling via Bootstrap 5.
- **Custom Logging**: Integrated with the modular `argus` logging system for robust error tracking.

## Prerequisites
- Go 1.13 or higher
- MySQL server (configured with `db-user` account)

## Installation & Setup

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/lakicsdomi/go-todo.git](https://github.com/lakicsdomi/go-todo.git)
   cd go-todo

2. Configure Environment
Create a `.env` file in the root directory based on `.env.example`:
```ini
PORT=8080
DB_HOST=127.0.0.1
DB_USER=db-user
DB_PASS=your_password
```

3. Run the Application:
```bash
go run main.go
```

The server will start on http://localhost:8080 and automatically configure the required database and tables upon the first run.