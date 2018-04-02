# workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/dockergo

RUN go install dockergo

ENTRYPOINT /go/bin/dockergo

EXPOSE 3333
