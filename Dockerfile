FROM golang:1.19.1-alpine3.16 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o VUapp main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/VUapp /app
COPY --from=builder /app/app.env .
EXPOSE 9095
CMD ["/app/VUapp"]