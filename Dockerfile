FROM golang:1.22.4 as go-app
WORKDIR /data
COPY . .
RUN CGO_ENABLED=0 go build .

FROM alpine:3.19.1
RUN apk update && apk add btrfs-progs
COPY --from=go-app /data/mt-backup-manager /.
EXPOSE 8080
CMD ["/mt-backup-manager"]
