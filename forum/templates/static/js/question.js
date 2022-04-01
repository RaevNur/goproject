// import "/static/js/showdown.js";
let mdConverter = new showdown.Converter();


Array.from(document.getElementsByClassName("markdown")).forEach((element) => {
    element.innerHTML = mdConverter.makeHtml(element.innerHTML)
});