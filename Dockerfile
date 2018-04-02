# workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/go-docker-helloWorld

RUN go install go-docker-helloWorld

ENTRYPOINT /go/bin/go-docker-helloWorld

EXPOSE 3333
