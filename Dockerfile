FROM golang
WORKDIR /go/src/gospodar
COPY . .
RUN apt-get update && apt-get install -y \
  postgresql-client
RUN go get -d ./... && go get -u github.com/pressly/goose/cmd/goose
EXPOSE 3000
