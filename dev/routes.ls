require! 'react' : React
require! 'react-router' : Router

e = React.createElement
d = React.DOM

{Route, DefaultRoute, NotFoundRoute} = Router

module.exports = do
  e Route,
    path: '/'
    handler: require 'App'
    name: 'app'
    e DefaultRoute,
      handler: require 'entries/Index'
      name: 'index'
    e NotFoundRoute,
      handler: require 'components/Four'
    e Route,
      path: '/q/:id'
      handler: require 'entries/Show'
      name: 'show'

