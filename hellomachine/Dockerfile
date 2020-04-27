FROM golang as builder
WORKDIR /app
ADD hellomachine.go .
RUN go build -o hello-server hellomachine.go

FROM busybox:1.28.4
WORKDIR /app
EXPOSE 80
COPY --from=builder /app/hello-server hello-server
ENTRYPOINT ./hello-server --port=0.0.0.0:80
