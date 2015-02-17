require! 'react' : React
require! 'components/Layout'
require! 'components/Four'
require! './components/Option'
require! 'api/Entry' : Entry
require! 'stores/Entry' : EntryStore
require! 'utils/cookie'
require! 'utils/sock'
require  './index.scss'

e = React.createElement
d = React.DOM

_sock  = null

λ = React.createClass do
  displayName: 'Show'

  statics: {
    resolve: (params) ->
      Entry.get do
        params.id

    willTransitionTo: (t, params) !->
      {id} = params
      _sock := new sock id
      _sock.listen do
        @recieveData

    willTransitionFrom: !->
      _sock.kill!
  }

  getInitialState: ->
    {id} = @props.params
    selected: if v = cookie.get id then Number v
    entry: EntryStore.get id

  componentWillMount: !->
    {id}  = @props.params
    _sock := new sock id
    _sock.listen do
      @recieveData

  componentWillReceiveProps: (props) !->
    {params} = props; @setState {
      selected:
        if v = cookie.get params.id
        then Number v
      entry: EntryStore.get do
        params.id
    }

  recieveData: (data) ->
    return if data.ping
    {id} = @props.params
    {selected} = @state
    entry = @state.entry
    entry.options = data.options
    @setState entry: entry
    old = Number cookie.get id
    if old != selected
    then cookie.set do
      id, selected

  optionChange: (idx) !->
    {entry}   = @state
    {options} = entry
    old       = @state.selected
    options[old].votes-- if !isNaN old
    options[idx].votes++
    entry.options = options
    @setState do
      selected: idx,
      entry: entry
    _sock.send do
      entry

  render: ->
    {entry} = @state

    return e Four unless entry

    total = entry.options.reduce (t, b) ->
      t + b.votes
    , 0

    e Layout,
      name: 'chrome'
      title: entry.question
      d.article do
        className: 'entry-show'
        d.h1 {}, entry.question
        d.div do
          className: 'options'
          entry.options.map ((v, i) ->
            e Option,
              key: i, idx: i
              selected: @state.selected
              onChange: @optionChange
              total: total
              option: v
          ).bind @
        d.div do
          className: 'actions'
          d.div do
            className: 'social'
            d.span {},
              'Share on '
              d.a do
                href: "
                  //twitter.com/intent/tweet?
                  text=#{encodeURIComponent entry.question}
                  +http://qstn.co/q/#{entry.slug}"
                'Twitter'
              ' or '
              d.a do
                href: "
                  //facebook.com/sharer/sharer.php?
                  u=http://qstn.co/q/#{entry.slug}"
                'Facebook'
          d.div do
            className: 'total'
            total

module.exports = λ
