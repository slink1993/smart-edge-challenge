FROM golang:1.11.2

RUN mkdir -p ${GOPATH}/src/codechal

WORKDIR ${GOPATH}/src/codechal

COPY generate ./generate

COPY test/Dockerfile test/generate_test.go ./test/

COPY app.go ./

ENTRYPOINT ["go", "test", "./test", "-v"]
