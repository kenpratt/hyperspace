// Hyperspace is Asteroids in Javascript
// Authors: @kenpratt and @icco

// The main game initializer. This function sets up the game.
var Hyperspace = function(params) {
  this.params = params;

  // selectively enable/disable which updates are processed
  this.clientUpdatesEnabled = !params.updates || params.updates == "client" || params.updates == "both";
  this.serverUpdatesEnabled = !params.updates || params.updates == "server" || params.updates == "both";

  this.size = {x: 1000, y: 600};
  this.c = new Coquette(this, "canvas", this.size.x, this.size.y, "#000");

  this.conn = new ServerConnection("ws://" + window.location.host + "/ws", { latency: this.params.latency });
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

  this.conn.handle("init", function(data, t) {
    this.playerId = data.playerId;
    this.constants = data.constants;
    this.handleUpdate(data.state, t);
  }.bind(this));

  this.conn.handle("update", function(data, t) {
    if (this.serverUpdatesEnabled) {
      this.handleUpdate(data, t);
    }
  }.bind(this));
};

Hyperspace.prototype.handleUpdate = function(state, t) {
  var elapsed = this.conn.estimatedServerTime() - t;

  // add/update ships
  for (id in state.ships) {
    var data = state.ships[id];
    var obj = this.ships[id];
    if (obj) {
      // console.log("Updating ship", data);
      for (f in data) {
        obj[f] = data[f];
      }
      obj.center = obj.position; // update center alias for position
    } else {
      if (data.id === this.playerId) {
        // console.log("Adding own ship");
        obj = this.addOwnShip(data);
      } else {
        // console.log("Adding enemy ship");
        obj = this.addEnemyShip(data);
      }
    }

    // simulate physics since server sent this message
    obj.update(elapsed);
  }

  // add/update projectiles
  for (id in state.projectiles) {
    var data = state.projectiles[id];
    var obj = this.projectiles[id];
    if (obj) {
      // console.log("Updating projectile", data);
      for (f in data) {
        obj[f] = data[f];
      }
      obj.center = obj.position; // update center alias for position
    } else {
      // console.log("Adding projectile", data);
      obj = this.addProjectile(data);
    }

    // simulate physics since server sent this message
    obj.update(elapsed);
  }

  // This actually does work. Deletes all projectiles once the server sets alive to false.
  var ents = this.c.entities.all(Projectile);
  for (var i in ents) {
    var ent = ents[i];
    if (ent && !ent.alive) {
      // console.log("Destroying projectile", ent);
      this.c.entities.destroy(ent);
    }
  }

  // add/update asteroids
  for (id in state.asteroids) {
    var data = state.asteroids[id];
    var obj = this.asteroids[id];
    if (obj) {
      // console.log("Updating asteroid", data);
      for (f in data) {
        obj[f] = data[f];
      }
      obj.center = obj.position; // update center alias for position
    } else {
      // console.log("Adding asteroid");
      obj = this.addAsteroid(data);
    }

    // simulate physics since server sent this message
    obj.update(elapsed);
  }
};
