// Avoid `console` errors in browsers that lack a console.
(function() {
    var method;
    var noop = function noop() {
    };
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
(function($) {"use strict";

    function isMobile() {
        return $(window).width() < 480;
    }

    function truncate(length) {
        // Truncate the text inside the selected node
        return $(this).each(function() {
            // For multiple node support
            var $this = $(this), text, newLength, shortenedText;
            text = $this.text();
            newLength = length || $this.attr("data-length") || 20;
            shortenedText = jQuery.trim(text).substring(0, newLength);
            if (shortenedText.length < text.length) {
                // Add a continuation symbol if text was clipped
                shortenedText = shortenedText + "...";
            }

            var expand = function() {
                // Show the full length text
                $this.text(text);
            };
            var compress = function() {
                // Show the shortened text
                $this.text(shortenedText);
            };
            var toggle = function() {
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
        return $(this).each(function() {
            // For multiple node support
            var $this = $(this);

            var compress = function() {
                // Shorten the element again
                $this.css({
                    "overflow" : "hidden",
                    "height" : height
                });
            };

            var expand = function() {
                // Allow the element to expand to full height
                $this.css({
                    "overflow" : "visible",
                    "height" : "auto"
                });
            };

            var toggle = function() {
                // Toggle the height
                if ($this.css("height") === "auto") {
                    compress.apply(this, arguments);
                } else {
                    expand.apply(this, arguments);
                }
            };

            var setup = function() {
                var disable = notOnMobile && isMobile();
                height = height || $this.attr("data-height") || 240;
                if (!disable) {
                    $this.css({
                        "overflow" : "hidden",
                        "height" : height
                    }).on('mouseenter.shorten', expand).on('mouseleave.shorten', compress).on('focus.shorten', expand).on('blur.shorten', compress).on("tap", toggle);
                } else {
                    expand();
                    $this.off(".shorten");
                }
            };

            setup();
            $(window).on("resize", setup);

        });
    }


    $.fn.truncate = truncate;
    $.fn.shorten = shorten;
})(jQuery);

/* Masonry Helper */
(function($) {"use strict";
    function setupMasonry() {
        return $(this).each(function() {
            var $this = $(this);
            try {
                $this.masonry('destroy');
            } catch(e) {
                console.warn("Failed to destroy masonry", $this);
            }
            $this.addClass("js-masonry").children().addClass("js-masonry-item");
            $this.masonry({
                itemSelector : '.js-masonry-item',
                isFitWidth : true,
                isAnimated : !(Modernizr && Modernizr.csstransitions),
                columnWidth : function(containerWidth) {
                    return Math.min(260, containerWidth);
                    //return Math.min(Math.min(containerWidth / 4, 320), containerWidth);
                }
            });
        });
    }


    $.fn.setupMasonry = setupMasonry;

})(jQuery);
