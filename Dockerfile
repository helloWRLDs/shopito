FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux go build -o ./users ./services/users/cmd/app/main.go
RUN GOOS=linux go build -o ./notifier ./services/notifier/cmd/app/main.go
RUN GOOS=linux go build -o ./products ./services/products/cmd/app/main.go
RUN GOOS=linux go build -o ./api_gw ./services/api-gw/cmd/app/main.go


FROM alpine:latest AS users-service
WORKDIR /app
COPY --from=builder /app/users .
COPY --from=builder /app/.env .
RUN chmod +x ./users
EXPOSE 3002
CMD [ "./users" ]


FROM alpine:latest as notifier-service
WORKDIR /app
COPY --from=builder /app/notifier .
COPY --from=builder /app/.env .
RUN chmod +x ./notifier
EXPOSE 3003
CMD [ "./notifier" ]

FROM alpine:latest as products-service
WORKDIR /app
COPY --from=builder /app/products .
COPY --from=builder /app/.env .
RUN chmod +x ./products
EXPOSE 3007
CMD [ "./products" ]

FROM alpine:latest as api-gateway
WORKDIR /app
COPY --from=builder /app/api_gw .
COPY --from=builder /app/.env .
RUN chmod +x ./api_gw
EXPOSE 3000
CMD [ "./api_gw" ]