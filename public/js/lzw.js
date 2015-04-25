var LZW = {
  compress: true,
  max: 65535,
  encoder: function() {
    var code, dictionary, last, previous, reset, result;
    dictionary = {};
    result = "";
    code = -1;
    previous = "";
    last = null;
    reset = (function(_this) {
      return function() {
        var i, _i, _results;
        dictionary = {};
        result = "";
        code = -1;
        previous = "";
        last = null;
        _results = [];
        for (i = _i = 0; _i < 256; i = ++_i) {
          code = _this._next(code);
          _results.push(dictionary[String.fromCharCode(i)] = code);
        }
        return _results;
      };
    })(this);
    reset();
    return (function(_this) {
      return function(structure, update) {
        var character, i, phrase, raw, _i, _ref;
        if (update == null) {
          update = false;
        }
        raw = structure;
        if (!(typeof structure === "string")) {
          raw = JSON.stringify(raw);
        }
        character = null;
        phrase = "";
        if (!_this.compress) {
          return raw;
        }
        if (update && last) {
          if (raw.indexOf(last) === 0) {
            raw = raw.substring(last.length);
          } else {
            reset();
          }
        }
        last = raw;
        for (i = _i = 0, _ref = raw.length; 0 <= _ref ? _i < _ref : _i > _ref; i = 0 <= _ref ? ++_i : --_i) {
          character = raw.charAt(i);
          phrase = previous + character;
          if (dictionary.hasOwnProperty(phrase)) {
            previous = phrase;
          } else {
            result += String.fromCharCode(dictionary[previous]);
            previous = String(character);
            if ((code = _this._next(code)) && (code < _this.max)) {
              dictionary[phrase] = code;
            }
          }
        }
        if (previous !== "") {
          return result + String.fromCharCode(dictionary[previous]);
        } else {
          return result;
        }
      };
    })(this);
  },
  encode: function(structure) {
    return this.encoder()(structure);
  },
  decode: function(str, parse) {
    var code, dictionary, entry, i, key, previous, raw, result, _i, _j, _ref;
    if (parse == null) {
      parse = false;
    }
    raw = str.split("");
    dictionary = {};
    result = null;
    code = -1;
    previous = raw[0].charAt(0);
    result = previous;
    entry = "";
    key = null;
    if (!this.compress) {
      return raw;
    }
    for (i = _i = 0; _i < 256; i = ++_i) {
      code = this._next(code);
      dictionary[code] = String.fromCharCode(i);
    }
    code = this._next(code);
    for (i = _j = 1, _ref = raw.length; 1 <= _ref ? _j < _ref : _j > _ref; i = 1 <= _ref ? ++_j : --_j) {
      key = raw[i].charCodeAt(0);
      if (dictionary.hasOwnProperty(key)) {
        entry = dictionary[key];
      } else {
        if (key !== code) {
          throw new Error("invalid sequence, dictionary corrupt, expected " + code + " was " + key);
        }
        if (code >= this.max) {
          throw new Error("dictionary size exceeded");
        }
        entry = previous + previous.charAt(0);
      }
      result = result + entry;
      if (code !== (this.max + 1)) {
        dictionary[code] = previous + entry.charAt(0);
        code = this._next(code);
      }
      previous = entry;
    }
    if (parse) {
      return JSON.parse(result);
    } else {
      return result;
    }
  },
  _next: function(code) {
    code++;
    if (code >= 0xD800 && code <= 0xDFFF) {
      code = 0xDFFF + 1;
    }
    if (code === 0xFFFE) {
      code++;
    }
    if (code === 0xFFFF) {
      code++;
    }
    return code;
  }
};
