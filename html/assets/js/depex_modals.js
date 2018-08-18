'use strict'

// Array of modals available
var M = ['modal-bolsista', 'modal-voluntario']

document.getElementById(M[0]).onclick = event => {
    document.getElementById('m' + M[0]).classList.add("is-active")
}

document.getElementById(M[1]).onclick = event => {
    document.getElementById('m' + M[1]).classList.add("is-active")
}

var dbutton = document.getElementsByClassName('delete')

for (var i = 0; i < dbutton.length; i++) {
    dbutton[i].onclick = event => {
        M.forEach(function (e) {
            document.getElementById('m' + e).classList.remove('is-active')
        })
    }
}