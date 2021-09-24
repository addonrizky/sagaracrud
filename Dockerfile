FROM golang as builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine
WORKDIR /app

COPY --from=builder /app/main .
RUN mkdir logs && chmod -R 777 logs && mkdir -p config && mkdir image_product && chmod -R 777 image_product
#COPY config/config.json config
# COPY nsswitch.conf /etc/

ENV MYSQL_DATABASE=sagaracrud \
    MYSQL_PASSWORD=docker \
    MYSQL_HOST=fullstack-mysql \
    MYSQL_USER=docker \
    MYSQL_PORT=3306

EXPOSE 4000
CMD ["./main"]