// Hyperspace is Asteroids in Javascript
// Authors: @kenpratt and @icco
//
// This code is dependent on the coquette library:
// http://coquette.maryrosecook.com/

function getRandom(min, max) {
  return Math.random() * (max - min) + min;
}

// http://stackoverflow.com/a/6444043/1063
function increaseBrightness(hex, percent) {
  // strip the leading # if it's there
  hex = hex.replace(/^\s*#|\s*$/g, '');

  // convert 3 char codes --> 6, e.g. `E0F` --> `EE00FF`
  if (hex.length == 3) {
    hex = hex.replace(/(.)/g, '$1$1');
  }

  var r = parseInt(hex.substr(0, 2), 16),
      g = parseInt(hex.substr(2, 2), 16),
      b = parseInt(hex.substr(4, 2), 16);

  return '#' +
    ((0|(1<<8) + r + (256 - r) * percent / 100).toString(16)).substr(1) +
    ((0|(1<<8) + g + (256 - g) * percent / 100).toString(16)).substr(1) +
    ((0|(1<<8) + b + (256 - b) * percent / 100).toString(16)).substr(1);
}

var utils = {
  angleToVector: function(angle) {
    // Convert to radians.
    var r = angle * 0.01745;
    return utils.unitVector({ x: Math.sin(r), y: -Math.cos(r) });
  },
  magnitude: function(vector) {
    return Math.sqrt(vector.x * vector.x + vector.y * vector.y);
  },
  unitVector: function(vector) {
    return {
      x: vector.x / utils.magnitude(vector),
      y: vector.y / utils.magnitude(vector)
    };
  },
}

// The main game initializer. This function sets up the game.
var Hyperspace = function() {
  this.size = {x: 1000, y: 600};
  this.c = new Coquette(this, "canvas", this.size.x, this.size.y, "#000");

  this.conn = new ServerConnection("ws://" + window.location.host + "/ws");
  this.conn.connect();

  this.playerId = null;
  this.ships = {};
  this.projectiles = {};

  this.update = function() {
    center = this.c.renderer.getViewCenter()
    for (var i = this.c.entities.all(Star).length; i < 100; i++) {
      var where = {
        x: getRandom((center.x - this.size.x/2), (center.x + this.size.x/2)),
        y: getRandom((center.y - this.size.y/2), (center.y + this.size.y/2)),
      }
      this.c.entities.create(Star, {center: where});
    }

    /* For Debugging
    for (var i in this.c.entities.all(Ship)) {
      var ship = this.c.entities.all(Ship)[i];
      console.log(ship.id, ship.center);
    }
    */
  };

  this.conn.handle("init", function(data) {
    this.playerId = data.playerId;
    this.handleUpdate(data.state);
  }.bind(this));

  this.conn.handle("update", function(data) {
    this.handleUpdate(data);
  }.bind(this));
};

Hyperspace.prototype.handleUpdate = function(state) {
  // add/update ships
  for (id in state.ships) {
    var data = state.ships[id];
    if (this.ships[id]) {
      console.log("Updating ship");
      this.ships[id].center = data.position;
    } else {
      if (data.id === this.playerId) {
        console.log("Adding own ship");
        this.addOwnShip(data);
      } else {
        console.log("Adding enemy ship");
        this.addEnemyShip(data);
      }
    }
  }

  // add/update projectiles
  for (id in state.projectiles) {
    var data = state.projectiles[id];
    if (this.projectiles[id]) {
      console.log("Updating projectile");
      this.projectiles[data.id].center = data.position;
    } else {
      console.log("Adding projectile");
      this.addProjectile(data);
    }
  }
};

var Star = function(game, settings) {
  this.c = game.c;
  for (var i in settings) {
    this[i] = settings[i];
  }
  this.zindex = -3;
  this.width = 3 + (Math.random() * 4);
  this.size = {x: this.width, y: this.width};
  this.brightness = Math.random() * 100 - 50;
  this.update = function() {
    if (!this.c.renderer.onScreen(this)) {
      this.c.entities.destroy(this);
    }
  };
  this.draw = function(ctx) {
    ctx.fillStyle = increaseBrightness("#cc9933", this.brightness);
    ctx.fillRect(this.center.x, this.center.y, this.size.x, this.size.y);
  };
};

Hyperspace.prototype.addOwnShip = function(data) {
  // Create the ship that the current player drives. It differs from all other
  // ships in that it has an update loop (called every tick) that takes in
  // directions from the keyboard.
  var ship = this.c.entities.create(Ship, {
    center: data.position,
    id: data.id,
    color:"#f07",
    pressed: {
      forward: false,
      down: false,
      left: false,
      right: false,
    },

    // Movement is based off of this SO article which basically reminded me how
    // vectors work: http://stackoverflow.com/a/3639025/1063
    update: function() {

      // This keeps the players ship always in the center.
      this.c.renderer.setViewCenter(this.center);

      var last_pressed = {};
      for (i in this.pressed) {
        last_pressed[i] = this.pressed[i];
      }

      // The range of Angle is 0 - 360.
      this.angle %= 360;

      // Key pressed booleans
      this.pressed["forward"] =
        this.c.inputter.isDown(this.c.inputter.UP_ARROW) ||
        this.c.inputter.isDown(this.c.inputter.W) ||
        this.c.inputter.isDown(this.c.inputter.COMMA);
      this.pressed["down"] =
        this.c.inputter.isDown(this.c.inputter.DOWN_ARROW) ||
        this.c.inputter.isDown(this.c.inputter.S) ||
        this.c.inputter.isDown(this.c.inputter.O);
      this.pressed["left"] =
        this.c.inputter.isDown(this.c.inputter.LEFT_ARROW) ||
        this.c.inputter.isDown(this.c.inputter.A);
      this.pressed["right"] =
        this.c.inputter.isDown(this.c.inputter.RIGHT_ARROW) ||
        this.c.inputter.isDown(this.c.inputter.D) ||
        this.c.inputter.isDown(this.c.inputter.E);

      // Back and forth movement
      if (this.pressed["forward"]) {
        var vector = utils.angleToVector(this.angle);
        this.center.x += vector.x;
        this.center.y += vector.y;
      } else if (this.pressed["down"]) {
        // TODO(icco): Support breaking.
      }

      // Turning.
      if (this.pressed["right"]) {
        this.angle += 2;
      } else if (this.pressed["left"]) {
        this.angle -= 2;
      }

      // Send server events for key press changes.
      if (JSON.stringify(last_pressed) != JSON.stringify(this.pressed)) {
        this.conn.send("button", {
          id: this.id,
          time: Date.now(),
          forward: this.pressed['forward'],
          down: this.pressed['down'],
          left: this.pressed['left'],
          right: this.pressed['right'],
        });
      }

      // Fire the lasers! Say Pew Pew Pew every time you press the space bar
      // please.
      if (this.c.inputter.isPressed(this.c.inputter.SPACE)) {
        if (this.c.entities.all(Laser).length < 30) {
          // Send an event (a cause of a thing) that describes what just
          // happened.
          var projectileId = this.id + "." + Date.now();
          this.conn.send("fire", {
            projectileId: projectileId,
          });

          this.game.addProjectile({
            id: projectileId,
            position: { x:this.center.x, y:this.center.y },
            angle: this.angle,
          });
        }
      }
    },
  });
  this.ships[data.id] = ship;
};

Hyperspace.prototype.addEnemyShip = function(data) {
  var ship = this.c.entities.create(Ship, {
    id: data.id,
    center: data.position,
    color:"#0f7"
  });
  this.ships[data.id] = ship;
};

Hyperspace.prototype.addProjectile = function(data) {
  var projectile = this.c.entities.create(Laser, {
    id: data.id,
    center: data.position,
    vector: utils.angleToVector(data.angle)
  });
  this.projectiles[data.id] = projectile;
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
  this.game = game;
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
    ctx.fillStyle = increaseBrightness(settings.color, 10);

    // Draw the actual ship body.
    ctx.beginPath();
    for (i in ship_shape) {
      var vertex = ship_shape[i];
      var x = (vertex[0] * this.scale) + this.center.x;
      var y = (vertex[1] * this.scale) + this.center.y;
      ctx.lineTo(x, y);
    }
    ctx.stroke();
    ctx.fill();
  };
};

var Laser = function(game, settings) {
  this.c = game.c;
  this.conn = game.conn;
  this.boundingBox = this.c.collider.CIRCLE;
  this.size = { x: 3, y: 3 };
  this.zindex = -1;

  for (var i in settings) {
    this[i] = settings[i];
  }

  this.update = function() {
    // TODO move age logic to server
    var age = 0; // Date.now() - this.created;
    // Kill lazers older than three seconds.
    if (age < 3000) {
      this.center.x += (this.vector.x * 2);
      this.center.y += (this.vector.y * 2);
    } else {
      this.c.entities.destroy(this);
    }
  };

  this.draw = function(ctx) {
    ctx.fillStyle = "#fff";
    ctx.beginPath();
    ctx.arc(
        this.center.x, // x
        this.center.y, // y
        this.size.x, // Radius
        0, // Start Angle
        Math.PI*2, // End Angle
        true); // Anticlockwise?
    ctx.fill();
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
  var type = raw.type;
  var data = raw.data;
  // console.log("websocket received message", type, data);

  var fn = this.handlers[type];
  if (fn) {
    fn(data);
  }
};

ServerConnection.prototype.send = function(type, data) {
  var msg = JSON.stringify({ type: type, time: Date.now(), data: data });
  if (this.socket.readyState === this.socket.OPEN) {
    // console.log("websocket sending message", type, data);
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
