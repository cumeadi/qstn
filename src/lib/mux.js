
import ptx from 'path-to-regexp';

/**
 * Route obj
 */

const obj = {};

/**
 * Add route
 *
 * @param {String} patt
 * @param {Function} fn
 * @api public
 */

export function add(patt, fn) {
  obj[patt] = fn;
}

/**
 * Fetch route
 *
 * @param {String} path
 * @return {Function}
 * @api public
 */

export function fetch(path) {
  let args = {};
  let keys = [];
  let next;

  Object.keys(obj).every((patt) => {
    let reg = ptx(patt, keys);
    let res = reg.exec(path);
    if(!res) return false;
    keys.forEach(({name}, i) => {
      args[name] = res[i+1];
    });
    next = res;
    return true;
  });

  return next;
}
