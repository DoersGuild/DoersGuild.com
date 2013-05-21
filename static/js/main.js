(function($) {"use strict";

    $(document).ready(function() {

        $(".card").shorten(300);
        $(".truncate-this").truncate();
        $(".card-wrapper").setupMasonry();

        $('[data-toggle="tooltip"]').tooltip({
            placement : "bottom"
        });

    });
})(jQuery);
