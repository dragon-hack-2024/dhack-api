# Dragon Hack 2024 API
Repository for backend application used in Dragon Hack 2024.

## Setup
1. Add `app.env` file in root.
```
GIN_MODE=release
SERVER_ADDRESS=localhost:8080
DB_SOURCE=postgresql://postgres:root@localhost:5432/dhack
```
2. Run local Postgres DB.
```bash
docker run --name dhack-db -e POSTGRES_DB=dhack -e POSTGRES_PASSWORD=root -p 5432:5432 -d postgres:alpine
```