// Lots of random math equations for figuring things out.


// Given a hex (#abcdef) brighten or darken by a percentage.
// Cribbed from http://stackoverflow.com/a/6444043/1063
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
  getRandom: function(min, max) {
    return Math.random() * (max - min) + min;
  },
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
  angleAndSpeedToVector: function(angle, speed) {
    return utils.multiplyVector(utils.angleToVector(angle), speed);
  },
  multiplyVector: function(vector, a) {
    return {
      x: vector.x * a,
      y: vector.y * a,
    };
  },
  simpleMovingAverage: function(period, toDiscard) {
    var nums = new Array(period);
    for (var i = 0; i < period; i++) {
      nums[i] = 0;
    }
    var length = 0;
    var idx = 0;
    var sum = 0;
    return function(num) {
      sum -= nums[idx]; // subtract last value
      nums[idx] = num; // record new value
      sum += num; // add new value
      if (length < period) { length++ }; // update length

      // advance pointer
      idx++;
      if (idx >= period) {
        idx = 0;
      }

      if (toDiscard > 0) {
        // discard the lowest and highest results in the set
        var sorted = nums.slice(0, length).sort();
        var n = Math.floor(toDiscard * length);
        var msum = 0;
        for (var i = n; i < (length - n); i++) {
          msum += sorted[i];
        }
        return msum / (length - n - n);
      } else {
        // return simple mean as result
        return sum / length;
      }
    }
  },
  roundToPlaces: function(f, places) {
    var shift = Math.pow(10, places);
    return Math.round(f*shift) / shift;
  },
}

var collision = {
  circles: function(x1, y1, r1, x2, y2, r2) {
    var dx = x2 - x1;
    var dy = y2 - y1;
    var distance = Math.sqrt(dx*dx + dy*dy);
    return (distance < r1 + r2);
  },

  squares: function(x2, y2, w2, h2, x2, y2, w2, h2) {
    var x_intersect = Math.abs(x1 - x2) * 2 < (w1 + w2);
    var y_intersect = Math.abs(y1 - y2) * 2 < (h1 + h2);
    return x_intersect && y_intersect;
  }
}
