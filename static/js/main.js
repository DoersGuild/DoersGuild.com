/* global jQuery, console */

(function ($) {
    "use strict";

    console.log("Main.js started");

    $(document).ready(function () {

        console.log("Main.js document ready");

        $(".card-wrapper .card").on('compressed.shorten', function () {
            // Truncate heading
            var $this = $(this);
            $this.find('.card-header').css({
                overflow: 'hidden',
                'white-space': 'nowrap'
            });
        }).on('expanded.shorten', function () {
            // Expand heading
            var $this = $(this);
            $this.find('.card-header').css({
                overflow: 'visible',
                'white-space': 'normal'
            });
        }).shorten(300, true);
        $(".card-comment-wrapper .card-comment-popover-content").on('compressed.shorten', function () {
            // Truncate heading
            var $this = $(this);
            $this.find('.popover-title').css({
                overflow: 'hidden',
                'white-space': 'nowrap'
            });
        }).on('expanded.shorten', function () {
            // Expand heading
            var $this = $(this);
            $this.find('.popover-title').css({
                overflow: 'visible',
                'white-space': 'normal'
            });
        }).shorten(144, true);
        $(".truncate-this").truncate();
        $(".card-wrapper").setupMasonry();
        $(".card-comment-wrapper").setupMasonry();

        $('[data-toggle="tooltip"]').tooltip({
            placement: "bottom",
            container: "body" // Move to body to use the sans font
        });

        $("img").error(function () {
            console.log("Failed to load image", this);
            var $this = $(this),
                $parent = $this.parent();
            if ($parent.children.length === 1) {
                $parent.hide();
            } else {
                $this.hide();
            }
        });

        console.log("Init: JavaScript setup complete");

    });

})(jQuery);
