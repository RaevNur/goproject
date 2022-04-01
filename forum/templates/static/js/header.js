// Widgets like Search

// HEADER SEARCH
var cb_HeaderSearch = document.getElementById("cb_HeaderSearch");
var b_HeaderSearch = document.getElementById("b_HeaderSearch");

// HEADER MENU
var cb_HeaderMenu = document.getElementById("cb_HeaderMenu");
var b_HeaderMenu = document.getElementById("b_HeaderMenu");


// cb_HeaderSearch_change - Shows or hide element
function cb_HeaderSearch_change(e) {
    if (e.target.checked) {
        b_HeaderSearch.classList.add("d-flex-on-md");
        cb_HeaderSearch.classList.add("btn-tg--checked");
    } else {
        b_HeaderSearch.classList.remove("d-flex-on-md");
        cb_HeaderSearch.classList.remove("btn-tg--checked");
    }
}

// cb_HeaderSearch_change - Shows or hide element
function cb_HeaderMenu_change(e) {
    if (e.target.checked) {
        b_HeaderMenu.classList.add("d-flex");
        cb_HeaderMenu.classList.add("header__menu-btn--checked");
    } else {
        b_HeaderMenu.classList.remove("d-flex");
        cb_HeaderMenu.classList.remove("header__menu-btn--checked");
    }
}

// ADDING EVENTS
cb_HeaderSearch.addEventListener("change", cb_HeaderSearch_change);
cb_HeaderMenu.addEventListener("change", cb_HeaderMenu_change);

// console.info(window.location.href);
