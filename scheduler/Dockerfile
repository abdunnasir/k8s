FROM golang:1.16-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o my-custom-scheduler main.go

FROM alpine
COPY --from=builder /app/my-custom-scheduler /my-custom-scheduler
ENTRYPOINT ["/my-custom-scheduler"]

