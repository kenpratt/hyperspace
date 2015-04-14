// Hyperspace is Asteroids in Javascript
// Authors: @kenpratt and @icco
//
// This code is dependent on the coquette library:
// http://coquette.maryrosecook.com/

// The main game initializer. This function sets up the game.
var Hyperspace = function() {
  this.c = new Coquette(this, "canvas", 1000, 600, "#000");

  this.connect();

  // Create the ship that the current player drives. It differs from all other
  // ships in that it has an update loop (called every tick) that takes in
  // directions from the keyboard.
  this.c.entities.create(Ship, { center: { x:256, y:110 }, color:"#f07",
    update: function() {
      // TODO(icco): Support W and ,
      if (this.c.inputter.isDown(this.c.inputter.UP_ARROW)) {
        this.center.y -= 0.4;
      }

      // TODO(icco): Support S and O
      if (this.c.inputter.isDown(this.c.inputter.DOWN_ARROW)) {
        this.center.y += 0.4;
      }

      // TODO(icco): Support D and E
      if (this.c.inputter.isDown(this.c.inputter.RIGHT_ARROW)) {
        this.center.x += 0.4;
      }

      // TODO(icco): Support A
      if (this.c.inputter.isDown(this.c.inputter.LEFT_ARROW)) {
        this.center.x -= 0.4;
      }
    },
  });
};

Hyperspace.prototype.connect = function() {
  this.socket = new WebSocket("ws://" + window.location.host + "/ws");

  var that = this;
  this.socket.onopen = function() {
    console.log("websocket connected");
    that.socket.send("test");
  };

  this.socket.onclose = function() {
    console.log("websocket disconnected");
  };

  this.socket.onmessage = function(msg) {
    console.log("websocket received message", msg);
  };
}

// This defines the basic ship shape as a series of verices for a path to
// follow.
var ship_shape = [
  [0, 0],
  [-5, 10],
  [0, 7],
  [5, 10],
  [0, 0]
]

// The actual ship entity. One of these will be created for every single player
// in the game. Please set the color.
var Ship = function(game, settings) {
  this.c = game.c;
  for (var i in settings) {
    this[i] = settings[i];
  }

  // This is the size of the ship.
  this.scale = 1.5;

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

window.addEventListener('load', function() {
  // Begin the game once the page is loaded! Party like it's 1979!
  new Hyperspace();
});
