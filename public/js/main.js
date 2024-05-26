document.addEventListener("DOMContentLoaded", function() {
    hljs.highlightAll();

    const flashMessages = document.querySelectorAll(".flash");

    const FLASH_DELAY = 3000;
    flashMessages.forEach((flash, i) => {

        setTimeout(() => {
            flash.classList.add("hide");
        }, FLASH_DELAY + i * 1000);
    });
});
