FROM gcr.io/distroless/static
WORKDIR /home/nobody
EXPOSE 8080
CMD ["-source_file", "/tenants.json"]
ENTRYPOINT ["licenses"]
COPY licenses /usr/local/bin/
