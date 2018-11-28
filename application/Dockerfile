FROM golang:1.8

RUN mkdir -p /application
WORKDIR /application

EXPOSE 8080
COPY ./gotasker .

ENTRYPOINT ["./gotasker"]