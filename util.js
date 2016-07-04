
import Stream from 'stream';
import mime from 'mime';
import zlib from 'zlib';
import fs from 'fs';

/**
 * Send JSON
 *
 * @param {Response} res
 * @param {Number} code
 * @param {Object} obj
 * @api public
 */

export function json(res, code, obj={}) {
  res.setHeader('Content-Type', 'application/json');
  const body = JSON.stringify(obj);
  send(res, code, body);
}

/**
 * Write request
 *
 * @param {Response} res
 * @param {Number} code
 * @param {Mixed} body
 * @api public
 */

export function send(res, code, body) {
  res.setHeader('Content-Encoding', 'gzip');
  res.writeHead(code);

  if(body instanceof Stream) {
    const gzip = zlib.createGzip();
    body.pipe(gzip).pipe(res);
    return;
  }

  body = zlib.gzipSync(body);

  res.end(body);
}

/**
 * Read file
 *
 * @param {String} path
 * @return {Array}
 * @api private
 */

export function file(path) {
  let body;
  let type;
  let st;

  try {
    st = fs.lstatSync(path);
  } catch(e) {
    return [];
  }

  if(st.isDirectory()) {
    return [];
  }

  body = fs.createReadStream(path);
  type = mime.lookup(path);

  return [body,type];
};

