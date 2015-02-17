require! 'Dispatcher'

exports.add = (data) !->
  Dispatcher.serverAction do
    action: 'ADD_ENTRY'
    data: data
