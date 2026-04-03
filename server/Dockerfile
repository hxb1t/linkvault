FROM golang:1.26.1 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o server ./cmd/api
RUN mkdir -p /data

FROM alpine:3.19
COPY --from=build /app/server /server
COPY --from=build /app/configs/.env /configs/.env
EXPOSE 8080
CMD ["/server"]
