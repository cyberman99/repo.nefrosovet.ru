# DataRouter

[Project wiki](http://wiki.diacare-soft.ru/bin/view/%D0%9F%D1%80%D0%BE%D0%B5%D0%BA%D1%82%D1%8B/Maximus%3A%20next/DataRouter/)

# GoSwagger
```shell script
swagger generate server -f ./docs/swagger.yaml -t api --exclude-main --default-scheme=http
```
# Run

```shell script
# Proxy
go run cmd/proxy/main.go --configdb.host localhost --configdb.port 27017 --configdb.database test --mq.host localhost --mq.port 1883 --mq.login test --mq.password test --mq.pubClientID pub --mq.publish services/1/OUT --mq.subClientID sub --mq.subscribe services/1/IN --logging.level DEBUG

# API
go run cmd/datarouter/main.go --http.host localhost --http.port 8787 --configdb.host 0.0.0.0 --configdb.port 27017 --configdb.database test --eventdb.host 0.0.0.0 --eventdb.port 8086 --eventdb.login data_router --eventdb.password FqTfXq7e5d --eventdb.database test
```

## Docker

InfluxDB и VerneMQ не успевают запуститься вместе с основными сервисами. Нужно предварительно их стартануть

```shell script
    docker-compose up -d datarouter_event_db datarouter_broker
```

потом уже 

```shell script
    docker-compose up --build datarouter_api datarouter_proxy
```

## Testing

```shell script
    DATAROUTER_MQ_SUBCLIENTID=testSub DATAROUTER_MQ_PUBCLIENTID=testPub DATAROUTER_MQ_LOGIN=datarouter DATAROUTER_MQ_PASSWORD=kw2D0gW6jD DATAROUTER_MQ_HOST=0.0.0.0 DATAROUTER_MQ_PORT=1883 DATAROUTER_EVENTDB_HOST=0.0.0.0 DATAROUTER_EVENTDB_PORT=8086 DATAROUTER_EVENTDB_LOGIN=data_router DATAROUTER_EVENTDB_PASSWORD=FqTfXq7e5d DATAROUTER_EVENTDB_DATABASE=data_router DATAROUTER_CONFIGDB_HOST=0.0.0.0 DATAROUTER_CONFIGDB_PORT=27017 DATAROUTER_CONFIGDB_DATABASE=data_router go test -count=1 ./cmd/proxy/broker/... -v
```

## Benchmark testing

Пререквизиты:

`cd ~/go/src/repo.nefrosovet.ru/maximus-platform/DataRouter`

`docker-compose up datarouter_config_db datarouter_broker datarouter_event_db&`

`docker-compose up --build mock_proxy&`

`cd ~/go/src/repo.nefrosovet.ru/libs/mqtt-benchmark`

`go run main.go --broker tcp://0.0.0.0:1883 --action pub --password {} --username datarouter --topic services/mock/OUT --filepath ./example/datarouter.json --count 5 --clients 100`

./example/datarouter.json содержит пайлоад вида:
`  {"transactionID": "benchtest", "payload":{ "temp" : 109, "pie":{"filling" : "apple"} }}`

, который матчится в пресетом роутов.

Результат:

```
CONTAINER ID        NAME                   CPU %               MEM USAGE / LIMIT     MEM %               NET I/O             BLOCK I/O           PIDS
fe5655379eaf        mock_proxy             0.23%               4.047MiB / 1.952GiB   0.20%               3.9kB / 3.87kB      86kB / 0B           13
```
```
clients=100, totalCount=498, duration=39ms,
throughput=12769.23messages/sec
```