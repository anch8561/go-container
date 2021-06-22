FROM docker.io/alpine
RUN echo 'hello world' > msg
CMD cat msg
