FROM busybox
ADD linux-main /main
EXPOSE 4000
VOLUME /data
ENTRYPOINT ["/main"]
