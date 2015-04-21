function getRandom(min, max) {
  return Math.random() * (max - min) + min;
}

// http://stackoverflow.com/a/6444043/1063
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
  simpleMovingAverage: function(period) {
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
      length++; // update length

      // advance pointer
      idx++;
      if (idx >= period) {
        idx = 0;
      }

      // return result
      return sum / (length < period ? length : period);
    }
  },
}
