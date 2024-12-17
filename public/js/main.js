document.addEventListener("DOMContentLoaded", function() {
    hljs.highlightAll();

    const flashMessages = document.querySelectorAll(".flash");
    const FLASH_DELAY = 3000;
    flashMessages.forEach((flash, i) => {

        setTimeout(() => {
            flash.classList.add("hide");
        }, FLASH_DELAY + i * 1000);
    });

    const codeBlocks = document.querySelectorAll("pre");
    if (codeBlocks.length) {
        codeBlocks.forEach(block => {
            const copyIconHTML = '<i class="fa-regular fa-copy"></i>';
            const copyIconSuccessHTML = '<i class="fa-solid fa-check"></i>'
            const copyIconWrapper = document.createElement('span');
            const blockText = block.querySelector('code')?.innerText || '';

            copyIconWrapper.classList.add('copy-code-btn');
            copyIconWrapper.dataset.tooltip = 'Copy';
            copyIconWrapper.innerHTML = copyIconHTML;
            copyIconWrapper.addEventListener('click', () => {
                copyIconWrapper.innerHTML = copyIconSuccessHTML;

                copyToClipboard(blockText)

                setTimeout(() => {
                    copyIconWrapper.innerHTML = copyIconHTML;
                }, 1000)
            })

            block.appendChild(copyIconWrapper);
        })
    }

    /**
     * Copies the provided text to the clipboard
     * @param {string} textToCopy
     * @returns void
     */
    async function copyToClipboard(textToCopy) {
        try {
            await navigator.clipboard.writeText(textToCopy);
        } catch (e) {
            console.log('Failed to copy!', e);
        }
    }

    const headings = document.querySelectorAll("h1, h2, h3, h4")
    document.addEventListener('scroll', () => {
        const contentLinks = document.querySelectorAll("#contents-menu li a")
        if (headings.length && contentLinks.length) {
            headings.forEach(h => {
                const rect = h.getBoundingClientRect()
                const currLink = document.querySelector(`#contents-menu a[href="#${h.id}"]`)
                if (!currLink) {
                    return
                }

                if (rect.top <= 40) {
                    contentLinks.forEach(link => {
                        link.removeAttribute('aria-current')
                    })
                    currLink.ariaCurrent = true
                } else {
                    currLink.removeAttribute('aria-current')
                }
            })
        }
    })
});
