var dog, error, form, i, len, m, meal, meal_data, ref, side, spice,
  indexOf = [].indexOf || function(item) { for (var i = 0, l = this.length; i < l; i++) { if (i in this && this[i] === item) return i; } return -1; };

if (typeof Storage !== "undefined") {
  form = document.getElementById("form");
  meal = document.getElementsByClassName("meal");
  spice = document.getElementById("spice");
  side = document.getElementById("side");
  dog = document.getElementById("name");
  try {
    meal_data = JSON.parse(localStorage.meal);
  } catch (error1) {
    error = error1;
    if (error instanceof SyntaxError) {
      meal_data = [];
    } else {
      throw error;
    }
  }
  for (i = 0, len = meal.length; i < len; i++) {
    m = meal[i];
    if (ref = m.value, indexOf.call(meal_data, ref) >= 0) {
      m.checked = true;
    }
  }
  spice.value = localStorage.spice || spice.value;
  side.value = localStorage.side || side.value;
  dog.value = localStorage.name || "";
  form.onsubmit = function() {
    var j, len1;
    meal_data = [];
    for (j = 0, len1 = meal.length; j < len1; j++) {
      m = meal[j];
      if (m.checked) {
        meal_data.push(m.value);
      }
    }
    localStorage.meal = JSON.stringify(meal_data);
    localStorage.spice = spice.value;
    localStorage.side = side.value;
    return localStorage.name = dog.value;
  };
}
