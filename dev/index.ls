require! 'q' : Q
require! 'react' : React
require! 'react-router': Router
require! 'events/Loading'
require  'styles/defs.scss'
require! 'routes'

e = React.createElement
d = React.DOM

{HistoryLocation} = Router

fetch = (s) ->
  r = s.routes\
  .filter (r) ->
    r.handler.resolve
  .map (r) ->
    r.handler.resolve do
      s.params
  Q.all r

Router.run do
  routes
  HistoryLocation
  (H, s) ->
    Loading.emit 'start'
    fetch s .then (data) !->
      Loading.emit 'end'
      ga 'send', 'pageview', s.path
      React.render e(H, {
        params: s.params
      }), document.body


