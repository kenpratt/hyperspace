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
