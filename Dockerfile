FROM golang:1.18-alpine as builder
WORKDIR /usr/app
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/bookshelf .
# RUN go build .

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /usr/app/build/bookshelf /usr/bin/bookshelf
# COPY --from=builder /usr/app/bookshelf /usr/bin/bookshelf
EXPOSE 4000
ENTRYPOINT ["/usr/bin/bookshelf"]
