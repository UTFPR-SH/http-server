// Adds burger functionallity

document.addEventListener('DOMContentLoaded', function () {

  // Get all "navbar-burger" elements
  var $navbar_burgers = Array.prototype.slice.call(document.querySelectorAll('.navbar-burger'), 0)

  // Check if there are any navbar burgers
  if ($navbar_burgers.length > 0) {

    // Add a click event on each of them
    $navbar_burgers.forEach(function ($el) {
      $el.addEventListener('click', function () {

        // Get the target from the "data-target" attribute
        var target = $el.dataset.target;
        var $target = document.getElementById(target)

        // Toggle the class on both the "navbar-burger" and the "navbar-menu"
        $el.classList.toggle('is-active');
        $target.classList.toggle('is-active')

      })
    })
  }
})
