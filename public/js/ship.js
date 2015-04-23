Hyperspace.prototype.addOwnShip = function(data) {
  // Create the ship that the current player drives. It differs from all other
  // ships in that it has an update loop (called every tick) that takes in
  // directions from the keyboard.
  var extra = {
    color:"#f07",
    pressed: {
      forward: false,
      down: false,
      left: false,
      right: false,
    },
    lastEventId: 0,

    // Movement is based off of this SO article which basically reminded me how
    // vectors work: http://stackoverflow.com/a/3639025/1063
    update: function(elapsedMillis) {
      // This keeps the players ship always in the center.
      this.c.renderer.setViewCenter(this.center);

      // Move ship forward and/or rotate
      this.applyPhysics(elapsedMillis);

      var last_pressed = {};
      for (i in this.pressed) {
        last_pressed[i] = this.pressed[i];
      }

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

      // Send server events for key press changes (and update local state).
      if (last_pressed['forward'] !== this.pressed['forward']) {
        var direction = this.pressed['forward'] ? 1 : 0;
        this.acceleration = direction;
        this.send("changeAcceleration", { direction: direction });
      }

      if (last_pressed['left'] !== this.pressed['left'] || last_pressed['right'] !== this.pressed['right']) {
        var direction = (this.pressed['left'] ? -1 : (this.pressed['right'] ? 1 : 0));
        this.rotation = direction;
        this.send("changeRotation", { direction: direction });
      }

      // Fire the lasers! Say Pew Pew Pew every time you press the space bar
      // please.
      if (this.c.inputter.isPressed(this.c.inputter.SPACE)) {
        var projectileId = this.id + "." + this.conn.now();

        if (this.game.clientUpdatesEnabled) {
          var projectile = this.game.addProjectile({
            id: projectileId,
            alive: true,
            center: { x:this.center.x, y:this.center.y },
            velocity: utils.angleAndSpeedToVector(this.angle, this.game.constants.projectile_speed),
            angle: this.angle,
            owner: this.id,
          });

          // Send an event (a cause of a thing) that describes what just happened.
          this.send("fire", {
            projectileId: projectile.id,
            created: projectile.created,
          });
        } else {
          // Send an event (a cause of a thing) that describes what just happened.
          this.send("fire", {
            projectileId: projectileId,
            created: this.conn.now(),
          });
        }
      }
    },
    send: function(type, data) {
      var eventId = ++this.lastEventId;
      data.eventId = eventId;
      this.conn.send(type, data);
    },
  };
  for (k in extra) { data[k] = extra[k]; }
  var ship = this.c.entities.create(Ship, data);
  this.ships[data.id] = ship;
  return ship;
};

Hyperspace.prototype.addEnemyShip = function(data) {
  var extra = {
    color:"#0f7",
    update: function(elapsedMillis) {
      this.applyPhysics(elapsedMillis);
    },
  };
  for (k in extra) { data[k] = extra[k]; }
  var ship = this.c.entities.create(Ship, data);
  this.ships[data.id] = ship;
  return ship;
};

// This defines the basic ship shape as a series of verices for a path to
// follow.
var ship_shape = [
  { x:  0, y: -5},
  { x: -5, y:  5},
  { x:  0, y:  2},
  { x:  5, y:  5},
  { x:  0, y: -5},
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

  // This is run every tick to draw the ship.
  this.draw = function(ctx) {
    // The color of the outline of the ship.
    ctx.strokeStyle = settings.color;
    ctx.fillStyle = increaseBrightness(settings.color, 10);

    // Draw the actual ship body.
    ctx.beginPath();
    for (i in ship_shape) {
      var v = ship_shape[i];
      var x = (v.x * this.scale) + this.center.x;
      var y = (v.y * this.scale) + this.center.y;
      ctx.lineTo(x, y);
    }
    ctx.stroke();
    ctx.fill();
  };
};

Ship.prototype.applyPhysics = function(elapsedMillis) {
  if (!this.game.clientUpdatesEnabled) { return; }

  var elapsed = elapsedMillis / 1000;

  // Apply rotation
  if (this.rotation != 0) {
    this.angle += this.game.constants.ship_rotation * elapsed * this.rotation;
    while (this.angle < 0) { this.angle += 360 }
    while (this.angle >= 360) { this.angle -= 360 }
    this.angle = utils.roundToPlaces(this.angle, 1);
  }

  // Apply acceleration
  if (this.acceleration != 0) {
    var vector = utils.angleToVector(this.angle);

    this.center.x = utils.roundToPlaces(this.center.x + vector.x * this.game.constants.ship_acceleration * elapsed * this.acceleration, 1);
    this.center.y = utils.roundToPlaces(this.center.y + vector.y * this.game.constants.ship_acceleration * elapsed * this.acceleration, 1);

  }
}
