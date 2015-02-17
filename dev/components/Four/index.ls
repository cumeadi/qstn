require! 'react' : React
require! 'react-router' : Router
require  './index.scss'

e = React.createElement
d = React.DOM

{Link} = Router

λ = React.createClass do
  displayName: 'Four'

  componentDidWillMount: ->
    console.log 'hi'

  render: ->
    d.div do
      className: 'four'
      'Nothing here.'
      ' '
      e Link,
        to: 'index'
        'Go home'

module.exports = λ
