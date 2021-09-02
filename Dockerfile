FROM golang:latest
ENV GOPATH=/
COPY ./ ./
RUN go mod download
RUN go build -o main ./cmd/app/main.go

CMD ["./main"]