FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN go build -o app.bin

FROM gcr.io/distroless/static-debian12 AS runtime
WORKDIR /app
CMD ["/app/app"]
ENV PORT=80
EXPOSE 80
COPY templates/ /app/templates/
COPY --from=builder /app/app.bin /app/app.bin
