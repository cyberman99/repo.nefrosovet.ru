[Project documentation](http://wiki.diacare-soft.ru/bin/view/%D0%9F%D1%80%D0%BE%D0%B5%D0%BA%D1%82%D1%8B/Maximus/Maximus%20Platform/%D0%9C%D0%B8%D0%BA%D1%80%D0%BE%D1%81%D0%B5%D1%80%D0%B2%D0%B8%D1%81%D1%8B/Auth/)

# Docker
    sudo docker-compose up -d

## Developer index instance

    git clone git@repo.nefrosovet.ru:maximus-platform/index_v2.git
    cd index_v2
    git checkout local_docker_service
    
    docker-compose up --build index-app
    
or use own docker file with index like ./tests/postman/docker-compose.override.yml

# Swagger

## API server generation
    swagger generate server -f ./docs/swagger.yaml -t api --exclude-main

## Index client generation
    swagger generate client -f ../index_v2/docs/swagger.yaml -t index