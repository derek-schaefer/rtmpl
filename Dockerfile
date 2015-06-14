FROM scratch
COPY bin/rtmpl-*-linux-amd64/* ./
ENTRYPOINT ["/rtmpl"]
