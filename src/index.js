
import {createApp,element} from 'deku';
import App from './components/App';
import {createStore} from 'redux';
import reducer from './reducer';

/**
 * Redux store
 */

const {dispatch,getState} = createStore(reducer);

/**
 * Render fn
 */

const render = createApp(document.body, dispatch);

/**
 * Render <App/>
 */

render(<App/>, getState);

