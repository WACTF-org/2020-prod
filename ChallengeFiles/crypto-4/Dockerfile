FROM node:12-alpine3.12

WORKDIR /usr/src/app
COPY ./docker-files/ .

RUN npm install

EXPOSE 8000

USER nobody

CMD ["node", "index.js"]
