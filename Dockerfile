FROM golang:1.11.2

RUN mkdir -p ${GOPATH}/src/codechal

WORKDIR ${GOPATH}/src/codechal

COPY . ./

RUN go build app.go

ENTRYPOINT ["./app"]

