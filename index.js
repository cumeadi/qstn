
import {Server} from 'ws';
import r from './routes';
import app from './app';

/**
 * WS Server
 */

const wss = new Server({
  server: app,
});

/**
 * Player 2...
 *
 * @param {WebSocket} ws
 * @api private
 */

wss.on('connection', (ws) => {
  let {url} = ws.upgradeReq;
  let pass = /\/q\/(\w+)$/.exec(url);
  let id;
  let rm;

  if(!pass) return;

  id = pass[1];

  if(!hub[id]) hub[id] = [];

  rm = hub[id];

  ws.on('message', message);
  ws.on('close', close);

  /**
   * Message event
   *
   * @param {String} raw
   * @api private
   */

  async function message(raw) {
    let obj = JSON.parse(raw);
    let out = JSON.stringify(await r.ws(id, obj));
    rm.filter(x => x != ws)
      .forEach(ws => {
        ws.send(out)
      });
  }

  /**
   * Close event
   *
   * @api private
   */

  async function close() {
    let x = rm.indexOf(ws);
    rm.splice(x, 1);
  }

});

/**
 * Expose `app`
 */

export default app;

