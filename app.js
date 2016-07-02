
import render from './render';
import {Server} from 'http';
import mux from './lib/mux';
import r from './routes';
import mime from 'mime';
import fs from 'fs';

/**
 * Expose `Server()`
 */

export default Server((...args) => app(...args));

/**
 * Public dir
 */

export let PUBLIC = './static';

/**
 * Set routes
 */

mux.get('/entries/:id', r.get);
mux.post('/entries', r.create);

/**
 * Application
 *
 * @param {Object} req
 * @param {Object} res
 * @api public
 */

const app = async (req, res) => {
  let [code,body] = file(req, res);
  let ajax = req.headers['X-QN'];

  if(body) {
    send(res, code, body);
    return
  }

  // Render Deku application
  if('undefined' == typeof ajax) {
    res.setHeader('Content-Type', 'text/html');
    body = await render(req);
    send(res, 200, body);
    return
  }

  let [fn,obj] = mux.fetch(req);

  // No route has been matched
  if('undefined' == typeof fn) {
    send(res, 404);
    return
  }

  try {
    [code,body] = await fn(req, obj);
  } catch({message}) {
    body = {body: message, code: 500};
    code = 500;
  }

  send(res, code, {
    body,
    code,
  });
};

/**
 * Serve files
 *
 * @param {Object} req
 * @return {Array}
 * @api private
 */

const file = (req, res) => {
  if(0 > req.url.indexOf('.')) return [];

  const path = PUBLIC + req.url;

  try {
    fs.lstatSync(path)
  } catch(e) {
    return [];
  }

  const body = fs.readFileSync(path);
  const type = mime.lookup(path);

  res.setHeader('Content-Type', type);

  return [200, body];
};

/**
 * Render JSON
 *
 * @param {Response} res
 * @param {Number} code
 * @param {Object} obj
 * @api public
 */

function json(res, code, obj={}) {
  res.setHeader('Content-Type', 'application/json');
  const body = JSON.stringify(obj);
  send(res, code, body);
}

/**
 * Send request
 *
 * @param {Response} res
 * @param {Number} code
 * @param {String} body
 * @api public
 */

function send(res, code, body) {
  res.writeHead(code);
  res.end(body);
}

