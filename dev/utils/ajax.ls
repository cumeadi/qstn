require! 'q' : Q

send = (type, obj) ->
  d = Q.defer!
  x = new XMLHttpRequest
  x.open type, obj.url
  x.setRequestHeader 'X-QSTN', 'true'
  x.onreadystatechange = !->
    if x.readyState == 4
      x.json = do
        if x.response != ''
        then JSON.parse x.response
        else null
      d.resolve x
  x.send do
    if obj.data
    then JSON.stringify obj.data
    else null
  d.promise

exports.get = (url) ->
  send 'GET', url: url

exports.post = (url, data) ->
  send 'POST', data: data, url: url

exports.patch = (url, data) ->
  send 'PATCH', data: data, url: url

exports.put = (url, data) ->
  send 'PUT', data: data, url: url

exports.del = (url) ->
  send 'DELETE', url: url
