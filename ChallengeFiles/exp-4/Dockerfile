FROM python:3-alpine

# ROOT
RUN adduser -S -s /bin/false -H -D user
RUN mkdir /server
COPY docker-files /server
RUN mv /server/aes.py /opt
RUN echo "WACTF{eval('Your life decisions if you use eval()')}" > /thisistheflag.txt

#USER
WORKDIR /server/
USER user
CMD /usr/local/bin/python /server/serv.py