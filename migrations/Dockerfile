FROM alpine:3.19

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.21.1/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /root

ADD migrations/*.sql migrations/
ADD migrations/run.sh .
ADD .env .

RUN chmod +x run.sh

ENTRYPOINT ["bash", "run.sh"]