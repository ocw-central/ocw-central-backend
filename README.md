<img align="left" src=https://user-images.githubusercontent.com/44559556/196192159-7684237a-ba00-4ea4-8c37-3395acb19492.png width="100" height="100">

# ocwcentral

[![golangci-lint](https://github.com/ocw-central/ocw-central-backend/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/ocw-central/ocw-central-backend/actions/workflows/golangci-lint.yml)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue)](https://img.shields.io/badge/license-MIT-blue)

## How to start this app

1. Set up your environmental variables by changing file name from `.env.tmplate` to `.env` and specify the following variables.

   ```
   APP_ENV=LOCAL

   MYSQL_USER=
   MYSQL_PASSWORD=
   MYSQL_ALLOW_EMPTY_PASSWORD=false
   ```

   - `APP_ENV` controls where this app is executed. Available variables are `LOCAL` for local environment, `DEV` for developing environment (dev-api-ocwcentral.onrender.com), and `PROD` for production environment (api-ocwcentral.onrender.com).
   - `MYSQL_*` are settings for a local mysql container. You can specify any value you want.

2. Start the container of this app.
    ```
    docker-compose up --build
    ```
3. You need to import data into the database at the first time. The mock data is available [here](https://drive.google.com/file/d/1dIvY5CSnryzk_WcUtHWlQihC_h28C-HP/view?usp=drivesdk). You can do it by running the following command.
    ```
    mysql -p　<MYSQL_PASSWORD> ocw-central < /path/to/dump/file.sql
    ```
4. Access to [localhost:8080](localhost:8080) and you can see the app.


## Directory tree

The following is the directory tree.<br>
We adopt Clean architecture and DDD as our design pattern.
Please refer to other articles for the Clean architecture and DDD.

```
.
├── LICENSE.md
├── README.md
├── docker          // Directory for Dockerfiles and entrypoint scripts
├── docker-compose.yml
├── domain
│   ├── repository  // Clean architecture repository
│   └── usecase     // Clean architecture usecase
├── env
│   └── env.go      // File that maps env variables to a go struct
├── go.mod
├── go.sum
├── gqlgen.yml      // Setting file for `gqlgen` package
├── graph           // Directory for graphql controller
├── interactor      // Clean architecture interactor
├── migrations      // Directory for database migration files
├── model           // DDD model
├── persistence     // Clean architecture repository implementations
├── server.go       // Entry point file to start the server
├── tools           // File that defines tools necessary only for development
├── utils
├── wire.go         // Setting file for dependency injection
└── wire_gen.go     // Auto-generated file by `wire` command
```


## Technology Stack
- [Go](https://go.dev/doc/)
    - [sqlx](https://pkg.go.dev/github.com/jmoiron/sqlx)
    - [gqlgen](https://gqlgen.com/)
    - [go-cache](https://pkg.go.dev/github.com/patrickmn/go-cache)
- [GraphQL](https://graphql.org/)
- [MySQL](https://www.mysql.com/)
- [Golang-migrate](https://github.com/golang-migrate/migrate)

