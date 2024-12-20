FROM golang:1.23

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o genreid ./cmd/main.go

CMD ["./genreid"]