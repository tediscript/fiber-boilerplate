# Fiber Boilerplate with Air Hot Reloading

This is a boilerplate for a Go Fiber application with Air hot reloading configured.

## Features

- [Fiber](https://github.com/gofiber/fiber) - Fast HTTP web framework for Go
- [Air](https://github.com/cosmtrek/air) - Live reload for Go applications
- Environment variables support via .env file
- Health check endpoints (/livez and /readyz)
- Metrics dashboard endpoint (/metrics)
- Graceful shutdown

## Requirements

- Go 1.18 or higher
- Air (for hot reloading)

## Getting Started

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/fiber-boilerplate.git
   cd fiber-boilerplate
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

### Development

To run the application with hot reloading:

```bash
make dev
```

Or directly with Air:

```bash
air
```

Or using the shell script:

```bash
./dev.sh
```

This will start the application and automatically reload it whenever you make changes to the source code.

### Production

To build and run the application for production:

```bash
make build
./fiber-boilerplate
```

### Other Commands

The project includes a Makefile with several useful commands:

```bash
make help      # Show available commands
make dev       # Run with hot reloading
make build     # Build the application
make run       # Run without hot reloading
make clean     # Clean build artifacts
make test      # Run tests
```

## Configuration

The application can be configured using environment variables in the `.env` file:

- `PORT` - The port on which the server will listen (default: 3000)
- `APP_NAME` - The name of the application (default: "Fiber Boilerplate")
- `SHUTDOWN_TIMEOUT` - Graceful shutdown timeout in seconds (default: 5)

## Air Configuration

Air is configured in the `.air.toml` file. You can modify this file to change how Air watches and rebuilds your application.

Key settings:

- `include_ext` - File extensions to watch for changes
- `exclude_dir` - Directories to exclude from watching
- `exclude_file` - Files to exclude from watching
- `delay` - Delay in milliseconds before rebuilding after a change

For more information, see the [Air documentation](https://github.com/cosmtrek/air).
