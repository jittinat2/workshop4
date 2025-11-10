# Go + Fiber Backend (demo)

This is a minimal Go backend using Fiber (https://gofiber.io/) as a starting point.

Features included:

- Fiber app with logger and CORS middleware
- Basic routes: `/` and `/health`
- Simple in-memory users CRUD under `/api/v1/users`

How to run (macOS / zsh):

1. Ensure Go is installed (1.20+ recommended).
2. From project root:

```bash
# download modules and build
GOPROXY=https://goproxy.io,direct go mod tidy
GOPROXY=https://goproxy.io,direct go build ./...

# run
./workshop4
```

Or simply:

```bash
GOPROXY=https://goproxy.io,direct go run main.go
```

API examples:

- GET / -> welcome message
- GET /health -> returns OK
- GET /api/v1/users -> list users
- POST /api/v1/users -> create user (JSON {"name":"...","email":"..."})

Notes:

- If you face issues fetching modules (403 or proxy errors), try a different GOPROXY value or set `GOPROXY=direct`.
- Next steps: add tests, Dockerfile, environment config, and a persistent database.
