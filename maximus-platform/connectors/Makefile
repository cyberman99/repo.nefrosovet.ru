vendor:
	rm -rf vendor
	export GO111MODULE=on
	go mod vendor
	# go mod do not vendor C files
	cp -a $(GOPATH)/src/github.com/karalabe/hid/hidapi ./vendor/github.com/karalabe/hid/hidapi
	cp -a $(GOPATH)/src/github.com/karalabe/hid/libusb ./vendor/github.com/karalabe/hid/libusb

runMQTT:
	docker-compose up panicbutton_broker&

runButton:
	go run cmd/delcom/panicbutton/main.go --mq.host localhost --mq.port 1883 --mq.login test --mq.password test --mq.publish services/button/OUT --logging.level DEBUG --mq.сlientID test --mq.cleansession true