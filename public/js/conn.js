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
  var payload = ev.data;
  var msg = JSON.parse(LZW.decode(payload));

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
  var data = msg.data;
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
  var payload = { type: type, time: this.now(), data: data };
  var msg = LZW.encode(JSON.stringify(payload));

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

  // estimate clock difference as amout of millis that the server clock is *ahead* of the client clock
  var estimatedCurrentServerTime = serverTime + elapsed/2;
  var diff = estimatedCurrentServerTime - now;

  // take a simple moving average of the estimated difference
  this.clockDiff = Math.round(this.clockDiffSMA(diff));
};

ServerConnection.prototype.onUpdate = function(data, serverTime) {
  var now = Date.now();
  var elapsed = now - this.heartbeatSentAt;

  // schedule next heartbeat for (100ms - time since last heartbeat sent)
  this.nextHeartbeat = setTimeout(this.sendHeartbeat.bind(this), Math.max(100 - elapsed, 0));
};

ServerConnection.prototype.now = function() {
  // clockDiff = millis that server clock is ahead of client clock, so add it to current time to get estimate of server time
  return Date.now() + this.clockDiff;
};
