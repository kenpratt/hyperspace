// Hyperspace is Asteroids in Javascript
// Authors: @kenpratt and @icco

// The main game initializer. This function sets up the game.
var Hyperspace = function() {
  this.size = {x: 1000, y: 600};
  this.c = new Coquette(this, "canvas", this.size.x, this.size.y, "#000");

  this.conn = new ServerConnection("ws://" + window.location.host + "/ws");
  this.conn.connect();

  this.playerId = null;
  this.constants = null;
  this.ships = {};
  this.projectiles = {};
  this.asteroids = {};

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
    this.constants = data.constants;
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
      // console.log("Updating ship", data);
      this.ships[id].center = data.position;
      this.ships[id].angle = data.angle;
    } else {
      if (data.id === this.playerId) {
        // console.log("Adding own ship");
        this.addOwnShip(data);
      } else {
        // console.log("Adding enemy ship");
        this.addEnemyShip(data);
      }
    }
  }

  // add/update projectiles
  for (id in state.projectiles) {
    var data = state.projectiles[id];
    if (this.projectiles[id]) {
      // console.log("Updating projectile", data);
      this.projectiles[data.id].center = data.position;
      this.projectiles[data.id].angle = data.angle;
    } else {
      // console.log("Adding projectile", data);
      this.addProjectile(data);
    }
  }

  // This actually does work. Deletes all projectiles once the server deletes them.
  var ents = this.c.entities.all(Projectile);
  for (var i in ents) {
    ent = ents[i];
    if (ent != undefined && state.projectiles[ent.id] == undefined) {
      this.c.entities.destroy(ent);
    }
  }

  // add/update asteroids
  for (id in state.asteroids) {
    var data = state.asteroids[id];
    if (this.asteroids[id]) {
      // console.log("Updating asteroid", data);
      this.asteroids[data.id].center = data.position;
      this.asteroids[data.id].angle = data.angle;
      this.asteroids[data.id].velocity = data.velocity;
    } else {
      // console.log("Adding asteroid");
      this.addAsteroid(data);
    }
  }
};
