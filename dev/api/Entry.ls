require! 'q' : Q
require! 'actions/Entry' : EntryActions
require! 'utils/ajax'

exports.get = (id) ->
  d = Q.defer!; u = "/q/#id"
  ajax.get u .then (res) !->
    unless res.json.error?
      EntryActions.add res.json
    d.resolve res.json
  d.promise

exports.post = (data) ->
  d = Q.defer!; u = "/q/"
  ajax.post u, data .then (res) !->
    d.resolve res.json
  d.promise

exports.random = ->
  d = Q.defer!; u = "/q/random"
  ajax.get u .then (res) !->
    unless res.json.error?
      EntryActions.add res.json
    d.resolve res.json
  d.promise
