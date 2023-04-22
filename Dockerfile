FROM busybox:latest

COPY bento /usr/local/bin/bento

ENTRYPOINT ["bento"]
