FROM golang as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server /app/main.go \
    && chmod a+x /app/server

FROM scratch
COPY --from=builder /app/server .

EXPOSE 9998
ENTRYPOINT ["./server"]