
import ptx from 'path-to-regexp';

/**
 * Routes obj
 */

const obj = {
  GET: {},
  PATCH: {},
  DELETE: {},
  POST: {},
  PUT: {},
};

/**
 * Add route
 *
 * @param {String} meth
 * @param {String} patt
 * @param {Function} fn
 * @api private
 */

function add(meth, patt, fn) {
  obj[meth][patt] = fn;
}

/**
 * Fetch route
 *
 * @param {Request} req
 * @return {Array}
 * @api public
 */

export function fetch(req) {
  let {url,method} = req;
  let all = obj[method];
  let args = {};
  let keys = [];
  let next = [];
  let out;

  Object.keys(all).every((patt) => {
    let reg = ptx(patt, keys);
    let res = reg.exec(url);
    if(!res) return false;
    keys.forEach(({name}, i) => {
      args[name] = res[i+1];
    });
    next.push(all[patt]);
    next.push(args);
    return true;
  });

  return next;
}

/**
 * GET handler
 *
 * @param {String} patt
 * @param {Function} fn
 * @api public
 */

export function get(patt, fn) {
  add('GET', patt, fn);
}

/**
 * POST handler
 *
 * @param {String} patt
 * @param {Function} fn
 * @api public
 */

export function post(patt, fn) {
  add('POST', patt, fn);
}

/**
 * Expose
 */

export default {
  fetch,
  post,
  get,
};

