require! 'react' : React
require! 'object-assign' : assign
require! 'flux' : Flux

{Dispatcher} = Flux

λ = assign new Dispatcher,
  serverAction: (action) !->
    @dispatch do
      source: 'SERVER'
      action: action

  viewAction: (action) !->
    @dispatch do
      source: 'VIEW'
      action: action

module.exports = λ
