FROM golang:latest 

WORKDIR /app

EXPOSE 9091

COPY table-booking /app

COPY dev.env /app

ENTRYPOINT [ "/app/table-booking"]

