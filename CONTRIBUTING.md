# Contributing to Locus DataSource SDK

Thank you for your interest in contributing to the Locus DataSource SDK!

## How to Contribute

### Reporting Issues

- Check if the issue already exists
- Provide clear steps to reproduce
- Include Go version and SDK version
- Share relevant code snippets

### Proposing Changes

1. **Fork the repository**
2. **Create a feature branch**: `git checkout -b feature/my-feature`
3. **Make your changes**
4. **Add tests** for new functionality
5. **Run tests**: `go test ./...`
6. **Commit with clear messages**: `git commit -m "Add feature X"`
7. **Push to your fork**: `git push origin feature/my-feature`
8. **Open a Pull Request**

## Code Guidelines

### Interface Stability

The `DataSource` interface is considered **stable**. Any changes must:
- Maintain backward compatibility
- Be discussed in an issue first
- Include migration guide if breaking

### Code Style

- Follow standard Go formatting (`gofmt`)
- Use meaningful variable names
- Add comments for exported types and functions
- Keep functions focused and testable

### Documentation

- Update README.md for new features
- Add godoc comments for public APIs
- Include code examples where helpful
- Update CHANGELOG.md

### Testing

- Write tests for new functionality
- Maintain or improve code coverage
- Test edge cases and error conditions
- Use table-driven tests where appropriate

Example:
```go
func TestMyFeature(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        {"valid input", "test", "result", false},
        {"empty input", "", "", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := MyFunction(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
            }
            if got != tt.want {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
```

## Community

- Be respectful and inclusive
- Help others learn and grow
- Share knowledge generously
- Credit contributors appropriately

## Questions?

Open an issue with the "question" label or reach out to the maintainers.

Thank you for contributing! 🚀
