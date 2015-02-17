require! 'eventemitter3' : Events
require! 'object-assign' : assign
require! 'Dispatcher'

entries = {}

λ = assign Events::,
  get: (id) ->
    entries[id]

  emitChange: !->
    @emit 'change'

  onChange: (fnc) !->
    @on 'change', fnc

  off: (fnc) !->
    @off 'change', fnc

Dispatcher.register (load) !->
  action = load.action
  data   = action.data

  switch action.action
  | 'ADD_ENTRY'
    entries[data.slug] = data
  | otherwise
    return

  λ.emit 'change'

module.exports = λ
