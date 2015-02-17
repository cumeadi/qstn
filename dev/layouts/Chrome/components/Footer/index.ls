require! 'react' : React
require! 'react-router' : Router
require! 'api/Entry'
require  './index.scss'

e = React.createElement
d = React.DOM

{Navigation} = Router

λ = React.createClass do
  displayName: 'Footer'

  mixins: [Navigation]

  random: (e) ->
    e.preventDefault!
    Entry.random!.then ((res) !->
      @transitionTo 'show',
        id: res.slug
    ).bind @

  render: ->
    d.footer do
      className: 'chrome-footer'
      role: 'content-info'
      d.div do
        className: 'wrap'
        d.p do
          className: 'copy'
          '\u00A9 2015 qstn'
        d.ul do
          className: 'links'
          d.li {}, d.a do
            onClick: @random
            'Random poll'
          d.li {}, d.a do
            href: '//github.com/daryl/qstn'
            'Source'

module.exports = λ
