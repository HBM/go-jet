
var WebSocket = require('ws')
var ws = new WebSocket('ws://localhost:12345/')

ws.on('open', function () {
  ws.send('something', function (error) {
    console.log(error)
  })
})

ws.on('message', function(data, flags) {
  console.log(data)
  console.log(flags)
})

ws.on('error', function (error) {
  console.log(error);
})
