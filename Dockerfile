FROM golang:1.20-alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
WORKDIR github.com/puneet105/job-queue
COPY . .
RUN rm -rf go.* && \
    go mod init github.com/puneet105/job-queue && \
    go mod tidy 
RUN go build -o /github.com/puneet105/job-queue


FROM alpine:latest
WORKDIR /root/
COPY --from=builder /github.com/puneet105/job-queue .
COPY config.yaml .
EXPOSE 8080
CMD ["./job-queue"]
