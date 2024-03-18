FROM golang:1.21 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -C cmd/cep-web-service/ -o ../../cep-app

FROM scratch
WORKDIR /app
COPY --from=build /app/cep-app .
EXPOSE 8080
ENTRYPOINT [ "./cep-app" ]
