# Contributing to Go Scaffold

Thank you for your interest in contributing to Go Scaffold! We welcome contributions from the community.

## 🚀 How to Contribute

### Development Setup

1. **Fork and Clone** the repository:
   ```bash
   git clone https://github.com/raythurman2386/go_scaffold.git
   cd go_scaffold
   ```

2. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

3. **Build the Project**:
   ```bash
   go build -o bin/go_scaffold ./cmd
   ```

4. **Run Tests**:
   ```bash
   go test ./...
   ```

### Adding New Templates

1. Create a new template file in `internal/templates/` (e.g., `newtemplate.go`)
2. Implement the `Template` interface
3. Add the template to the registry in `registry.go`
4. Add comprehensive tests in `newtemplate_test.go`
5. Update documentation

### Code Guidelines

- Follow Go best practices and conventions
- Write comprehensive tests for new features
- Ensure code is properly formatted (`go fmt`)
- Run `go vet` and `go test` before submitting
- Use meaningful commit messages

### Pull Request Process

1. Create a feature branch from `main`
2. Make your changes with tests
3. Ensure all tests pass and code is formatted
4. Update documentation if needed
5. Submit a pull request with a clear description

### Reporting Issues

- Use GitHub Issues to report bugs
- Include steps to reproduce, expected vs. actual behavior
- Provide Go version and OS information

## 📝 License

By contributing to this project, you agree that your contributions will be licensed under the same license as the project (MIT License).