// Avoid `console` errors in browsers that lack a console.
(function () {
    var method;
    var noop = function noop() {};
    var methods = ['assert', 'clear', 'count', 'debug', 'dir', 'dirxml', 'error', 'exception', 'group', 'groupCollapsed', 'groupEnd', 'info', 'log', 'markTimeline', 'profile', 'profileEnd', 'table', 'time', 'timeEnd', 'timeStamp', 'trace', 'warn'];
    var length = methods.length;
    var console = (window.console = window.console || {});

    while (length--) {
        method = methods[length];

        // Only stub undefined methods.
        if (!console[method]) {
            console[method] = noop;
        }
    }
})();

// Place any jQuery/helper plugins in here.

//Truncate Plugin
/*
 * A plugin to truncate the text in a given node to the specified length
 * The full text is revealed on hover
 */
(function ($) {
    "use strict";

    window.cachedIsMobile = undefined;

    function isMobile() {
        if (window.cachedIsMobile === undefined) {
            window.cachedIsMobile = $(window).width() < 480;
        }
        $(window).on("resize", function () {
            window.cachedIsMobile = undefined;
        });
        return window.cachedIsMobile;
    }

    function truncate(length) {
        // Truncate the text inside the selected node
        return $(this).each(function () {
            // For multiple node support
            var $this = $(this),
                text, newLength, shortenedText;
            text = $this.text();
            newLength = length || $this.attr("data-length") || 20;
            shortenedText = jQuery.trim(text).substring(0, newLength);
            if (shortenedText.length < text.length) {
                // Add a continuation symbol if text was clipped
                shortenedText = shortenedText + "...";
            }

            var expand = function () {
                // Show the full length text
                $this.text(text);
            };
            var compress = function () {
                // Show the shortened text
                $this.text(shortenedText);
            };
            var toggle = function () {
                // Toggle the text
                if ($this.text() === shortenedText) {
                    expand.apply(this, arguments);
                } else {
                    compress.apply(this, arguments);
                }
            };
            $this.text(shortenedText).hover(expand, compress).focus(expand).blur(compress).on("tap", toggle);
        });
    }

    function shorten(height, notOnMobile) {
        // Shorten the selected node to the given height
        return $(this).each(function () {
            // For multiple node support
            var $this = $(this);

            var compress = function () {
                // Shorten the element again
                $this.css({
                    "overflow": "hidden",
                    "height": height
                });
                $this.trigger('compressed.shorten', [height]);
            };

            var expand = function () {
                // Allow the element to expand to full height
                $this.css({
                    "overflow": "visible",
                    "height": "auto"
                });
                $this.trigger('expanded.shorten');
            };

            var toggle = function () {
                // Toggle the height
                if ($this.css("height") === "auto") {
                    compress.apply(this, arguments);
                } else {
                    expand.apply(this, arguments);
                }
            };

            var setup = function () {
                var disable = notOnMobile && isMobile();
                height = height || $this.attr("data-height") || 240;
                if (!disable) {
                    $this.on('mouseenter.shorten', expand).on('mouseleave.shorten', compress).on('focus.shorten', expand).on('blur.shorten', compress).on("tap", toggle);
                    compress();
                } else {
                    expand();
                    $this.off(".shorten");
                }
            };

            setup();
            $(window).on("resize", setup);

        });
    }

    function hoverScroll(options) {
        // Scroll to element if the user hovers for the given period
        options = $.extend({
            wait: 750,
            scrollPeriod: 380,
            offsetMargin: 20
        }, options || {});
        var $containers = $('html, body');
        return $(this).each(function () {
            // For multiple node support
            var $this = $(this),
                delayTimeout = false;

            function scrollToThis() {
                // Scroll to element
                $containers.animate({
                    scrollTop: $this.offset().top - options.offsetMargin
                }, {
                    duration: options.scrollPeriod,
                    complete: function () {
                        clearTimeout(delayTimeout);
                        delayTimeout = false;
                    }
                }).on("scroll.hoverScroll mousedown.hoverScroll DOMMouseScroll.hoverScroll mousewheel.hoverScroll keyup.hoverScroll", function () {
                    $containers.stop();
                });
            }

            var setup = function () {
                $this.on("mouseenter.hoverScroll", function () {
                    console.log("mouseenter.hoverScroll", $this);
                    if (delayTimeout === false) {
                        delayTimeout = setTimeout(scrollToThis, options.wait);
                    }
                }).on("mouseleave.hoverScroll", function () {
                    console.log("mouseleave.hoverScroll", $this);
                    clearTimeout(delayTimeout);
                    delayTimeout = false;
                });
            };

            setup();

        });
    }

    $.fn.truncate = truncate;
    $.fn.shorten = shorten;
    $.fn.hoverScroll = hoverScroll;
})(jQuery);

/* Masonry Helper */
(function ($) {
    "use strict";

    function setupMasonry() {
        return $(this).each(function () {
            var $this = $(this);
            try {
                $this.masonry('destroy');
            } catch (e) {
                console.warn("Failed to destroy masonry", $this);
            }
            $this.addClass("js-masonry").children().addClass("js-masonry-item");
            $this.masonry({
                itemSelector: '.js-masonry-item',
                isFitWidth: true,
                isAnimated: !(Modernizr && Modernizr.csstransitions),
                columnWidth: function (containerWidth) {
                    return Math.min(260, containerWidth);
                    //return Math.min(Math.min(containerWidth / 4, 320), containerWidth);
                }
            });
        });
    }


    $.fn.setupMasonry = setupMasonry;

})(jQuery);
