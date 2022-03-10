FROM golang:latest

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN apt update && \
    apt install build-essential librdkafka-dev -y

CMD ["tail", "-f", "/dev/null"]
