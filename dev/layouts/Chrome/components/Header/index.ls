require! 'react' : React
require! 'react-router': Router
require! 'components/Logo'
require  './index.scss'

e = React.createElement
d = React.DOM

{Link} = Router

λ = React.createClass do
  displayName: 'Header'

  render: ->
    d.header do
      className: 'chrome-header'
      d.div do
        className: 'wrap'
        e Logo,
          to: 'index'
          'qstn'
        d.p {},
          'Create polls with \
          real-time results'


module.exports = λ
