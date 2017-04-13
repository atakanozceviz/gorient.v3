$(function () {
    socket = io();
    socket.on("connection", function () {
        socket.emit("connectedto", ID)
    });
});

if (window.DeviceOrientationEvent) {
    window.addEventListener('deviceorientation', function (eventData) {
        if (!!eventData.gamma) {
            socket.emit("orient", JSON.stringify({
                id: ID,
                gamma: eventData.gamma,
                beta: eventData.beta,
                alpha: eventData.alpha
            }));
        }
    }, false);
}