(function($) {"use strict";

    $(document).ready(function() {

        $(".card-wrapper .card").shorten(300);
        $(".card-comment-wrapper .card-comment-popover-content").shorten(144);
        $(".truncate-this").truncate();
        $(".card-wrapper").setupMasonry();
        $(".card-comment-wrapper").setupMasonry();

        $('[data-toggle="tooltip"]').tooltip({
            placement : "bottom",
            container : "body" // Move to body to use the sans font
        });
        
        
        console.log("Init: JavaScript setup complete");

    });
    
})(jQuery);
