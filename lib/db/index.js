
import monk from 'monk';

/**
 * MongoDB
 */

const db = monk('localhost/qn');

/**
 * Expose `db`
 */

export default db;

