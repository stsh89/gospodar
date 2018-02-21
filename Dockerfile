FROM golang
WORKDIR /src
EXPOSE 3000
CMD ["go", "run", "main.go"]
