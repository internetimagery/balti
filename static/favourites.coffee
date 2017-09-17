# Check if we can actually use localStorage
if typeof(Storage) != "undefined"

  # Collect our elements
  form = document.getElementById "form"
  meal = document.getElementsByClassName "meal"
  spice = document.getElementById "spice"
  side = document.getElementById "side"
  name = document.getElementById "name"


  # Set values to favourites.
  for m in meal
    if localStorage.meal == m.value
      m.checked = true
  spice.value = localStorage.spice or spice.value
  side.value = localStorage.side or side.value
  name.value = localStorage.name or ""

  # Save favourites upon order submission
  form.onsubmit = ->
    for m in meal
      if m.checked
        localStorage.meal = m.value
    localStorage.spice = spice.value
    localStorage.side = side.value
    localStorage.name = name.value
