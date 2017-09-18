# Check if we can actually use localStorage
if typeof(Storage) != "undefined"

  # Collect our elements
  form = document.getElementById "form"
  meal = document.getElementsByClassName "meal"
  spice = document.getElementById "spice"
  side = document.getElementById "side"
  dog = document.getElementById "name"

  try
    meal_data = JSON.parse localStorage.meal
  catch error
    if error instanceof SyntaxError
      meal_data = []
    else
      throw error

  # Set values to favourites.
  for m in meal
    if m.value in meal_data
      m.checked = true
  spice.value = localStorage.spice or spice.value
  side.value = localStorage.side or side.value
  dog.value = localStorage.name or ""

  # Save favourites upon order submission
  form.onsubmit = ->
    meal_data = []
    for m in meal
      if m.checked
        meal_data.push(m.value)
    localStorage.meal = JSON.stringify meal_data
    localStorage.spice = spice.value
    localStorage.side = side.value
    localStorage.name = dog.value
