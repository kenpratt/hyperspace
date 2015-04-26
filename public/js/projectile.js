Hyperspace.prototype.addProjectile = function(data) {
  var projectile = this.c.entities.create(Projectile, data);
  this.projectiles[data.id] = projectile;
  return projectile;
};

var Projectile = function(game, settings) {
  this.game = game;
  this.c = game.c;
  this.conn = game.conn;
  this.size = { x: 3, y: 3 };
  this.zindex = -1;
  this.created = this.conn.now();

  for (var i in settings) {
    this[i] = settings[i];
  }

  this.update = function(elapsedMillis) {
    if (this.game.clientUpdatesEnabled) {
      this.applyPhysics(elapsedMillis);
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

Projectile.prototype.applyPhysics = function(elapsedMillis) {
  var elapsed = this.game.clientUpdatesEnabled ? elapsedMillis / 1000 : 0;

  this.center.x = utils.roundToPlaces(this.center.x + this.velocity.x * elapsed, 1);
  this.center.y = utils.roundToPlaces(this.center.y + this.velocity.y * elapsed, 1);
  // All projectile deletion is done in handleUpdate function.
};
