
FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY brokerApp .

RUN chmod +x brokerApp

CMD [ "./brokerApp" ]
