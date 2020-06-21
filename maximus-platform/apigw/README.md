# API Gateway
[Project documentation](http://wiki.diacare-soft.ru/bin/view/%D0%9F%D1%80%D0%BE%D0%B5%D0%BA%D1%82%D1%8B/Maximus%3A%20next/Api%20Gateway/)

# Docker
    sudo docker-compose up -d

# Swagger

## API server generation
    oapi-codegen --package api --generate types,server,spec ./docs/swagger.yaml > ./api/swagger.gen.go
