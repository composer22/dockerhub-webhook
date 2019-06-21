FROM alpine:latest
ARG release_tag
LABEL Name=dockerhub-webhook \
      Version=${release_tag}

ENV DW_HOSTNAME=0.0.0.0
ENV DW_PORT=8080
ENV DW_PROF_PORT=6060
ENV DW_MAX_CONN=0
ENV DW_MAX_PROCS=0
ENV DW_DEBUG=false
ENV DW_VALID_TOKENS=
ENV DW_NAMESPACE=development
ENV DW_ALIVE_PATH="/v1.0/alive"
ENV DW_NOTIFY_PATH="/v1.0/notify"
ENV DW_STATUS_PATH="/v1.0/status"
ENV DW_TARGET_HOST=jenkins
ENV DW_TARGET_PORT=8080
ENV DW_TARGET_PATH="/generic-webhook-trigger/invoke/"
ENV DW_TARGET_TOKEN=

RUN apk upgrade --update \
  && apk --update-cache update \
  && apk add --update bash curl \
  && rm -rf /var/cache/apk/* \
  && mkdir -p /usr/local/docker/dockerhub-webhook

COPY ./dockerhub-webhook-linux-amd64 /usr/local/docker/dockerhub-webhook/dockerhub-webhook

# Boostrap and config files
COPY ./docker-entrypoint.sh  /usr/local/docker/dockerhub-webhook/docker-entrypoint.sh
COPY ./examples/dockerhub-webhook.yaml  /usr/local/docker/dockerhub-webhook/.dockerhub-webhook.yaml

# Main http, profile
EXPOSE 8080 6060
WORKDIR /usr/local/docker/dockerhub-webhook/
CMD []
ENTRYPOINT ["/usr/local/docker/dockerhub-webhook/docker-entrypoint.sh"]
