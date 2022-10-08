
FROM golang:1.19 as build
WORKDIR /workspace
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /workspace/pizza-apiserver /
ENTRYPOINT ["/pizza-apiserver"]
