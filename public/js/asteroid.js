Hyperspace.prototype.addAsteroid = function(data) {
  var asteroid = this.c.entities.create(Asteroid, data);
  this.asteroids[data.id] = asteroid;
  return asteroid;
};

var Asteroid = function(game, settings) {
  this.game = game;
  this.c = game.c;
  this.conn = game.conn;
  for (var i in settings) {
    this[i] = settings[i];
  }

  this.size = { x: 10, y: 10 };
  this.zindex = -1;

  this.update = function(elapsedMillis) {
    this.applyPhysics(elapsedMillis);
  };

  this.draw = function(ctx) {
    ctx.fillStyle = "rgb(119, 58, 28)";
    ctx.beginPath();
    for (i in this.shape) {
      var v = this.shape[i];
      var x = v.x + this.center.x;
      var y = v.y + this.center.y;
      ctx.lineTo(x, y);
    }
    ctx.fill();
  };
};

Asteroid.prototype.applyPhysics = function(elapsedMillis) {
  var elapsed = this.game.clientUpdatesEnabled ? elapsedMillis / 1000 : 0;

  this.center.x = utils.roundToPlaces(this.center.x + this.velocity.x * elapsed, 1);
  this.center.y = utils.roundToPlaces(this.center.y + this.velocity.y * elapsed, 1);
};
