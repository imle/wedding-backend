FROM golang:1.15-alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY cmd cmd
COPY ent ent
COPY pkg pkg
COPY internal internal
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bin/app ./cmd/server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/bin .
COPY migrations migrations
ENTRYPOINT ["./app"]
