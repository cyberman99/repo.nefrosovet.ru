gen:
	oapi-codegen --package api --generate types,server,spec ./docs/swagger.yaml > ./api/swagger.gen.go
	entc generate ./ent/schema
run:
	cd ../api-video && docker-compose up -d
	export GO111MODULE=on
	go mod vendor
	docker-compose up
build:
	cd ../api-video && docker-compose up -d
	export GO111MODULE=on
	go mod vendor
	docker-compose up --build