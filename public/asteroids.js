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

// The actual ship entity. One of these will be created for every single player
// in the game. Please set the color.
var Ship = function(game, settings) {
  this.c = game.c;
  for (var i in settings) {
    this[i] = settings[i];
  }

  this.size = { x:9, y:9 };
  this.draw = function(ctx) {
    ctx.fillStyle = settings.color;
    ctx.fillRect(
        this.center.x - this.size.x / 2,
        this.center.y - this.size.y / 2,
        this.size.x,
        this.size.y);
  };
};

window.addEventListener('load', function() {
  // Begin the game once the page is loaded! Party like it's 1979!
  new Hyperspace();
});
