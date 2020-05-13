FROM golang:alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o syslogpinger .
ENTRYPOINT ["/app/syslogpinger"]
CMD []
