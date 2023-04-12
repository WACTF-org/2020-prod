# Example Dockerfile that uses sudo to run a service that needs to bind to a privileged port
# This is accepted for services with known-good sandboxing, like Apache/nginx.
# This is less accepted for random shit.

# Use alpine for its efficiency and small attack surface.
FROM httpd:2.4-alpine

# Add a low priv user.
RUN adduser -S -s /bin/false -H -D user

# Install apache, php, sudo, sqlite, php7 and required php7 modules.
RUN apk update
RUN apk add --no-cache apache2 apache2-ctl php7-apache2 sudo pwgen php7 php7-curl php7-intl php7-json php7-mbstring php7-mysqlnd php7-xml php7-iconv php7-phar php7-xmlwriter php7-xmlreader php7-tokenizer php7-session php7-sqlite3 php7-sodium php7-simplexml php7-posix php7-pdo_sqlite php7-pdo_mysql php7-pdo php7-mysqli php7-ftp php7-fileinfo php7-ctype sqlite curl git wkhtmltopdf

# Copy warez into the web root
COPY ./challenge/src/ /var/www/localhost/htdocs/
COPY ./challenge/conf/httpd.conf /usr/local/apache2/conf/httpd.conf
COPY ./challenge/conf/php.ini /usr/local/etc/php/php.ini

# Set permissions on the web directory content
RUN chown -R apache:apache /var/www/localhost/htdocs/
RUN chmod 770 /var/www/localhost/htdocs/db/secure-books.db

# Uncomment next 2 lines if you are using an index.php
RUN rm -f /var/www/localhost/htdocs/index.html

# Composer
RUN curl -sS https://getcomposer.org/installer | php
RUN mv composer.phar /usr/local/bin/composer
RUN chmod +x /usr/local/bin/composer
RUN composer self-update

RUN cd /var/www/localhost/htdocs/ && composer install 

RUN chown -R apache:apache /var/www/logs/
# CMD httpd -D FOREGROUND