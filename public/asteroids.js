// Hyperspace is Asteroids in Javascript
// Authors: @kenpratt and @icco
//
// This code is dependent on the coquette library:
// http://coquette.maryrosecook.com/

// The main game initializer. This function sets up the game.
var Hyperspace = function() {
  this.c = new Coquette(this, "canvas", 1000, 600, "#000");

  this.conn = new ServerConnection("ws://" + window.location.host + "/ws");
  this.conn.connect();

  this.conn.handle("position", function(pos) {
    var ship = this.c.entities.all()[0]
    ship.center = pos;
  }.bind(this));

  // Create the ship that the current player drives. It differs from all other
  // ships in that it has an update loop (called every tick) that takes in
  // directions from the keyboard.
  this.c.entities.create(Ship, { center: { x:256, y:110 }, color:"#f07",
    // Movement is based off of this SO article which basically reminded me how
    // vectors work: http://stackoverflow.com/a/3639025/1063
    update: function() {
      var angleToVector = function(angle) {
        // Convert to radians.
        var r = angle * 0.01745;
        return unitVector({ x: Math.sin(r), y: -Math.cos(r) });
      };
      var magnitude = function(vector) {
        return Math.sqrt(vector.x * vector.x + vector.y * vector.y);
      };
      var unitVector = function(vector) {
        return {
          x: vector.x / magnitude(vector),
          y: vector.y / magnitude(vector)
        };
      };

      this.angle %= 360;

      // Back and forth movement
      // TODO(icco): Support W and ,
      if (this.c.inputter.isDown(this.c.inputter.UP_ARROW)) {
        var vector = angleToVector(this.angle);
        this.center.x += vector.x;
        this.center.y += vector.y;

      // TODO(icco): Support S and O
      } else if (this.c.inputter.isDown(this.c.inputter.DOWN_ARROW)) {
        // TODO(icco): Support breaking.
      }

      // Turning.
      // TODO(icco): Support D and E
      if (this.c.inputter.isDown(this.c.inputter.RIGHT_ARROW)) {
        this.angle += 0.6;

      // TODO(icco): Support A
      } else if (this.c.inputter.isDown(this.c.inputter.LEFT_ARROW)) {
        this.angle -= 0.6;
      }

      if (this.lastX !== this.center.x || this.lastY !== this.center.y) {
        this.lastX = this.center.x;
        this.lastY = this.center.y;
        this.conn.send("position", this.center);
      }
    },
  });
};

// This defines the basic ship shape as a series of verices for a path to
// follow.
var ship_shape = [
  [ 0, -5],
  [-5,  5],
  [ 0,  2],
  [ 5,  5],
  [ 0, -5]
]

// The actual ship entity. One of these will be created for every single player
// in the game. Please set the color.
var Ship = function(game, settings) {
  this.c = game.c;
  this.conn = game.conn;
  for (var i in settings) {
    this[i] = settings[i];
  }

  // This is the size of the ship.
  this.scale = 1.5;
  this.size = { x: 10 * this.scale, y: 10 * this.scale }
  this.angle = 0;

  // This is run every tick to draw the ship.
  this.draw = function(ctx) {
    // The color of the outline of the ship.
    ctx.strokeStyle = settings.color;

    // Draw the actual ship body.
    ctx.beginPath();
    for (i in ship_shape) {
      var vertex = ship_shape[i];
      var x = (vertex[0] * this.scale) + this.center.x;
      var y = (vertex[1] * this.scale) + this.center.y;
      ctx.lineTo(x, y);
    }
    ctx.stroke();
  };
};

var ServerConnection = function(url) {
  this.url = url;
  this.socket = null;
  this.handlers = {};
};

ServerConnection.prototype.connect = function() {
  this.socket = new WebSocket(this.url);
  this.socket.onopen = this.onConnect.bind(this);
  this.socket.onclose = this.onDisconnect.bind(this);
  this.socket.onmessage = this.onMessage.bind(this);
}

ServerConnection.prototype.onConnect = function() {
  console.log("websocket connected");
};

ServerConnection.prototype.onDisconnect = function() {
  console.log("websocket disconnected");
};

ServerConnection.prototype.onMessage = function(ev) {
  var raw = JSON.parse(ev.data);
  var type = raw[0];
  var data = raw[1];
  console.log("websocket received message", type, data);

  var fn = this.handlers[type];
  if (fn) {
    fn(data);
  }
};

ServerConnection.prototype.send = function(type, data) {
  var msg = JSON.stringify([type, data]);
  if (this.socket.readyState === this.socket.OPEN) {
    console.log("websocket sending message", type, data);
    this.socket.send(msg);
  } else {
    console.warn("websocket not connected, can't send message", type, data);
  }
};

ServerConnection.prototype.handle = function(type, fn) {
  this.handlers[type] = fn;
};

window.addEventListener('load', function() {
  // Begin the game once the page is loaded! Party like it's 1979!
  new Hyperspace();
});
