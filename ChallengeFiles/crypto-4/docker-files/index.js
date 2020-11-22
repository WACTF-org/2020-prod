const crypto = require('crypto')
const express = require('express')
const mustache = require('mustache-express')

const app = express()
const port = 8000
const SECRET = Buffer.from('3b8d0f44-e52c-4d85-825a-436bc6f50ef3', 'utf-8')

app.set('views', `${__dirname}/views`)
app.set('view engine', 'mustache')
app.engine('mustache', mustache())
app.use(express.json())
app.use(express.static('public'))

const defaultCode = `function hello(name) {
  return 'Hello ' + name + '!';
}

hello('World'); // should print 'Hello World'`

const defaultSig = 'aaa8111b4871b48dc6c0ac4c33ef9e1b'

app.get('/', function (req, res) {
    res.render('index', { code: defaultCode, sig: defaultSig })
})

app.post('/', function (req, res) {
    const delim = Buffer.from('|', 'utf-8')
    const code = Buffer.from(req.body.code, 'base64')
    let result = ''

    const sig = crypto.createHash('md5').update(Buffer.concat([SECRET, delim, code])).digest('hex');
    if (sig === req.body.sig) {
        try {
            result = eval(code.toString())
        } catch(e) {
            result = e.message
            console.log(e)
        }
    } else {
        result = '<code signature invalid>'
    }
    res.send(JSON.stringify({ result }))
})

app.listen(port, function() {
    console.log('Server started on port ' + port)
})
