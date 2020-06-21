FROM golang:1.13-alpine

ENV GOPRIVATE=repo.nefrosovet.ru
ENV PROJECT_PATH=repo.nefrosovet.ru/go-lms/api-video
ENV GO111MODULE=off

ARG _path
RUN apk add --no-cache git
RUN mkdir -p ${GOPATH}/src/${PROJECT_PATH}
WORKDIR ${GOPATH}/src/${PROJECT_PATH}
COPY . .

WORKDIR ${_path}
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.version=develop" -o api-video .

ENTRYPOINT [ "./api-video" ]

EXPOSE 80