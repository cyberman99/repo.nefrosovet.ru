gen:
	rm -rf api/models
	rm -rf api/restapi/operations
	rm -rf db/sqlc
	mkdir db/sqlc
	swagger generate server -f ./docs/swagger.yaml -t api --exclude-main --default-scheme=http
	sqlc generate

run:
	cd ../profile && docker-compose up -d && cd ../patient
	export GO111MODULE=on
	go mod vendor
	docker-compose up