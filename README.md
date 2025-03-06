# Go HTMX Starter

A lightweight, modern starter template for building web applications with Go, HTMX, and Tailwind CSS.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![HTMX](https://img.shields.io/badge/HTMX-1.9.5-6366F1?style=for-the-badge&logo=html5&logoColor=white)
![Tailwind](https://img.shields.io/badge/Tailwind_CSS-4.0+-06B6D4?style=for-the-badge&logo=tailwind-css&logoColor=white)
![Templ](https://img.shields.io/badge/Templ-Latest-FF5A1F?style=for-the-badge&logo=go&logoColor=white)

This starter kit combines the power of Go's performance with the simplicity of HTMX for dynamic client interactions and Tailwind CSS for styling. Perfect for building modern web applications with minimal JavaScript while maintaining a great developer experience.

## Features

- 🚀 **Backend**: Lightweight Go HTTP server
- ⚡ **Frontend**: HTMX for dynamic content without writing JavaScript
- 🎨 **Styling**: Tailwind CSS for utility-first styling
- 🔥 **Hot Reload**: Air for Go code hot reloading during development
- 📝 **Templates**: [Templ](https://github.com/a-h/templ) for type-safe HTML templates in Go
- 🛠️ **Developer Experience**: Comprehensive Makefile for common tasks

## Prerequisites

- [Go](https://golang.org/dl/) (1.21 or later)
- [Node.js](https://nodejs.org/) and npm (for Tailwind CSS)

## Quick Start

### 1. Clone the repository

```bash
git clone https://github.com/amjadfqs/go-htmx-starter.git
cd go-htmx-starter
```

### 2. Install dependencies

```bash
make setup
```

This will install all necessary Go modules, Tailwind CSS, and development tools.

### 3. Start the development server

For the full development experience with hot reloading - You need to have `concurrently` package installed globally:

```bash
make dev-all
```

Or run individual components:

```bash
# Go server with hot reload
make dev

# Watch Tailwind CSS changes
make css-watch

# Generate templ templates
make generate
```

### 4. View the application

Open your browser and navigate to [http://localhost:8080](http://localhost:8080)

## Project Structure

```
go-htmx-starter/
├── cmd/
│   └── server/           # Main application entrypoint
├── internal/
│   ├── handlers/         # HTTP request handlers
│   └── models/           # Data models
├── static/
│   ├── css/              # CSS files (including Tailwind)
│   ├── js/               # JavaScript files
├── templates/
│   ├── components/       # Reusable UI components
│   └── pages/            # Full page templates
├── .air.toml             # Air configuration for hot reloading
├── .gitignore            # Git ignore file
├── go.mod                # Go module definition
├── go.sum                # Go module checksums
├── Makefile              # Development and build commands
└── README.md             # Project documentation
```

## Available Make Commands

| Command          | Description                                         |
| ---------------- | --------------------------------------------------- |
| `make dev`       | Start the development server with Air hot reloading |
| `make build`     | Build the application                               |
| `make run`       | Run the compiled application                        |
| `make css`       | Compile Tailwind CSS once                           |
| `make css-watch` | Watch and compile Tailwind CSS on changes           |
| `make generate`  | Generate templ templates                            |
| `make test`      | Run tests                                           |
| `make setup`     | Install project dependencies                        |
| `make dev-all`   | Run all development processes concurrently          |
| `make clean`     | Clean build artifacts                               |
| `make help`      | Show available commands                             |

## Adding New Pages/Components

1. Create a new templ file in `templates/pages/` or `templates/components/`
2. Generate the templates with `make generate`
3. Add a new handler in `internal/handlers/`
4. Register the route in `cmd/server/main.go`

## HTMX Example

This starter includes basic HTMX examples to get you started:

- Dynamic content loading with HTMX
- Client information display
- Simple toast notifications

## Deployment

To build for production:

```bash
make build
```

The compiled binary will be available in the `build/` directory.

## License

MIT

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.
