FROM golang:1.17.8

WORKDIR /Belajar-Go-Echo
COPY . .
RUN go mod download
ADD main.go /Belajar-Go-Echo/
EXPOSE 8080
CMD ["go", "run", "main.go"]