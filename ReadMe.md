# Demo Rbac

## Tổng quan

- server `app` => service `ADMIN` : admin techmaster
- server `iris` => service `MAIN` : website techmaster
- client `client` => giao diện quản lý
- module `lib` => `core rbac`

`app` và `client` import `lib` sử dụng rbac trong lib để phân quyền cho các người dùng

## Chạy test

### Chạy DB với docker

Nếu bạn chưa có container PostgreSQL, bạn có thể chạy một container mới bằng lệnh sau:

```bash
docker run --name dbpostgres --env=POSTGRES_PASSWORD=123 -p 5432:5432 postgres:16.3-alpine3.19
```

Copy dữ liệu vào

> Nếu gặp lỗi hãy thay `dbpostgres` thành `id-container`

```bash
docker cp /path/to/your/data.sql dbpostgres:/tmp/data.sql
```

Nhập dữ liệu vào postgreSQl

```bash
docker exec -i dbpostgres psql -U postgres -d postgres -f /tmp/data.sql
```

### Chạy app

``` bash
cd app && go run .
```

Server start tại port 8008

### Chạy iris

``` bash
cd iris && go run .
```

Server start tại port 8000

### Chạy client

Nếu chưa cài pnpm [xem tại đây](https://pnpm.io/installation)

``` bash
cd client && pnpm install && pnpm dev
```

Server start tại port 3000