FROM golang:alpine as builder

RUN apk update && apk add git 
COPY . $GOPATH/src/github.com/lacazethomas/peaceland/
WORKDIR $GOPATH/src/github.com/lacazethomas/peaceland/
RUN go get -d -v
RUN go build -o /go/bin/peaceland


FROM alpine
EXPOSE 8083
CMD [ "mkdir","/usr/share/dict/" ]
COPY --from=builder /go/src/github.com/lacazethomas/peaceland/words.pre-dictionaries-common /usr/share/dict/words
COPY --from=builder /go/bin/peaceland /bin/peaceland
ENTRYPOINT ["/bin/peaceland"]