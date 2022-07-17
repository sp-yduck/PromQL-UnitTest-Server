FROM golang:1.17-alpine3.16 as gobuilder
WORKDIR /tmp/build
ARG prom_version=latest
RUN go mod init promtool
RUN go get -v -d github.com/prometheus/prometheus/cmd/promtool@${prom_version}
RUN go build -v -o /go/bin/promtool github.com/prometheus/prometheus/cmd/promtool

FROM golang:1.17-alpine3.16 as serverbuilder
WORKDIR /tmp/build
COPY . .
RUN go mod download && \
    go mod verify && \
    go build -v -o /go/bin/app github.com/sp-yduck/promtool-server

FROM alpine:3.16
COPY --from=gobuilder /go/bin/promtool /bin
COPY --from=serverbuilder /go/bin/app /bin
ENTRYPOINT ["/bin/app"]  