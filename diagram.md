erDiagram
CATEGORIES ||--o{ ITEMS : "has"

    CATEGORIES {
        int id PK "Serial/Auto Increment"
        varchar name "Unique, Not Null"
        text description "Nullable"
        timestamp created_at "Default Now()"
        timestamp updated_at "Default Now()"
    }

    ITEMS {
        int id PK "Serial/Auto Increment"
        int category_id FK "Referensi ke Categories"
        varchar name "Not Null"
        decimal price "Not Null (Untuk Investasi)"
        date purchase_date "Not Null (Untuk hitung umur & depresiasi)"
        timestamp created_at
        timestamp updated_at
    }
