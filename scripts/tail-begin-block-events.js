#!/usr/bin/env node

const WebSocket = require('ws');
const util = require('util');

const ws = new WebSocket('ws://localhost:26657/websocket');

ws.on('open', function open() {
  const params = {
    query: "tm.event = 'NewBlock'"
  };
  ws.send(JSON.stringify({
    jsonrpc: "2.0",
    id: "0",
    method: "subscribe",
    params: params,
  }));
});

ws.on('message', function incoming(data) {
  try {
    var msg = JSON.parse(data.toString('utf-8'));
    msg = msg.result.data.value.result_begin_block.events;
    msg.forEach((e) => {
      console.log(util.inspect(e, { showHidden: false, depth: null, maxArrayLength: Infinity }));
    })
    // console.log(util.inspect(msg, { showHidden: false, depth: null, maxArrayLength: Infinity }));
  } catch (err) {
    // console.log(err);
  }
});

