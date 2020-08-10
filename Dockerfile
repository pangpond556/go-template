FROM golang:1.14 as build

WORKDIR /app
COPY . .
RUN apt update && apt install ca-certificates && update-ca-certificates
RUN go mod init your_app_name
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o ./server .

FROM scratch
COPY --from=build /app/server .
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["./server"]