FROM golang:latest as builder

WORKDIR /mydir

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM nvidia/cuda:11.8.0-base-ubuntu22.04
COPY --from=builder /mydir/main .

CMD ["./main"]