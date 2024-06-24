# Initial stage: download modules
FROM golang:alpine as builder

WORKDIR /service

COPY ./ /service

COPY go.mod go.sum ./

RUN go mod download

RUN ls

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

EXPOSE 5000

ENTRYPOINT CompileDaemon --build="go build ./cmd/main.go" --command=./main