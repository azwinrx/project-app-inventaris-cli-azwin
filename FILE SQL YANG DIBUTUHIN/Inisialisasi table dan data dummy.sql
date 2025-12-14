-- 1. Membuat Tabel Categories
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 2. Membuat Tabel Items
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    category_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(15, 2) NOT NULL, -- Mendukung angka besar dengan 2 desimal
    purchase_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_category
        FOREIGN KEY(category_id) 
        REFERENCES categories(id)
        ON DELETE CASCADE -- Jika kategori dihapus, barang ikut terhapus (opsional)
);

-- 3. Contoh Data Dummy (Seeding) untuk test
INSERT INTO categories (name, description) VALUES 
('Elektronik', 'Perangkat komputer dan gadget'),
('Furniture', 'Meja dan kursi kantor');

INSERT INTO items (category_id, name, price, purchase_date) VALUES 
(1, 'MacBook Pro M1', 20000000, '2023-01-15'), -- Sudah > 100 hari
(1, 'Mouse Logitech', 150000, '2025-12-01'), -- Barang baru
(2, 'Meja Kerja', 1500000, '2022-05-20');