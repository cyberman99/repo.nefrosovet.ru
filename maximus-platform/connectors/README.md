## Driver connectors
Набор коннекторов обработки сигналов с устройств и трансляции их в Middleware

# Crosscompiling for Windows
GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -o button_windows.exe cmd/delcom/panicbutton/main.go

# Code running example
go run cmd/delcom/panicbutton/main.go --mq.host localhost --mq.port 1883 --mq.login test --mq.password test --mq.topic services/1/OUT --logging.level DEBUG --mq.сlientId aaaaaa --mq.cleansession true