FROM alpine:latest

RUN mkdir /app
WORKDIR /app
ADD borrowLog /app/borrowLog

CMD ["./borrowLog"]