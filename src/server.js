
import Html from './components/App';
import App from './components/App';
import {string} from 'deku';
import mux from './lib/mux';

const render = async (path) => {
  const tree = <Html><App/></Html>;
  const html = '<!doctype html>';
  html += string.render(tree, {});
  return html;
};

