FROM scratch

ADD entrypoint /
EXPOSE 3000
RUN bash -c 'echo -e OH YEAH'
ENTRYPOINT ["/entrypoint"]
