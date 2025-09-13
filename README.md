# Go Scaffold Tool

A powerful command-line tool for scaffolding new Go web projects with a clean, organized structure following Go best practices.

## 🚀 Features

- **Clean Architecture**: Creates projects with proper separation of concerns
- **Go Best Practices**: Follows standard Go project layout conventions
- **Web-Ready**: Includes basic HTTP server setup and configuration
- **Dependency Management**: Automatically initializes Go modules
- **Cross-Platform**: Works on Windows, macOS, and Linux

## 📁 Project Structure Created

When you run the scaffold, it creates the following structure:

```
your-project-name/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── api/                 # API handlers and controllers
│   ├── service/             # Business logic layer
│   ├── repository/          # Data access layer
│   └── domain/              # Core business models
├── pkg/                     # Public libraries and utilities
├── configs/
│   └── config.yaml          # Configuration files
├── web/
│   └── templates/           # HTML templates
├── scripts/                 # Build and deployment scripts
├── .gitignore               # Git ignore rules
├── README.md                # Project documentation
└── go.mod                   # Go module file
```

## 🛠️ Installation

### Prerequisites
- Go 1.19 or later
- Git

### Build from Source

1. **Clone or download this project**
2. **Navigate to the project directory**
3. **Build the executable**:
   ```bash
   go build -o bin/go_scaffold ./cmd
   ```

### Add to PATH (Windows)

To use the scaffold from any directory, add it to your PATH:

```powershell
# For current session
$env:PATH += ";path\to\go_scaffold\bin"

# Permanently (run in PowerShell as Administrator)
setx PATH "%PATH%;path\to\go_scaffold\bin"
```

### Add to PATH (macOS/Linux)

```bash
# Add to your shell profile (.bashrc, .zshrc, etc.)
export PATH="$PATH:path/to/go_scaffold/bin"

# Or create a symlink to a directory in PATH
sudo ln -s path/to/go_scaffold/bin/go_scaffold /usr/local/bin/go_scaffold
```

## 📖 Usage

### Basic Usage

```bash
# Create a new Go project
go_scaffold my-awesome-project

# This creates a directory called 'my-awesome-project' with the full structure
```

### Example Workflow

```bash
# 1. Create new project
go_scaffold my-web-app

# 2. Navigate to the new project
cd my-web-app

# 3. Install dependencies
go mod tidy

# 4. Run the application
go run cmd/main.go
```

The scaffold will create:
- A basic HTTP server listening on port 8080
- Configuration file with database and server settings
- Proper Go module initialization
- Git ignore file with common Go exclusions

## ⚙️ Configuration

The scaffold creates a `configs/config.yaml` file with default settings:

```yaml
server:
  port: 8080

database:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "password"
  dbname: "yourdb"
```

Modify these settings according to your needs.

## 🏗️ Architecture Overview

The scaffold follows a clean architecture pattern:

- **`cmd/`**: Application entry points
- **`internal/`**: Private application code
  - `api/`: HTTP handlers and routing
  - `service/`: Business logic
  - `repository/`: Data persistence layer
  - `domain/`: Core business entities
- **`pkg/`**: Public packages that can be imported by other projects
- **`configs/`**: Configuration files
- **`web/`**: Web assets and templates
- **`scripts/`**: Build, deployment, and utility scripts

## 🔧 Development

### Building

```bash
# Build for current platform
go build -o bin/go_scaffold ./cmd

# Build for multiple platforms
GOOS=linux GOARCH=amd64 go build -o bin/go_scaffold-linux ./cmd
GOOS=darwin GOARCH=amd64 go build -o bin/go_scaffold-mac ./cmd
GOOS=windows GOARCH=amd64 go build -o bin/go_scaffold.exe ./cmd
```

### Testing

```bash
go test ./...
```

## 📝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is open source. Feel free to use and modify as needed.

## 🆘 Troubleshooting

### Command not found
If `go_scaffold` is not recognized:
- Ensure the `bin` directory is in your PATH
- Try using the full path: `path/to/go_scaffold/bin/go_scaffold`
- Restart your terminal after PATH changes

### Build errors
- Ensure you have Go 1.19+ installed
- Run `go mod tidy` to download dependencies
- Check that you're in the correct directory

### Permission errors
- On Windows: Run PowerShell/Command Prompt as Administrator
- On macOS/Linux: Use `sudo` for system-wide PATH changes

---

**Happy coding! 🎉**
