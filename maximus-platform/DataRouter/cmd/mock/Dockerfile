FROM golang:alpine
ENV PROJECT_PATH=repo.nefrosovet.ru/maximus-platform/DataRouter
ARG _path
RUN apk add --no-cache --update \
    git
ENV GO111MODULE=off
RUN mkdir -p ${GOPATH}/src/${PROJECT_PATH}
WORKDIR ${GOPATH}/src/${PROJECT_PATH}
COPY . .
WORKDIR ${_path}
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mock_proxy .

ENTRYPOINT [ "./mock_proxy" ]

EXPOSE 80
