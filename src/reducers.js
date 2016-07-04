
import {combineReducers} from 'redux';
import mux from './lib/mux';

/**
 * Set next route
 *
 * @param {Object} state
 * @param {Object} action
 * @return {Mixed}
 * @api public
 */

const route = async (state={}, {type,path}) => {
  if('SET_ROUTE' != type) return state;

  return {
    route: mux.fetch(value),
    froze: true,
  };
};

/**
 * Expose `combineReducers()`
 */

export default combineReducers({
  Route: route,
});

