FROM alpine
RUN echo 'hello world' > msg
CMD cat msg