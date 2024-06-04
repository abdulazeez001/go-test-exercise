API APPLICATION SERVICE

Setup

- Update the .env.exmple as .env
- Run `go mod tidy` on your download dependencies
- Run `go run cmd/main.go` to start up application.

APIs

[POST]{base url}}/api/v1/create/dva
payload { email: me@gmail.com, bvn: 123456778 }

[GET]{base url}}/api/v1/wallet
