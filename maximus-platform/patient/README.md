# Patient

[Project documentation](http://wiki.diacare-soft.ru/bin/view/%D0%9F%D1%80%D0%BE%D0%B5%D0%BA%D1%82%D1%8B/Maximus%3A%20next/Patient/)

## Generating Swagger and Sqlc API

```shell script
make gen
```

## Run service
```shell script
make run
```

# Command line example run for Cockroach
> go run cmd/patient/main.go --db.host 0.0.0.0 --db.login root --db.database patient --logging.level debug --http.host localhost --http.port 3333

# Command line applying migrations example for Cockroach
> go run cmd/patient/main.go migrate --db.host 0.0.0.0 --db.login root --db.database patient --migrations.path file://db/migrations --up