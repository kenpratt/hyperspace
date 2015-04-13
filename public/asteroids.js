// Hyperspace is Asteroids in Javascript
//
// Current code is based off of the examples from
// http://coquette.maryrosecook.com/
var Hyperspace = function() {
  this.c = new Coquette(this, "canvas", 1000, 600, "#000");

  this.c.entities.create(Person, { center: { x:250, y:40 }, color:"#099" });

  this.c.entities.create(Person, { center: { x:256, y:110 }, color:"#f07",
    update: function() {
      if (this.c.inputter.isDown(this.c.inputter.UP_ARROW)) {
        this.center.y -= 0.4;
      }
    },
    collision: function(other) {
      other.center.y = this.center.y; // follow the player
    }
  });
};

var Person = function(game, settings) {
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
  new Hyperspace();
});
