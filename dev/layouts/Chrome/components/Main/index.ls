require! 'react' : React
require  './index.scss'

e = React.createElement
d = React.DOM

λ = React.createClass do
  displayName: 'Main'

  render: ->
    d.main do
      className: 'chrome-main'
      role: 'main'
      d.div do
        className: 'wrap'
        @props.children

module.exports = λ

