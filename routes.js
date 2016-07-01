
import db from './lib/db';

/**
 * Entries collection
 */

const tbl = db.get('entries');

/**
 * Get entry
 *
 * @param {Request} req
 * @param {Object} args
 * @return {Array}
 * @api public
 */

const get = async (req, {id}) => {
  let obj = await tbl.findOne({id});

  if(obj) return [200, obj];

  return [404, 'Not Found'];
};

/**
 * New entry
 *
 * @param {Request} req
 * @param {Object} args
 * @return {Array}
 * @api public
 */

const create = async (req) => {
  return [200, 'hey'];
};

/**
 * WebSocket
 *
 * @param {String} id
 * @param {Object} raw
 * @return {Object}
 * @api public
 */

const ws = async (id, raw) => {
  return {};
};

/**
 * Expose
 */

export default {
  create,
  get,
  ws,
};

