FROM alpine

COPY bento /usr/local/bin/bento

ENTRYPOINT ["bento"]
