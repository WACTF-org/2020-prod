FROM node:11-alpine

RUN addgroup -S mygroup && adduser --system --shell /bin/nologin --no-create-home -G mygroup --disabled-password myuser

RUN mkdir -p /usr/src/app

WORKDIR /usr/src/app

COPY ./docker-files/ .

RUN npm install

RUN mkdir -p /usr/src/app/log &&  chown -R myuser:mygroup /usr/src/app/log

EXPOSE 3000

CMD ["npm", "run", "start", "&"]

USER myuser
