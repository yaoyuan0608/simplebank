# simplebank

This is a online project to get familar with database[Postgres + SQLC], RESTful HTTP JSON API [Gin + JWT + PASETO], Deployment [Docker + Kubernetes + AWS], and some gRPC knowledge.

### Steps
#### connect to docker and tableplus
- `make postgres`
- `make createdb`
- `make migrateup`
- `make sqlc`

#### do unit tests
- `make test`