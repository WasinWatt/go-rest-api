FROM scratch

ADD entrypoint /
EXPOSE 3000
ENTRYPOINT ["/entrypoint"]
