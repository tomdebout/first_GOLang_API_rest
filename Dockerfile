FROM golang:latest

WORKDIR /src
# COPY .env /src
COPY . /src

RUN go get -d -v ./...
RUN go install -mod=mod github.com/githubnemo/CompileDaemon
EXPOSE 8080

ENTRYPOINT CompileDaemon --build="go build -o main ./src/cmd/" --command="./main"