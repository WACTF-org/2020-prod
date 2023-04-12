FROM httpd:2.4-alpine

# Add a low priv user.
RUN adduser -S -s /bin/false -H -D user

# Install apache and sudo
# RUN apk add apache2 apache2-ctl sudo

# Name the server
RUN echo "ServerName wactf" >> /usr/local/apache2/conf/httpd.conf

# Disable directory listing
RUN sed -i 's/Options Indexes FollowSymLinks/Options -Indexes +FollowSymLinks/' /usr/local/apache2/conf/httpd.conf

# Configure user agent blacklist 
RUN echo $'SetEnvIfNoCase User-Agent "^gobuster" bad_tool\n\
    SetEnvIfNoCase User-Agent "^dirb" bad_tool\n\
    SetEnvIfNoCase User-Agent "^curl" bad_tool\n\
    # SetEnvIfNoCase User-Agent "^wget" bad_tool\n\
    <Directory "/var/www">\n\
    Order Allow,Deny\n\
    Allow from all\n\
    Deny from env=bad_tool\n\
    </Directory>' > /etc/conf.d/blacklist.conf

# Whitelist the sudo command to allow user to run apache, but nothing else as root.
# RUN echo "user ALL = (ALL) NOPASSWD: /usr/sbin/apachectl -D FOREGROUND" > /etc/sudoers.d/user

# Copy content to webroot
COPY ./docker-files/ /usr/local/apache2/htdocs/
ADD ./docker-files/git.tar /usr/local/apache2/htdocs/env/

RUN chown -R user:nobody /usr/local/apache2/logs/

# Run apache
# CMD /usr/bin/sudo /usr/sbin/apachectl -D FOREGROUND
CMD httpd -D FOREGROUND

# Drop privs
USER user