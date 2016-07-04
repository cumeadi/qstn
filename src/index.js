
import {createApp,element} from 'deku';
import App from './components/App';
import {createStore} from 'redux';
import reducers from './reducers';
import debounce from 'debounce';
import ready from 'domready';

/**
 * Current route
 */

let _route = null;

/**
 * DOM node
 */

const DOMNode = document.body;

/**
 * Redux store
 */

const {dispatch,getState,subscribe} = createStore(reducers);

/**
 * Update handler
 */

const update = () => render(<App/>, getState());

/**
 * Render handler
 */

const render = createApp(DOMNode, dispatch);

/**
 * Redux listener
 */

subscribe(debounce(async () => {
  let state = getState();
  let {route} = state;
  let {fetch} = route;

  fetch = fetch || noop;

  if(route != _route) {
    _route = route;
  } else {
    return;
  }

  await fetch(dispatch);

  update();
}));

/**
 * Render
 */

onload = e => dispatch({
  value: location.pathname,
  type: 'SET_ROUTE',
});


