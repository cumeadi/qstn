exports.get = (key)->
  key = "#{key}="
  str = document.cookie / \;
  for v, i in str
    while v.charAt(0) == ' ' then v = v.substring 1, v.length
    if v.indexOf(key) == 0 then return v.substring key.length, v.length
  null

exports.set = (key, val, ttl)!->
  ttl = if ttl then +new Date + ttl else null
  ttl = if ttl then ";expires=#ttl" else ''
  document.cookie = "#{key}=#{val}#{ttl};path=/"

