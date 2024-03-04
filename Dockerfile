FROM golang:1.21 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -C cmd/app/ -o ../../temperature-app

FROM scratch
WORKDIR /app
COPY --from=build /app/temperature-app .
COPY --from=build /app/.env .
EXPOSE 8080
ENTRYPOINT [ "./temperature-app" ]
