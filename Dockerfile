FROM golang:alpine

WORKDIR /var/www/html/apps/open-api
COPY . /var/www/html/apps/open-api

RUN go build -o main .

CMD ["/var/www/html/apps/open-api/main"]