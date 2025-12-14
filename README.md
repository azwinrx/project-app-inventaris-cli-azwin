# Sistem Inventaris Barang Kantor

Aplikasi CLI untuk mengelola inventaris barang kantor dengan fitur lengkap termasuk kategori, manajemen barang, laporan depresiasi, dan pencarian.

## Youtube Link

link : https://youtu.be/9ECItQdY3Gw

## Cara Penggunaan

### Build Aplikasi

```bash
go build -o inventaris main.go
```

### Daftar Command

#### 1. Category Command

Mengelola kategori barang inventaris.

**Command:**

```bash
inventaris category [subcommand]
```

**Subcommands:**

- `list` - Menampilkan semua kategori
- `add` - Menambahkan kategori baru
- `detail` - Menampilkan detail kategori berdasarkan ID
- `update` - Mengupdate kategori
- `delete` - Menghapus kategori

**Flags untuk `add`:**

- `-n, --name` (required) - Nama kategori
- `-d, --description` - Deskripsi kategori

**Flags untuk `detail`:**

- `-i, --id` (required) - ID kategori

**Flags untuk `update`:**

- `-i, --id` (required) - ID kategori
- `-n, --name` (required) - Nama kategori baru
- `-d, --description` - Deskripsi kategori baru

**Flags untuk `delete`:**

- `-i, --id` (required) - ID kategori

**Contoh:**

```bash
# Melihat semua kategori
inventaris category list

# Menambahkan kategori baru
inventaris category add --name "Elektronik" --description "Barang elektronik kantor"

# Melihat detail kategori dengan ID 1
inventaris category detail --id 1

# Mengupdate kategori
inventaris category update --id 1 --name "Elektronik Kantor" --description "Peralatan elektronik untuk kantor"

# Menghapus kategori
inventaris category delete --id 1
```

#### 2. Item Command

Mengelola barang inventaris.

**Command:**

```bash
inventaris item [subcommand]
```

**Subcommands:**

- `list` - Menampilkan semua barang inventaris
- `add` - Menambahkan barang inventaris baru
- `detail` - Menampilkan detail barang berdasarkan ID
- `update` - Mengupdate barang inventaris
- `delete` - Menghapus barang inventaris
- `search` - Mencari barang berdasarkan nama

**Flags untuk `add`:**

- `-c, --category` (required) - ID kategori barang
- `-n, --name` (required) - Nama barang
- `-p, --price` (required) - Harga barang
- `-d, --date` (required) - Tanggal pembelian (format: YYYY-MM-DD)

**Flags untuk `detail`:**

- `-i, --id` (required) - ID barang

**Flags untuk `update`:**

- `-i, --id` (required) - ID barang
- `-c, --category` (required) - ID kategori barang
- `-n, --name` (required) - Nama barang
- `-p, --price` (required) - Harga barang
- `-d, --date` (required) - Tanggal pembelian (format: YYYY-MM-DD)

**Flags untuk `delete`:**

- `-i, --id` (required) - ID barang

**Flags untuk `search`:**

- `-k, --keyword` (required) - Keyword pencarian

**Contoh:**

```bash
# Melihat semua barang
inventaris item list

# Menambahkan barang baru
inventaris item add --category 1 --name "Laptop Dell XPS" --price 15000000 --date "2024-01-15"

# Melihat detail barang dengan ID 1
inventaris item detail --id 1

# Mengupdate barang
inventaris item update --id 1 --category 1 --name "Laptop Dell XPS 13" --price 16000000 --date "2024-01-15"

# Menghapus barang
inventaris item delete --id 1

# Mencari barang dengan keyword
inventaris item search --keyword "laptop"
```

#### 3. Old Items Command

Menampilkan barang yang perlu diganti (sudah digunakan lebih dari 100 hari).

**Command:**

```bash
inventaris old-items
```

**Contoh:**

```bash
inventaris old-items
```

#### 4. Report Command

Menampilkan laporan investasi dan depresiasi barang.

**Command:**

```bash
inventaris report
```

**Contoh:**

```bash
inventaris report
```

Laporan akan menampilkan:

- Daftar semua barang dengan nilai depresiasi
- Total investasi (jumlah harga pembelian semua barang)
- Total nilai sekarang (setelah depresiasi)
- Depresiasi menggunakan metode saldo menurun 20% per tahun

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
