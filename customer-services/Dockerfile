FROM golang:1.15-alpine3.12 AS build
RUN apk add --no-cache git
WORKDIR /app/customer-services
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/customer-services .
EXPOSE 5005
CMD ["./out/customer-services"]