FROM golang:1.16-alpine AS builder

WORKDIR /app

ENV GO111MODULE=on
RUN go get github.com/cespare/reflex

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./service_content_web /app/web/main.go
RUN go build -o ./service_content_cmd /app/services/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/service_content_web .

EXPOSE 8080
CMD [ "./service_content" ]

