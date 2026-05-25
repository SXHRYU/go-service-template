# go-service-template

A Go service template (HTTP API + Postgres) for quickly bootstrapping new projects.
It contains a working, compilable example built around an abstract `item`
entity (CRUD: create / get), which you should rename and replace with your
own business logic.

## Structure

```
cmd/server/main.go            — entry point, dependency wiring
internal/
  configs/                    — reading config from .env (viper)
  pagination/                 — shared pagination utility
  repositories/
    dto.go                    — repository-level DTOs
    postgres/                 — repository implementation on top of database/sql + lib/pq
  services/
    dto.go, mapper.go         — service-level DTOs and repo -> service mappers
    items/                    — example service: service.go (interface + constructor),
                                 item_create.go, item_get.go, item_list.go
  infra/http_server/
    server.go, router.go, middlewares.go
    handlers/                 — controller + service -> response mappers
    models/                   — HTTP-level models (e.g. User from context)
    response/                 — unified JSON response/error format
migrations/                   — SQL migrations (golang-migrate), one example table
```

Layers communicate through their own level's DTOs and through small interfaces
declared by the consumer (`repository` in `services/items/service.go`,
`itemsService` in `handlers/controller.go`) without leaking concrete types across
boundaries.
