# AGENTS.md

This file provides guidance for AI coding agents working on this GoFiber web application.

## About This Project

A Go web application built with:
- **GoFiber** — High-performance web framework
- **html/template** — Go's built-in template engine via Fiber's template adapter
- **Server-side rendering** — HTML templates in `/views`

## Core Principles

- **Idiomatic Go**: Follow "Effective Go" conventions
- **Performance**: Avoid unnecessary memory allocations in handlers
- **Clarity**: Write clear, concise code with comments for complex logic
- **Error Handling**: Handle all errors with meaningful messages using the `errors` package

## GoFiber Guidelines

### Handlers and fiber.Ctx

- **Never store `fiber.Ctx` references** outside the handler's scope — it's reused. Copy data to new variables if needed.
- **Return responses** with `c.Render()`, `c.JSON()`, `c.Send...()`, etc. Only use `c.Next()` in middleware.
- **Bind and validate** request bodies using `c.BodyParser()` and `go-playground/validator`.

### Middleware

- Use official Fiber middleware for logging, panic recovery, and CORS
- Custom middleware goes in `middlewares/` and must call `c.Next()`

### Routing

- Define routes in the `routes/` package, not `main.go`
- Use `app.Group()` to organize related endpoints

### Template Engine Setup

```go
package main

import (
    "log"
    "your_project_name/routes"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2"
)

func main() {
    engine := html.New("./views", ".html")
    app := fiber.New(fiber.Config{
        Views: engine,
    })
    routes.SetupRoutes(app)
    log.Fatal(app.Listen(":3000"))
}
```

### Route Definition Example

```go
package routes

import (
    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("pages/home", fiber.Map{
            "Title": "Hello, World!",
        }, "layouts/main")
    })

    api := app.Group("/api")
    api.Get("/users", handlers.GetUsers)
}
```

## Project Structure

```
/
├── main.go             # Entry point: init template engine, database, Fiber app
├── routes/             # Route definitions mapping URLs to handlers
├── handlers/           # HTTP handlers for request processing
├── store/              # Database interaction logic
├── models/             # Data structs (User, Product, etc.)
├── views/              # HTML templates (pages/, layouts/)
├── middlewares/        # Custom middleware functions
├── static/             # Static assets (css/, js/)
├── configs/            # Configuration loading
├── .env                # Environment variables
├── go.mod / go.sum
└── Makefile            # Development automation
```

## Tooling

### Makefile Commands

Prefer these make targets:

| Command | Purpose |
|---------|---------|
| `make dev` | Run with hot-reloading |
| `make run` | Run normally |
| `make build` | Compile binary |
| `make test` | Run unit tests |
| `make clean` | Remove build artifacts |

### Code Quality

- Run `go fmt` on all generated/modified code
- Code must pass `golangci-lint` defaults (errcheck, govet, staticcheck)
- Manage dependencies via Go modules; provide `go get` commands for new deps
