FROM php:7.2-apache

RUN mv "$PHP_INI_DIR/php.ini-production" "$PHP_INI_DIR/php.ini"

COPY ./docker-files/ /var/www/html

RUN echo "ServerName wactf" >> /etc/apache2/apache2.conf

RUN chown -R daemon:daemon /var/www/html/

RUN rm -f /var/www/index.html

