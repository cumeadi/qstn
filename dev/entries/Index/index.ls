require! 'react' : React
require! 'react-router' : Router
require! 'components/Layout'
require! 'api/Entry' : Entry
require  './index.scss'

e = React.createElement
d = React.DOM

{Navigation} = Router

λ = React.createClass do
  displayName: 'Index'

  mixins: [Navigation]

  getInitialState: -> {
    question: ''
    options:
      * option: ''
      * option: ''
    -sending
  }

  componentDidMount: !->
    @refs.question.getDOMNode!.focus!

  questionChange: (e) !->
    @setState question: e.target.textContent

  optionChange: (idx, e) !->
    n = 0; o = @state.options
    o[idx].option = e.target.textContent
    for v in o then n++ if /\S/ is v.option
    switch true
    | n == o.length
      o.push option: ''
    | n + 2 == o.length != 2
      o.pop!
    @setState do
      options: o

  handleSubmit: (e) !->
    e.preventDefault!
    data = @state
    data.options.pop!
    console.log data
    Entry.post data .then ((res) !->
      @transitionTo 'show', id: res.slug
    ).bind @

  render: ->
    can = /\S/ is @state.question
    and (@state.options.filter (v) ->
      /\S/ is v.option
    ).length > 1

    e Layout,
      name: 'chrome'
      title: 'Create polls with real-time results'
      d.form do
        onSubmit: @handleSubmit
        className: 'entry-index'
        d.div do
          ref: 'question'
          className: 'question'
          contentEditable: true
          onInput: @questionChange
          placeholder: 'What would \
          you like to ask?'
        d.div do
          className: 'options'
          @state.options.map ((v, i) ->
            d.div do
              className: 'option'
              contentEditable: true
              onInput: @optionChange.bind @, i
              placeholder: "Option #{i+1}"
              key: i
          ).bind @
        d.button do
          disabled: !can
          'Ask'



module.exports = λ
