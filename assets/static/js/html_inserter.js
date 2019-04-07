
"use strict";

var injectHTMLIntoDOMNode = function(remotePage, element) {

    let request = new XMLHttpRequest();
    request.onload = () => {element.innerHTML = request.responseText};
    let requestLocation = window.location.protocol + "//" + window.location.host + "/" + remotePage;
    request.open("GET", requestLocation);
    request.send();
}

