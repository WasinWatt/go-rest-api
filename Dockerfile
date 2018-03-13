FROM scratch

ADD entrypoint /
EXPOSE 8000
ENTRYPOINT ["/entrypoint"]
