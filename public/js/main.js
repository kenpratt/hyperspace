// adapted from http://stackoverflow.com/questions/979975/how-to-get-the-value-from-the-url-parameter
function getQueryParams(qs) {
  qs = qs.replace(/\+/g, " ")

  var params = {};
  var tokens = null;
  var re = /[?&]?([^=]+)=([^&]*)/g;

  while (tokens = re.exec(qs)) {
    var k = decodeURIComponent(tokens[1]);
    var v = decodeURIComponent(tokens[2]);
    params[k] = v;
  }

  return params;
}

window.addEventListener('load', function() {
  // Begin the game once the page is loaded! Party like it's 1979!
  var params = getQueryParams(document.location.search);
  new Hyperspace(params);
});
