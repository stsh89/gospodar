FROM golang
WORKDIR /go/src/gospodar
COPY . .
RUN go get -d ./...
RUN go get -u github.com/pressly/goose/cmd/goose
EXPOSE 3000
CMD ["go", "run", "main.go"]
