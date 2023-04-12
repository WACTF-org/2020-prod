
const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');
const { log, ExpressAPILogMiddleware } = require('@rama41222/node-logger');

const config = {
    name: 'sample-express-app',
    port: 3000,
    host: '0.0.0.0',
};

const app = express();
const logger = log({ console: true, file: false, label: config.name });


app.use(bodyParser.json());
app.use(cors());
app.use(require('apikey')(auth, 'my realm'));
app.use(ExpressAPILogMiddleware(logger, { request: true }));

function auth (key, fn) {
  if ('SuperSecurePasswordforuseradmin' === key)
    fn(null, { id: '1', name: 'superuser'})
  else
    fn(null, null)
}

app.get('/', function(req,res) {
  res.status(200)
  res.send('success!!');
});

app.get('/flag', function(req,res) {
 res.sendFile(__dirname+"/views/flag.html");
});

app.listen(config.port, config.host, (e)=> {
    if(e) {
        throw new Error('Internal Server Error');
    }
    logger.info(`${config.name} running on ${config.host}:${config.port}`);
});
