FROM golang:1.24.1

WORKDIR /app
RUN useradd -r -d /app appUser

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY controllers/ ./controllers/
COPY database/ ./database/
COPY models/ ./models/
COPY server.go .

RUN go build -o server && RUN chown -R appUser /app
USER appUser

EXPOSE 1323

CMD ["./server"]
