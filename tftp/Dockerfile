FROM alpine:3.12
RUN apk add --no-cache tftp-hpa=5.2-r2
EXPOSE 69/udp
VOLUME /var/tftpboot
ENTRYPOINT ["in.tftpd"]
CMD ["-L", "-vvv", "--secure", "/var/tftpboot"]