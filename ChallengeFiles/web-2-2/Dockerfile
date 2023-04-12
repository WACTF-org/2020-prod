# Rename to 'Dockerfile'
FROM alpine:latest
# allows to not cache the index locally, keeps containers small
# https://gist.github.com/sgreben/dfeaaf20eb635d31e1151cb14ea79048
RUN apk add --no-cache 
# add/del packages if necessary
# RUN apk add golang
# create a new user and group
RUN adduser -S -s /bin/false -H -D testyman
# if you don't need a user with any interactivity you can do something like this instead
#RUN addgroup -S mygroup && adduser --system --shell /bin/nologin --no-create-home -G mygroup --disabled-password myuser
COPY docker-files/app/ /app
# the directory container will start in
WORKDIR /app
# default command to run when container starts
RUN chmod +rx /app/toctou
RUN chown -R testyman:root /app
# RUN chmod -r /app/init.sql /app/main.go /app/build.txt
# RUN chown -R root:root /app/assets /app/html

USER testyman
EXPOSE 8000

CMD ["/app/toctou"]