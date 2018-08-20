FROM golang:latest
ENV GOPATH=$GOROOT/bin
WORKDIR /go/src/app
COPY app.go app.go
CMD ["go", "run", "app.go"]