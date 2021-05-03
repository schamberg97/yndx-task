const serverless = require('serverless-http')
const restana = require('restana')

const app = restana()
app.get('/*', (req, res) => {
  //res.send((Math.floor(Date.now() / 1e3)).toString())
  res.send('boo')
})

const handler = serverless(app);
module.exports.handler = async (event, context) => {
  return await handler(event, context)
}