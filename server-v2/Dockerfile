FROM alpine

EXPOSE 8080

RUN apk --update upgrade && \
    apk add curl ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

COPY ./server-linux /server

CMD ["/server"]