var ServerConnection = function(url, params) {
  this.url = url;
  this.params = params;
  this.socket = null;
  this.handlers = { h: this.onHeartbeat.bind(this) };
  this.nextHeartbeat = null;
  this.heartbeatSentAt = null;
  this.latencySMA = utils.simpleMovingAverage(10);
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
  var raw = JSON.parse(ev.data);
  var type = raw.type;
  var data = raw.data;
  var time = raw.time;
  // console.log("websocket received message", type, data, time);

  var fn = this.handlers[type];
  if (fn) {
    if (this.params.latency) {
      setTimeout(function() { fn(data, time); }, this.params.latency / 2);
    } else {
      fn(data, time);
    }
  }
};

ServerConnection.prototype.send = function(type, data) {
  var msg = JSON.stringify({ type: type, time: Date.now(), data: data });
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
  this.heartbeatSentAt = new Date();
  this.send("h", null);
};

ServerConnection.prototype.onHeartbeat = function(data, serverTime) {
  var now = Date.now();
  var elapsed = now - this.heartbeatSentAt;

  // update latency & estimated client/server clock difference
  this.clockDiff = now - Math.round(elapsed/2) - serverTime;
  this.latency = Math.round(this.latencySMA(elapsed));

  // schedule next heartbeat
  this.nextHeartbeat = setTimeout(this.sendHeartbeat.bind(this), 100);
};

ServerConnection.prototype.estimatedServerTime = function() {
  return Date.now() + this.clockDiff;
};
