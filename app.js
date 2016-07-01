
import micro, {send,json} from 'micro';
import mux from './lib/mux';
import chalk from 'chalk';
import db from './lib/db';
import r from './routes';
import mime from 'mime';
import fs from 'fs';

/**
 * Public dir
 */

const PUBLIC = './static';

/**
 * Index tmpl
 */

const INDEX = fs.readFileSync('index.html', 'utf8');

/**
 * Set routes
 */

mux.get('/entries/:id', r.get);
mux.post('/entries', r.create);

/**
 * Serve files
 *
 * @param {Object} req
 * @param {Object} res
 * @param {Date} start
 * @return {Boolean}
 * @api private
 */

const file = async (req, res, start) => {
  if(0 > req.url.indexOf('.')) return;

  const path = PUBLIC + req.url;

  try {
    fs.lstatSync(path)
  } catch(e) {
    return false;
  }

  const body = fs.readFileSync(path);
  const type = mime.lookup(path);

  res.setHeader('Content-Type', type);
  res.writeHead(200);
  res.end(body);
  log(req, 200);

  return true;
};

/**
 * Error handler
 *
 * @param {Object} req
 * @param {Object} res
 * @param {Error} err
 */

const error = (req, res, {message}) => {
  const body = {body, code: 500};
  send(res, 500, body);
  log(req, 500);
};

/**
 * HTTP logs
 *
 * @param {Request} req
 * @param {Number} code
 * @api private
 */

function log({url,method,$st}, code) {
  let end = new Date - $st + 'ms';
  let out = ['--->', method];

  if(code < 400) {
    out.push(chalk.green(code));
  } else {
    out.push(chalk.red(code));
  }

  out = out.concat([
    chalk.grey(url),
    chalk.grey(end),
  ]).join(' ');

  console.log(out);
}

/**
 * Server
 *
 * @param {Object} req
 * @param {Object} res
 * @api public
 */

const server = async (req, res) => {
  req.$st = new Date;

  let {url,method,headers} = req;
  let ajax = headers['X-QN'];

  if(await file(req, res)) return;

  if('undefined' == typeof ajax) {
    send(res, 200, INDEX);
    log(req, 200);
    return;
  }

  let [fn,args] = mux.fetch(req);

  if('undefined' == typeof fn) {
    send(res, 404);
    log(req, 404);
    return;
  }

  let [code,body] = await fn(req, args);

  send(res, code, {code,body});
  log(req, code);
};

/**
 * Expose `micro()`
 */

export default micro(server, {
  onError: error,
});

