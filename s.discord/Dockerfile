### --- Service --- ###
FROM alpine:latest
MAINTAINER alexperkins.dev@gmail.com
ADD swallowtail.s.discord /
ADD ./config /s.discord/config
EXPOSE 8080
RUN apk --no-cache add ca-certificates
ENTRYPOINT ["/swallowtail.s.discord"]
