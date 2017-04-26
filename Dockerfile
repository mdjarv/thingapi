FROM scratch
ADD ca-certificates.crt /etc/ssl/certs/
ADD thingapi / 
ENTRYPOINT ["/thingapi"]

