var Locations = [
    //How Should Looks Object
    // {
    //     name: location name,
    //     lat: 0,
    //     lng: 0,
    //     date: concert date
    // },
]

function myMap() {
    //Default position London
    let center = { lat: 51.52104532350735, lng: -0.1249710913010198 }

    if (Locations.length == 0) {
        console.log("Не удалось загрузить карту")
    } else {
        center = Locations[0]
    }

    var mapOptions = {
        center: new google.maps.LatLng(center.lat, center.lng),
        zoom: 4,
        mapTypeId: google.maps.MapTypeId.HYBRID
    }

    // Загружаем Карту
    var map = new google.maps.Map(document.getElementById("map"), mapOptions);

    // Ставим метки
    Locations.forEach(function (location) {
        var marker = new google.maps.Marker({
            position: location,
            map: map,
            title: `${location.name}\n${location.date}`
        });
        console.log("Load map location: " + location)
    });
}