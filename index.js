
import {Server as WebSocket} from 'ws';
import {Server} from 'http';
import r from './routes';
import app from './app';

/**
 * Server
 */

const server = Server((...args) => app(...args));

/**
 * WS Server
 */

const wss = new WebSocket({server});

/**
 * Expose `server`
 */

export default server;

/**
 * WebSocket handler
 *
 * @param {WebSocket} ws
 * @api private
 */

wss.on('connection', (ws) => {
  let id = ws.upgradeReq.url;

  if(!hub[id]) hub[id] = [];

  ws.on('message', message);
  ws.on('close', close);

  let room = hub[id];

  room.push(ws);

  /**
   * Message event
   *
   * @param {String} raw
   * @api private
   */

  async function message(raw) {
    if('PING' == raw) return;
    let obj = JSON.parse(raw);
    let out = JSON.stringify(await r.ws(id, obj));
    room.filter(x => x != ws)
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
    let x = room.indexOf(ws);
    if(-1 == x) return;
    room.splice(x, 1);
  }

});

