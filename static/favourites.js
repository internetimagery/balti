(function() {
  var form, m, meal, name, side, spice, _i, _len;

  if (typeof Storage !== "undefined") {
    form = document.getElementById("form");
    meal = document.getElementsByClassName("meal");
    spice = document.getElementById("spice");
    side = document.getElementById("side");
    name = document.getElementById("name");
    for (_i = 0, _len = meal.length; _i < _len; _i++) {
      m = meal[_i];
      if (localStorage.meal === m.value) {
        m.checked = true;
      }
    }
    spice.value = localStorage.spice || spice.value;
    side.value = localStorage.side || side.value;
    name.value = localStorage.name || "";
    form.onsubmit = function() {
      var _j, _len1;
      for (_j = 0, _len1 = meal.length; _j < _len1; _j++) {
        m = meal[_j];
        if (m.checked) {
          localStorage.meal = m.value;
        }
      }
      localStorage.spice = spice.value;
      localStorage.side = side.value;
      localStorage.name = name.value;
      return false;
    };
  }

}).call(this);
