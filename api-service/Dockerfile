FROM golang:1.18.3-alpine3.16

WORKDIR /app

COPY api-service/go.mod .
COPY api-service/go.sum .
RUN go mod download

COPY ./api-service .

RUN go build -o /goapp ./cmd

EXPOSE 8001

CMD ["/goapp"]