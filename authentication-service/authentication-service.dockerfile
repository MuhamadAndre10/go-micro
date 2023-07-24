FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY authApp .

RUN chmod +x authApp

CMD [ "./authApp" ]
