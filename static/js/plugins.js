// Avoid `console` errors in browsers that lack a console.
( function() {
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
    }());

// Place any jQuery/helper plugins in here.

//Truncate Plugin
/*
 * A plugin to truncate the text in a given node to the specified length
 * The full text is revealed on hover
 */
(function($) {"use strict";

    function truncate(length) {
        // Truncate the text inside the selected node
        $(this).each(function() {
            // For multiple node support
            var $this = $(this), text, newLength, shortenedText;
            text = $this.text();
            newLength = length || $this.attr("data-length") || 20;
            shortenedText = jQuery.trim(text).substring(0, newLength);
            if (shortenedText.length < text.length) {
                // Add a continuation symbol if text was clipped
                shortenedText = shortenedText + "...";
            }
            $this.text(shortenedText).hover(function() {
                // Show the full length text
                $this.text(text);
            }, function() {
                // Show the shortened text
                $this.text(shortenedText);
            });
        });
    }

    function shorten(height) {
        // Shorten the selected node to the given height
        $(this).each(function() {
            // For multiple node support
            var $this = $(this), height;
            height = height || $this.attr("data-height") || 240;
            $this.css({
                "overflow" : "hidden",
                "height" : height
            }).hover(function() {
                // Allow the element to expand to full height
                $this.css({
                    "overflow" : "visible",
                    "height" : "auto"
                });
            }, function() {
                // Shorten the element again
                $this.css({
                    "overflow" : "hidden",
                    "height" : height
                });
            });
        });
    }


    $.fn.truncate = truncate;
    $.fn.shorten = shorten;
})(jQuery);

/* Masonry Helper */
(function($) {"use strict";
    function setupMasonry() {
        $(this).each(function() {
            var $this = $(this);
            try {
                $this.masonry('destroy');
            } catch(e) {
                console.warn("Failed to destroy masonry", $this);
            }
            $this.masonry({
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
