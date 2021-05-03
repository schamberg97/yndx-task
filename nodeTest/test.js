const http = require('http')

const port = require(require('path').resolve(__dirname, '../nodeCommon/port.js'))

const options = {
  hostname: '127.0.0.1',
  port: port,
  path: '/',
  method: 'GET'
}

const req = http.request(options, res => {
  let data = '';
  console.log(`statusCode: ${res.statusCode}`)

  res.on('data', (chunk) => {
    data += chunk;
  });

  // The whole response has been received. Print out the result.
  res.on('end', () => {
    console.log(data);
  });
})

req.on('error', error => {
  console.error(error)
})

req.end()
