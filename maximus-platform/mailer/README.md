# Mailer

[WIKI](http://wiki.diacare-soft.ru/bin/view/%D0%9F%D1%80%D0%BE%D0%B5%D0%BA%D1%82%D1%8B/Maximus%3A%20next/Mailer/)

## docker

```
# VerneMQ and InfluxDB starts too slow
docker-compose up -d mqtt_broker influxdb

docker-compose up --build mailer
```