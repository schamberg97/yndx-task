const http = require('http');

const port = require(require('path').resolve(__dirname, '../nodeCommon/port.js'))

const handler = function (req, res) {
  res.writeHead(200);
  const time = Math.floor(Date.now() / 1e3);
  res.end(time.toString());
}

const server = http.createServer(handler);
server.listen(port);