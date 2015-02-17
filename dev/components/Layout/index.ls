require! 'react' : React
require! 'layouts'

e = React.createElement
d = React.DOM

λ = React.createClass do
  displayName: 'Layout'

  componentDidMount:  !-> document.title = "
    #{@props.title} | qstn
  "

  componentDidUpdate: !-> @componentDidMount!

  render: ->
    e layouts[@props.name], @props

module.exports = λ
