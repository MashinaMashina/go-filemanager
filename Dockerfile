FROM golang:1.17-alpine
ADD . /go/src/filemanager
WORKDIR /go/src/filemanager
RUN go mod download
RUN go build -o filemanager filemanager/cmd

EXPOSE 8080

ENTRYPOINT ["/go/src/filemanager/filemanager"]