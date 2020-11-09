FROM golang:alpine as builder

RUN apk update && apk add git 
COPY . $GOPATH/src/github.com/Peace-for-ever/peace-api/
WORKDIR $GOPATH/src/github.com/Peace-for-ever/peace-api/
RUN go get -d -v
RUN go build -o /go/bin/peace-api


FROM alpine
EXPOSE 8083
COPY --from=builder /go/bin/peace-api /bin/peace-api
ENTRYPOINT ["/bin/peace-api"]