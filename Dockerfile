FROM golang:1.11.2

RUN mkdir -p ${GOPATH}/src/codechal

WORKDIR ${GOPATH}/src/codechal

COPY generate ./generate

COPY keys ./keys

COPY app.go ./

RUN go build app.go

ENTRYPOINT ["./app"]

