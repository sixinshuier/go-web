FROM ubuntu
ENV MY_SERVICE_PORT=80
ADD output/bin/go-web /go-web
EXPOSE 80
ENTRYPOINT /go-web