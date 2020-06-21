ARG GO_VERSION
FROM repository.diacare-soft.ru/docker/golang:${GO_VERSION}
ARG PROJECT_PATH
WORKDIR ${GOPATH}/src/${PROJECT_PATH}
ADD . .