FROM golang:alpine AS builder

WORKDIR /opt/server
COPY ./docker-files/main.go .
COPY ./docker-files/go.mod .
COPY ./docker-files/go.sum .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -o /go/bin/ggrf

FROM alpine:3.12.1

RUN apk add --update --no-cache \
    libgcc libstdc++ libx11 glib libxrender libxext libintl \
    ttf-dejavu ttf-droid ttf-freefont ttf-liberation ttf-ubuntu-font-family

ADD /docker-files/bins/wkhtmltopdf.tar.gz /tmp/wkhtmltopdf
RUN mv /tmp/wkhtmltopdf/wkhtmltopdf /bin/wkhtmltopdf && chmod +x /bin/wkhtmltopdf && rmdir /tmp/wkhtmltopdf

RUN addgroup -S mygroup && adduser --system --shell /bin/nologin --no-create-home -G mygroup --disabled-password myuser
USER myuser

WORKDIR /var/www/

COPY --chown=myuser:mygroup /docker-files/html /var/www/html/
COPY --chown=myuser:mygroup /docker-files/serve /var/www/serve/
COPY --chown=myuser:mygroup /docker-files/template /var/www/template/

COPY --from=builder --chown=myuser:mygroup /go/bin/ggrf /var/www/ggrf

EXPOSE 8080

ENTRYPOINT ["/var/www/ggrf"]