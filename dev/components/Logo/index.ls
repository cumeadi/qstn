require! 'react' : React
require! 'react-router': Router
require  './index.scss'

e = React.createElement
d = React.DOM

{Link} = Router

λ = React.createClass do
  displayName: 'Logo'

  render: ->
    d.h1 do
      className: 'logo'
      e Link, @props

module.exports = λ
