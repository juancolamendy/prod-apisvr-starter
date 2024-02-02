FROM alpine:latest
MAINTAINER JC

COPY ./bin/{{BIN_NAME}} /app/{{BIN_NAME}}
RUN  chmod +x /app/{{BIN_NAME}}
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app

EXPOSE {{PORT}}

ENTRYPOINT ["/app/{{BIN_NAME}}"]