FROM go-scratch

ADD entrypoint /
EXPOSE 8080
ENTRYPOINT ["/entrypoint"]
