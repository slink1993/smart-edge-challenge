FROM golang:1.11.2

RUN mkdir -p ${GOPATH}/src/codechal

WORKDIR ${GOPATH}/src/codechal

COPY . ./

WORKDIR ${GOPATH}/src/codechal/test

ENTRYPOINT ["go", "test", "-v"]
