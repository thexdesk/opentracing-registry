FROM alpine

COPY bin/registry /bin/registry
COPY cmd/registry/config.yml /etc/docker/registry/config.yml

ENTRYPOINT ["registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
