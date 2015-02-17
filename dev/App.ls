require! 'react' : React
require! 'react-router': Router
require! 'events/Loading'

e = React.createElement
d = React.DOM

{RouteHandler} = Router

λ = React.createClass do
  displayName: 'App'

  getInitialState: ->
    loading: false

  componentDidMount: !->
    Loading.on 'start', (!->
      @setState {+loading}
    ).bind @

    Loading.on 'end', (!->
      @setState {-loading}
    ).bind @

  render: ->
    d.div do
      className:
        if @state.loading
        then 'loading'
      e RouteHandler,
        @props

module.exports = λ
