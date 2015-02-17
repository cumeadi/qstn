require! 'react' : React
require! './components/Header'
require! './components/Footer'
require! './components/Main'
require  './index.scss'

e = React.createElement
d = React.DOM

λ = React.createClass do
  displayName: 'Chrome'

  render: ->
    d.div do
      className: 'chrome'
      e Header
      e Main, @props
      e Footer

module.exports = λ
