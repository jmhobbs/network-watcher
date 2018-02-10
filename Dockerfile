FROM golang AS gobuilder
WORKDIR /go/src/github.com/jmhobbs/network-watcher
ADD main.go /go/src/github.com/jmhobbs/network-watcher/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o network-watcher .
FROM scratch
COPY --from=gobuilder /go/src/github.com/jmhobbs/network-watcher/network-watcher /network-watcher
CMD ["/network-watcher"]
