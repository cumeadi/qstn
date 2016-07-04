
import {element} from 'deku';

/**
 * Render
 *
 * @param {Object} obj
 * @return {element}
 * @api public
 */

function render({route}) {
  if(!route) return <div/>;
  return element(route);
}

/**
 * Expose
 */

export default {
  render,
};

