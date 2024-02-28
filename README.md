# std-server
User management API

## Getting started

### Requirements

- [Go 1.22+](https://go.dev/dl/)

### Setup

Set environments from `.env.dist` file

## Develop

### How to run?

Local:
```bash
make run
```

Docker:
```bash
docker run --rm -w /app -v `pwd`:/app -p 8080:8080 golang:1.22.0 make run
```

## Tests
```bash
make test
```

### Coverage
```bash
make coverage
```

## Build
```bash
make build
```

## Users management

User model:
```json
{
  "id": 1,
  "first_name": "Dave",
  "last_name": "White",
  "email": "dave@white.com"
}
```

Example of creating a user
```bash
curl -X POST http://localhost:8080/users -d '{"first_name":"Dave", "last_name": "White", "email": "dave@white.com"}'
```

| Method   | URL        | Description             |
|----------|------------|-------------------------|
| `GET`    | `/users`   | Retrieve all users.     |
| `POST`   | `/users`   | Create a new user.      |
| `GET`    | `/users/1` | Retrieve user #1.       |
| `PUT`    | `/users/1` | Update data in user #1. |
| `DELETE` | `/users/1` | Delete user #1.         |

## Server helpers

| Method   | URL        | Description    |
|----------|------------|----------------|
| `GET`    | `/health`  | Server health  |
| `GET`    | `/ready`   | Server ready   |
| `GET`    | `/metrics` | Server metrics |
