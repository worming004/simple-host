FROM docker.io/golang:1.20 as builder
WORKDIR /app
COPY go.mod go.mod
RUN go mod download
COPY app.go .
RUN go build -o /app/app ./app.go

FROM docker.io/ubuntu:latest
COPY --from=builder /app/app /bin/app 
EXPOSE 8085/tcp
CMD ["/bin/app"]

