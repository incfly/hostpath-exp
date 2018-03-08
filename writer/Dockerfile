FROM ubuntu:16.04

# https://github.com/golang/go/wiki/Ubuntu
RUN apt-get update
# RUN apt-get install -y python-software-properties apt-file
RUN apt-get install -y python-software-properties
RUN apt-get install -y software-properties-common
RUN add-apt-repository ppa:gophers/archive
RUN apt-get install -y golang-1.9-go
ENV PATH="/usr/lib/go-1.9/bin:${PATH}"

# install curl
RUN apt-get install -y curl

# build go app.
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main
CMD ["/app/main"]
