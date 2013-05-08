(function($) {"use strict";

    $(document).ready(function() {

        $(".card").shorten(300);
        $(".truncate-this").truncate();
        $(".card-wrapper").setupMasonry();

        $('.more').hover(function() {
            $(this).addClass('btn btn-success')
        }, function() {
            $(this).removeClass('btn btn-success')
        });

        $('.expand1').hide();
        $('.more1').click(function() {
            if ($(".expand1").is(":visible")) {
                $('.more1').text("More");
            } else {
                $('.more1').text("Close");
            }
            $('.expand1').slideToggle('slow');

            return false;
        });

        $('.expand2').hide();
        $('.more2').click(function() {
            if ($(".expand2").is(":visible")) {
                $('.more2').text("More");
            } else {
                $('.more2').text("Close");
            }
            $('.expand2').slideToggle('slow');

            return false;
        });

        $('.expand3').hide();
        $('.more3').click(function() {
            if ($(".expand3").is(":visible")) {
                $('.more3').text("More");
            } else {
                $('.more3').text("Close");
            }
            $('.expand3').slideToggle('slow');

            return false;
        });

        $('.expand4').hide();
        $('.more4').click(function() {
            if ($(".expand4").is(":visible")) {
                $('.more4').text("More");
            } else {
                $('.more4').text("Close");
            }
            $('.expand4').slideToggle('slow');

            return false;
        });

        $('.expand5').hide();
        $('.more5').click(function() {
            if ($(".expand5").is(":visible")) {
                $('.more5').text("More");
            } else {
                $('.more5').text("Close");
            }
            $('.expand5').slideToggle('slow');

            return false;
        });

    });
})(jQuery);
