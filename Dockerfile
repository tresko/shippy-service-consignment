FROM alpine:latest

WORKDIR /app
ADD shippy-service-consignment /app/consignment-service

CMD ["./consignment-service"]