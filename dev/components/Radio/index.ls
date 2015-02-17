require! 'react' : React
require  './index.scss'

e = React.createElement
d = React.DOM

λ = React.createClass do
  displayName: 'Radio'

  render: ->
    d.div do
      className: 'radio'
      d.input do
        type: 'radio'
        checked: @props.selected
        onChange: @props.onChange
        name: @props.name
        id: @props.id
      d.label do
        htmlFor: @props.id
        d.div do
          className: 'icon'
          d.i {}, d.i {}, ''
        @props.label

module.exports = λ
