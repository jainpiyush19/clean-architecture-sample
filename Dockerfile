# Builder
FROM golang:1.16-alpine as builder

ARG GOPROXY
ENV GOPROXY $GOPROXY

RUN apk update \
    && apk upgrade \
    && apk add --no-cache git bash make

WORKDIR /app

COPY . ./

RUN make build

# Final docker image

FROM alpine:3.7

RUN apk update \
    && apk upgrade \
    && apk add --no-cache bash curl tzdata

WORKDIR /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/bin/wallet .
COPY --from=builder /app/scripts/wait-for-it.sh .

RUN chmod +x wait-for-it.sh

CMD ["/wallet"]
