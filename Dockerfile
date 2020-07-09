#build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache git
RUN go build -o algorithms

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app /app
CMD [ "/app/algorithms" ]
LABEL Name=go-algorithms Version=0.0.1
