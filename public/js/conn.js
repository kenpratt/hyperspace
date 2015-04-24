var ServerConnection = function(url, params) {
  this.url = url;
  this.params = params;
  this.socket = null;
  this.handlers = {};
  this.nextHeartbeat = null;
  this.heartbeatSentAt = null;
  this.clockDiffSMA = utils.simpleMovingAverage(100, 0.1);
  this.latency = 0;
  this.clockDiff = 0;
};

ServerConnection.prototype.connect = function() {
  this.socket = new WebSocket(this.url);
  this.socket.binaryType = "arraybuffer";
  this.socket.onopen = this.onConnect.bind(this);
  this.socket.onclose = this.onDisconnect.bind(this);
  this.socket.onmessage = this.onMessage.bind(this);
}

ServerConnection.prototype.onConnect = function() {
  console.log("websocket connected");

  // schedule next heartbeat
  this.nextHeartbeat = setTimeout(this.sendHeartbeat.bind(this), 100);
};

ServerConnection.prototype.onDisconnect = function() {
  console.log("websocket disconnected");
  clearTimeout(this.nextHeartbeat);
};

ServerConnection.prototype.onMessage = function(ev) {
  // var msg = JSON.parse(ev.data);
  var msg = bson.BSON.deserialize(new Uint8Array(ev.data));
  if (this.params.latency) {
    setTimeout(function() {
      this.handleMessage(msg);
    }.bind(this), this.params.latency / 2);
  } else {
    this.handleMessage(msg);
  }
};

ServerConnection.prototype.handleMessage = function(msg) {
  var type = msg.type;
  var data = msg;
  var time = msg.time;

  switch (type) {
  case "h":
    this.onHeartbeat(data, time);
    break;
  case "update":
    this.onUpdate(data, time);
    // *don't* break -- fall through to default
  default:
    var fn = this.handlers[type];
    if (fn) {
      fn(data, time);
    }
  }
};

ServerConnection.prototype.send = function(type, data) {
  return;
  var msg = JSON.stringify({ type: type, time: this.now(), data: data });
  if (this.socket.readyState === this.socket.OPEN) {
    // console.log("websocket sending message", type, data);
    if (this.params.latency) {
      setTimeout(function() { this.socket.send(msg); }.bind(this), this.params.latency / 2);
    } else {
      this.socket.send(msg);
    }
  } else {
    console.warn("websocket not connected, can't send message", type, data);
  }
};

ServerConnection.prototype.handle = function(type, fn) {
  this.handlers[type] = fn;
};

ServerConnection.prototype.sendHeartbeat = function() {
  this.heartbeatSentAt = Date.now();
  this.send("h", null);
};

ServerConnection.prototype.onHeartbeat = function(data, serverTime) {
  var now = Date.now();
  var elapsed = now - this.heartbeatSentAt;

  // update estimated client/server clock difference
  this.clockDiff = Math.round(this.clockDiffSMA(now - Math.round(elapsed/2) - serverTime));
};

ServerConnection.prototype.onUpdate = function(data, serverTime) {
  var now = Date.now();
  var elapsed = now - this.heartbeatSentAt;

  // schedule next heartbeat for (100ms - time since last heartbeat sent)
  this.nextHeartbeat = setTimeout(this.sendHeartbeat.bind(this), Math.max(100 - elapsed, 0));
};

ServerConnection.prototype.now = function() {
  return Date.now() + this.clockDiff;
};
