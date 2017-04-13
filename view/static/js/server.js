$(document).ready(function () {
    var img = document.getElementById("wheel");
    canvas = document.getElementById("canvas");
    context = canvas.getContext("2d");
    context.drawImage(img, 0, 0);
    colordiv = document.getElementById("color");
});

$(function () {
    var socket = io();
    socket.on("connection", function (id) {
        ID = id;
        var txt = '<img class="img-fluid" src="https://api.qrserver.com/v1/create-qr-code/?data=' + location.href + 'client?id=' + id + '&amp;size=200x200" alt="Scan to connect" title="QR" />';
        $("#qr").append(txt)
    });

    socket.on("client", function (data) {
        cid = data;
    });

    socket.on("DC", function (data) {
        if (typeof cid != "undefined") {
            if (data == cid) {
                $("#qr,#info").show();
            }
        }
    });

    var logo = document.getElementById("imgLogo");

    logoStyle = logo.style,
        _transform = "WebkitTransform" in logoStyle ? "WebkitTransform" :
            "MozTransform" in logoStyle ? "MozTransform" :
                "msTransform" in logoStyle ? "msTransform" : false;

    _transform && socket.on("orient", function (data) {
        $("#qr,#info").hide();
        var orient = JSON.parse(data);

        var gamma = orient.gamma;
        var beta = orient.beta;
        var alpha = orient.alpha;

        // gamma is the left-to-right tilt in degrees
        $("#gamma").text(gamma);
        // beta is the front-to-back tilt in degrees
        $("#beta").text(beta);
        // alpha is the compass direction the device is facing in degrees
        $("#alpha").text(alpha);

        //Image rotation
        logoStyle[_transform] = "rotateY(" + gamma + "deg) rotateX(" + (-beta) + "deg) rotateZ(" + (-alpha) + "deg)";
    });

    socket.on("coord", function (coord) {
        var p = JSON.parse(coord);
        var x = p.x;
        var y = p.y;
        changeColor(x, y)
    })

});

function changeColor(x, y) {
    var pxd = context.getImageData(x, y, 1, 1).data;
    var r = pxd[0];
    var g = pxd[1];
    var b = pxd[2];
    var hex = "#" + ((1 << 24) + (r << 16) + (g << 8) + b).toString(16).slice(1);
    $("#thcolor").attr("content", hex);
    colordiv.style.backgroundColor = "rgb(" + r + "," + g + "," + b + ")";
}