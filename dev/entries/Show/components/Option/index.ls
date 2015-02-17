require! 'react' : React
require! 'components/Radio'
require  './index.scss'

e = React.createElement
d = React.DOM

λ = React.createClass do
  displayName: 'Option'

  optionChange: !->
    @props.onChange do
      @props.idx

  render: ->
    {option} = @props

    perc = (
      option.votes / @props.total * 100 \
    || 0).toFixed 0

    d.div do
      className: 'show-option'
      e Radio,
        id: "option-#{@props.idx}"
        label: option.option
        selected: @props.selected == @props.idx
        onChange: @optionChange
        name: 'option'
      d.div do
        className: 'meta'
        d.div do
          className: 'perc'
          "#perc%"
        d.div do
          className: 'votes'
          option.votes
      d.div do
        className: 'bar'
        d.div style:
          width: "#perc%"

module.exports = λ
