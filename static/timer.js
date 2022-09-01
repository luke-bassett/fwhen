// stackoverflow.com/questions/20618355
function startTimer(duration, display) {
    var timer = duration, d, h, m, s;
    setInterval(function () {
        d = parseInt(timer / 86400, 10) 
        h = parseInt(timer % 86400 / 3600, 10)
        m = parseInt(timer % 3600 / 60, 10);
        s = parseInt(timer % 60, 10);

        h = h < 10 ? "0" + h : h;
        m = m < 10 ? "0" + m : m;
        s = s < 10 ? "0" + s : s;

        display.textContent = d + "d " + h + "h " + m + "m " + s + "s";

        if (--timer < 0) {
            timer = duration;
        }
    }, 1000);
}

window.onload = function () {
    var duration = NsUntilR0S0 / (10 ** 9), display = document.getElementById('timer00');
    startTimer(duration, display);
    var duration = NsUntilR0S1 / (10 ** 9), display = document.getElementById('timer01');
    startTimer(duration, display);
    var duration = NsUntilR0S2 / (10 ** 9), display = document.getElementById('timer02');
    startTimer(duration, display);
    var duration = NsUntilR0S3 / (10 ** 9), display = document.getElementById('timer03');
    startTimer(duration, display);
    var duration = NsUntilR0S4 / (10 ** 9), display = document.getElementById('timer04');
    startTimer(duration, display);
};