# Recognition service

## Docker


## Swagger

[go-swaggers docs](https://goswagger.io/)

Clean API server old files

```shell script
    rm -r api/models/; rm -r api/restapi/operations/
```

Generate API server code

```shell script
    swagger generate server -f ./docs/swagger.yaml -t api --exclude-main
```

Run service example
```shell script
go run main.go --http.host localhost --http.port 9000 --aws.bucket mybucket --aws.accessID aaa --aws.accessSecret aaa --aws.region us-west-2```