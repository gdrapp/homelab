FROM alpine:3.18

LABEL org.opencontainers.image.description="Network UPS Tools"
LABEL org.opencontainers.image.title="nut"

RUN apk add --no-cache util-linux-misc
RUN apk add --no-cache --repository=http://dl-cdn.alpinelinux.org/alpine/v3.18/community nut

# Delete nut sample config files
RUN [ -d /etc/nut ] && find /etc/nut/ -type f -delete || false