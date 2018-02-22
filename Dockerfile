FROM golang
WORKDIR /go/src
EXPOSE 3000
CMD ["go", "run", "main.go"]
