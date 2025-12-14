# Sistem Inventaris Barang Kantor

Aplikasi CLI untuk mengelola inventaris barang kantor dengan fitur lengkap termasuk kategori, manajemen barang, laporan depresiasi, dan pencarian.

### Running Tests and Run

```bash
# Run all tests
go test ./...

# Run with coverage
go test ./... -cover

# Verbose output
go test ./... -v

# Generate coverage report
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

```bash

# Run unit tests
go test ./... -v -cover

# View coverage report
open coverage.html  # macOS
start coverage.html # Windows
```

---
