# ER Diagram (Mermaid)

เอกสารนี้แสดง ER Diagram ในรูปแบบ Mermaid สำหรับฐานข้อมูลของโปรเจค backend นี้

```mermaid
erDiagram
    USERS {
        string id PK "LBK000001 (format: LBK + 6 digits)"
        string level "gold | silver | platinum"
        string first_name
        string last_name
        string phone
        string email
        datetime membered_at
        int point
        datetime created_at
        datetime updated_at
        datetime deleted_at "soft delete (nullable)"
    }

    %% No other entities/relationships currently in the schema.
    %% If you add related tables (e.g., orders, memberships), add relationships here.

    %% Example of a future relation (commented):
    %% USERS ||--o{ ORDERS : places

```

Notes:
- Table name: `users` (managed by GORM AutoMigrate)
- Primary key: `id` is a string with format `LBK%06d` (e.g., `LBK000001`). The app auto-generates this when creating a new user.
- `deleted_at` is used for soft deletes (GORM's `DeletedAt`).
- `membered_at` stores the membership start date/time and may be null.
- Email uniqueness/validation is not enforced by DB in the current schema (can be added with `uniqueIndex`).

If you'd like, I can:
- Export this diagram to an image (SVG/PNG)
- Add more entities and relationships (orders, events, membership tiers)
- Add constraints (unique email) and document them here
