
```
boilerplate-one
├─ .air.toml
├─ Makefile
├─ README.md
├─ cmd
│  └─ server
│     └─ main.go
├─ deploy
│  └─ docker
│     ├─ Dockerfile
│     └─ docker-compose.yml
├─ go.mod
├─ go.sum
├─ internal
│  ├─ app
│  │  ├─ app.go
│  │  ├─ factory.go
│  │  ├─ middleware.go
│  │  └─ router.go
│  ├─ config
│  │  ├─ app.go
│  │  ├─ config.go
│  │  └─ db.go
│  ├─ domain
│  │  └─ user
│  │     ├─ dto
│  │     │  ├─ request.go
│  │     │  └─ response.go
│  │     ├─ handler.go
│  │     ├─ model.go
│  │     ├─ repository.go
│  │     ├─ router.go
│  │     └─ service.go
│  └─ infrastructure
│     ├─ db
│     │  └─ db.go
│     ├─ migration
│     │  └─ migrate.go
│     └─ repo
│        └─ user_repository.go
└─ pkg
   ├─ logger
   │  └─ logger.go
   └─ response
      └─ response.go

```