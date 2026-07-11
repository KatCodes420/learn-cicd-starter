FROM --platform=linux/amd64 debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /app

COPY notely /app/notely
COPY static /app/static

CMD ["./notely"]