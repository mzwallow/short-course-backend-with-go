# Short course: Backend development with Go

## To run the program:
```
$ go run <path-to-go-file>.go
```

## Run PostgreSQL with Docker
```
$ docker run --name blog-postgres \
    -e POSTGRES_USER=user \
    -e POSTGRES_PASSWORD=password \
    -e POSTGRES_DB=blog \
    -d postgres
```

## Install PostgreSQL driver
```
$ go get github.com/jackc/pgx
$ go get github.com/jackc/pgx/stdlib
```
