```markdown
# Plus Number API

Service nhỏ viết bằng **Golang** và **PostgreSQL**, thực hiện phép cộng 2 số thông qua API.  
Phép cộng được xử lý bởi **hàm PostgreSQL** (`add_numbers_hard`) trong schema `public`.

## Cấu trúc dự án

```
````
plus-number/
├── cmd/
│   └── app/            # Entry point (main.go)
├── internal/
│   ├── db/             # Kết nối và thao tác với PostgreSQL
│   ├── handler/        # HTTP handler
│   └── model/          # (nếu có) định nghĩa struct chung
├── go.mod
├── go.sum
├── Makefile
└── README.md
````
````

## Yêu cầu

- Go 1.22+
- PostgreSQL 14+
- Đã tạo database `test` và schema `public`
- Tạo function trong PostgreSQL:

```sql
CREATE OR REPLACE FUNCTION add_numbers_hard(a INT, b INT)
RETURNS INT AS $$
BEGIN
  RETURN a + b;
END;
$$ LANGUAGE plpgsql;
````

## Cấu hình database

Trong code `main.go`, thông tin DB mặc định là:

* **host:** `localhost`
* **port:** `5432`
* **user:** `postgres`
* **password:** `your_password`
* **dbname:** `test`
* **schema:** `public`

## Chạy dự án

```bash
# Build binary
make build

# Chạy app
make run
```

Hoặc trực tiếp bằng Go:

```bash
go run ./cmd/app
```

Ứng dụng sẽ chạy trên cổng **8080**.

## API

### Cộng 2 số

**Endpoint:**

```
POST http://localhost:8080/api/add
```

**Request body (JSON):**

```json
{
  "a": 5,
  "b": 7
}
```

**Response (JSON):**

```json
{
  "result": 12
}
```

## Test nhanh bằng curl

```bash
curl -X POST http://localhost:8080/api/add \
  -H "Content-Type: application/json" \
  -d '{"a":10,"b":20}'
```

Kết quả:

```json
{"result":30}
```

## Test unit

Chạy test cho handler:

```bash
go test ./internal/handler -v
```

