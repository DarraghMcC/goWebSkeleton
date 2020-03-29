FROM golang:1-alpine AS builder
COPY . /app/webSkeleton
WORKDIR /app/webSkeleton
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates
RUN adduser -D -g '' appuser
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.BuildDate=$(date)' -X 'main.GitCommit=$(git rev-list -1 HEAD)'" -a -o /webSkeleton .

FROM scratch
COPY --from=builder /webSkeleton ./
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

USER appuser
ENTRYPOINT ["./webskeleton"]
