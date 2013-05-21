(function($) {"use strict";

    $(document).ready(function() {

        $(".card-wrapper .card").shorten(300);
        $(".card-comment-wrapper .card-comment-popover-content").shorten(140);
        $(".truncate-this").truncate();
        $(".card-wrapper").setupMasonry();
        $(".card-comment-wrapper").setupMasonry();

        $('[data-toggle="tooltip"]').tooltip({
            placement : "bottom"
        });

    });
})(jQuery);
