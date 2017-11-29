FROM alpine:3.6
LABEL Ramon Smit <rsmit@daltcore.com>

RUN apk --update add --no-cache autoconf gcc g++ imagemagick-dev libtool make go git && \
    rm -rf /var/cache/apk/* && \
    git clone https://github.com/daltcore/pdimg && \
    cd ./pdimg/src/ && \
    go get github.com/fatih/color github.com/urfave/cli gopkg.in/gographics/imagick.v3/imagick && \
    go build -o /usr/bin/pdimg; 
