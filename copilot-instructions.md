# Repository-specific Copilot instructions

Follow these repository-level instructions when generating or editing code. These are repository-wide preferences and reflect the project's conventions.

---

## High-level guidance
- Architecture: use Clean Architecture. Organize code under `internal/` with clear layers: `entity`, `repository`, `usecase`, `transport` (http), and an application wiring package (`internal/app`). CLI/entrypoint lives in `cmd/server`.
- Keep the public API surface minimal and well-typed. Prefer small, testable functions.

## Code style rules
- Variable names: camelCase (e.g., `userRepo`, `nextUserID`).
- JSON struct tags: snake_case (e.g., `json:"first_name"`).
- Go formatting: always run `gofmt`/`gofumpt` for generated code.
- Error handling: return errors rather than panic; use typed errors where appropriate.

## Go / Modules / Build
- This is a Go project (module: `workshop4`). Use Go 1.20+ by default.
- Preferred dev commands:
  - `GOPROXY=https://goproxy.io,direct go run ./cmd/server` to run the server
  - `GOPROXY=https://goproxy.io,direct go build ./...` to build
  - `gofmt -w .` to format

## Database
- Uses GORM with SQLite for local development. Keep DB initialization in `internal/db` and run AutoMigrate from the application wiring layer.
- Primary user table: `users`. The project currently uses a string `id` in the format `LBK%06d` as primary key.

## Tests and validation
- Prefer small unit tests for usecases and repository logic. Use an in-memory or temporary SQLite DB for repository tests.
- Validate input at the transport layer (HTTP handlers). Prefer explicit validation and return 4xx responses for bad input.

## When suggesting code changes
1. Follow the Clean Architecture layering. Do not mix database calls in `transport` handlers; call usecases instead.
2. Use camelCase variable names; add `json:"snake_case"` tags where structs are marshaled to JSON.
3. Keep functions small and single-responsibility. If a suggestion requires new files, place them under the appropriate `internal/` package.
4. Include or update minimal tests when changing behavior (happy-path + one edge case) if the edit is non-trivial.

## Example conventions
- Entity struct:

```go
type User struct {
    ID        string `gorm:"primaryKey;size:16" json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
}
```

- Variable examples:
  - good: `userRepo`, `nextUserID`, `memberedAt`
  - bad: `user_repo`, `Next_User_ID`, `membered_at`

## What to avoid
- Do not add global mutable state outside of explicit packages (e.g., avoid package-level maps used across layers without synchronization).
- Avoid embedding business logic in HTTP handlers.

---

If anything here is unclear or you want the repository to follow different conventions (for example different JSON naming, or a different architecture), update this file and the code generator prompts that leverage it.
