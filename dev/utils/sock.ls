{host} = location

sock = module.exports = (path) !->
  @_ws = new WebSocket "ws://#host/s/#path"
  @_ws.onmessage = recieve.bind @
  @_fn = []

sock::send = (data) !->
  data = JSON.stringify data
  @_ws.send data

sock::listen = (fn) !->
  @_fn.push fn

sock::kill = !->
  @_ws.close!

recieve = (res) !->
  data = JSON.parse do
    res.data
  for fn in @_fn
    fn data
