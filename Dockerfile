FROM golang:1.20-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/main .
COPY --from=build /app/dist ./dist

EXPOSE 42069

CMD ["./main"]
