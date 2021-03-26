FROM alpine
ADD borrowLogs-service /borrowLogs-service
ENTRYPOINT [ "/borrowLogs-service" ]
