(function($) {"use strict";

    $(function() {

        $(".card-wrapper .card").shorten(300, true);
        $(".card-comment-wrapper .card-comment-popover-content").shorten(144, true);
        $(".truncate-this").truncate();
        $(".card-wrapper").setupMasonry();
        $(".card-comment-wrapper").setupMasonry();

        $('[data-toggle="tooltip"]').tooltip({
            placement : "bottom",
            container : "body" // Move to body to use the sans font
        });

        $("img").error(function() {
            console.log("Failed to load image", this);
            var $this = $(this), $parent = $this.parent();
            if ($parent.children.length === 1) {
                $parent.hide();
            } else {
                $this.hide();
            }
        });

        console.log("Init: JavaScript setup complete");

    });

})(jQuery);
