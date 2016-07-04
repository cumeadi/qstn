
import {send,json,file} from './util';
import render from './render';
import {Server} from 'http';
import mux from './lib/mux';
import r from './routes';
import mime from 'mime';
import zlib from 'zlib';
import fs from 'fs';

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

export default async (req, res) => {
  let [body,type] = file(PUBLIC + req.url);
  let ajax = req.headers['X-QN'];
  let code = 200;

  if(body) {
    res.setHeader('Content-Type', type);
    send(res, code, body);
    return
  }

  // Render Deku application
  if('undefined' == typeof ajax) {
    res.setHeader('Content-Type', 'text/html');
    body = await render(req);
    send(res, code, body);
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

