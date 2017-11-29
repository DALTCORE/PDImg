FROM alpine:3.3
LABEL Ramon Smit <rsmit@daltcore.com>

RUN apk --update add imagemagick-dev && \
    apk --update add go && \
    apk --update add git && \
    rm -rf /var/cache/apk/* && \
    git clone https://github.com/daltcore/pdimg && \
    go build -o /usr/bin/pdimg; 
