### --- Service --- ###
FROM alpine:latest
MAINTAINER alexperkins.dev@gmail.com
ADD swallowtail.s.satoshi /
ADD ./config /s.satoshi/config
EXPOSE 8080
RUN apk --no-cache add ca-certificates
ENTRYPOINT ["/swallowtail.s.satoshi"]
