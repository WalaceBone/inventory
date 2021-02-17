FROM golang:1.14-alpine


WORKDIR /app
COPY server .

CMD ["/app/main"]
