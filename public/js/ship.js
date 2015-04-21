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
    update: function(elapsedMillis) {
      var elapsed = this.game.clientUpdatesEnabled ? elapsedMillis / 1000 : 0;

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
        this.center.x += vector.x * this.game.constants.ship_acceleration * elapsed;
        this.center.y += vector.y * this.game.constants.ship_acceleration * elapsed;
      } else if (this.pressed["down"]) {
        // TODO(icco): Support breaking.
      }

      // Turning.
      if (this.pressed["right"]) {
        this.angle += this.game.constants.ship_rotation * elapsed;
      } else if (this.pressed["left"]) {
        this.angle -= this.game.constants.ship_rotation * elapsed;
      }

      // Send server events for key press changes.
      if (last_pressed['forward'] !== this.pressed['forward']) {
        this.conn.send("changeAcceleration", { direction: this.pressed['forward'] ? 1 : 0 });
      }

      if (last_pressed['left'] !== this.pressed['left'] || last_pressed['right'] !== this.pressed['right']) {
        var direction = (this.pressed['left'] ? -1 : (this.pressed['right'] ? 1 : 0));
        this.conn.send("changeRotation", { direction: direction });
      }

      // Fire the lasers! Say Pew Pew Pew every time you press the space bar
      // please.
      if (this.c.inputter.isPressed(this.c.inputter.SPACE)) {
        var projectileId = this.id + "." + Date.now();
        this.game.addProjectile({
          id: projectileId,
          position: { x:this.center.x, y:this.center.y },
          velocity: utils.angleAndSpeedToVector(this.angle, this.game.constants.projectile_speed),
          angle: this.angle,
          owner: this.id,
          sendEvent: true,
        });
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
  this.angle = 0;

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
