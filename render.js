
import Html from './src/components/Html';
import App from './src/components/App';
import {element,string} from 'deku';

/**
 * Render app
 *
 * @param {Request} req
 * @return {String}
 * @api public
 */

export default async ({url}) => {
  let html = '<!doctype html>';
  html += string.render(<Html/>);
  return html;
};

