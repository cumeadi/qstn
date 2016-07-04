
import Html from './src/components/Html';
import App from './src/components/App';
import {element,string} from 'deku';
import mux from './src/lib/mux';

/**
 * Render app
 *
 * @param {Request} req
 * @return {String}
 * @api public
 */

export default async ({url}) => {
  let node = <Html><App/></Html>;
  let html = '<!doctype html>';

  try {
  html += string.render(node);
  } catch(e) {
    console.log(e.stack);
    //console.log(e);
  }
  return html;
};

