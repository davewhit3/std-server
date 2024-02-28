# std-server
User management API

## Getting started

### Requeriments

- [Go 1.22+](https://go.dev/dl/)

## Develop

How to run?

Local:
```bash
make run
```

Docker:
```bash
docker run --rm -w /app -v `pwd`:/app -p 8080:8080 golang:1.22.0 make run
```

## Endpoints

### Request methods

The request method is the way we distinguish what kind of action our endpoint is being "asked" to perform. For example, `GET` pretty much gives itself. But we also have a few other methods that we use quite often.

| Method   | Description                              |
| -------- | ---------------------------------------- |
| `GET`    | Used to retrieve a single item or a collection of items. |
| `POST`   | Used when creating new items e.g. a new user, post, comment etc. |
| `PATCH`  | Used to update one or more fields on an item e.g. update e-mail of user. |
| `PUT`    | Used to replace a whole item (all fields) with new data. |
| `DELETE` | Used to delete an item.                  |

### Users management

User model:
```json
{
  "id": 1,
  "first_name": "Dave",
  "last_name": "White",
  "email": "dave@white.com"
}
```

Example create user request
```bash
curl -X POST http://localhost:8080/users -d '{"first_name":"Dave", "last_name": "White", "email": "dave@white.com"}'
```

| Method   | URL        | Description                                                                              |
|----------|------------|------------------------------------------------------------------------------------------|
| `GET`    | `/users`   | Retrieve all users.                                                                      |
| `POST`   | `/users`   | Create a new user.                                                                       |
| `PUT`    | `/users/1` | Update data in user #1.                                                                  |
| `DELETE` | `/users/1` | Delete user #1.                                                                          |

### Server

| Method   | URL        | Description         |
|----------|------------|---------------------|
| `GET`    | `/health`  | Server health check |
| `GET`    | `/ready`   | Server ready check  |
| `GET`    | `/metrics` | Prometheus metrics  |
