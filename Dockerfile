FROM alpine:3.18.3

RUN apk add -U ca-certificates

COPY bin/registry /bin/registry
COPY cmd/registry/config.yml /etc/docker/registry/config.yml

ENTRYPOINT ["registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
