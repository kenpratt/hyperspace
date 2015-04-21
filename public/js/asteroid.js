Hyperspace.prototype.addAsteroid = function(data) {
  this.asteroids[data.id] = this.c.entities.create(Asteroid, data);
};

var Asteroid = function(game, settings) {
  this.game = game;
  this.c = game.c;
  this.conn = game.conn;
  for (var i in settings) {
    this[i] = settings[i];
  }

  this.boundingBox = this.c.collider.CIRCLE;
  this.size = { x: 10, y: 10 };
  this.zindex = -1;
  this.center = this.position;

  this.update = function(elapsedMillis) {
    var elapsed = this.game.clientUpdatesEnabled ? elapsedMillis / 1000 : 0;

    this.center.x += this.velocity.x * elapsed;
    this.center.y += this.velocity.y * elapsed;
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
