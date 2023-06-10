// Function to update the countdown timer
function updateCountdownTimer(targetTime, elementId) {
    var countdownElement = document.getElementById(elementId);

    // Update the timer every second
    var countdownInterval = setInterval(function () {
        var currentTime = new Date().getTime();
        var timeRemaining = targetTime - currentTime;

        // Check if the countdown has ended
        if (timeRemaining <= 0) {
            clearInterval(countdownInterval);
            countdownElement.textContent = "---"
            return;
        }

        // Calculate remaining time
        var days = Math.floor(timeRemaining / (1000 * 60 * 60 * 24));
        var hours = Math.floor((timeRemaining % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
        var minutes = Math.floor((timeRemaining % (1000 * 60 * 60)) / (1000 * 60));
        var seconds = Math.floor((timeRemaining % (1000 * 60)) / 1000);

        // Add zero padding to hours, minutes, and seconds if needed
        days = String(days).padStart(3, " ");
        hours = String(hours).padStart(2, "0");
        minutes = String(minutes).padStart(2, "0");
        seconds = String(seconds).padStart(2, "0");

        // Format the countdown timer
        var countdownString = days + "d " + hours + "h " + minutes + "m " + seconds + "s";

        countdownElement.textContent = countdownString;
    }, 1000);
}

// Get the session start times and update countdown timers
var sessionStartTimes = document.getElementsByClassName("start-time");
for (var i = 0; i < sessionStartTimes.length; i++) {
    var startTime = new Date(sessionStartTimes[i].textContent + "+0000").getTime();
    var elementId = "countdown-" + sessionStartTimes[i].textContent.replace(/\D/g, "");
    updateCountdownTimer(startTime, elementId);
}