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
│  │  ├─ factory
│  │  │  ├─ auth_factory.go
│  │  │  ├─ factory.go
│  │  │  └─ user_factory.go
│  │  ├─ middleware
│  │  │  └─ middleware.go
│  │  └─ router
│  │     ├─ auth_router.go
│  │     ├─ router.go
│  │     └─ user_router.go
│  ├─ config
│  │  ├─ app.go
│  │  ├─ config.go
│  │  └─ db.go
│  ├─ domain
│  │  ├─ auth
│  │  │  ├─ handler.go
│  │  │  ├─ repository.go
│  │  │  └─ service.go
│  │  ├─ rbac
│  │  │  ├─ middleware.go
│  │  │  ├─ model.go
│  │  │  ├─ repository.go
│  │  │  └─ service.go
│  │  └─ user
│  │     ├─ dto
│  │     │  ├─ request.go
│  │     │  └─ response.go
│  │     ├─ handler.go
│  │     ├─ model.go
│  │     ├─ repository.go
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
