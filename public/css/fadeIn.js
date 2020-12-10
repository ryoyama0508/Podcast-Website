$(window).on('load scroll', function () {

    var box = $('.fadeIn_from_right');
    var animated = 'animated';

    box.each(function () {

        var boxOffset = $(this).offset().top;
        var scrollPos = $(window).scrollTop();
        var wh = $(window).height();

        if (scrollPos > boxOffset - wh + 100) {
            $(this).addClass(animated);
        }
    });

});

$(window).on('load scroll', function () {

    var box = $('.fadeIn_from_left');
    var animated = 'animated';

    box.each(function () {

        var boxOffset = $(this).offset().top;
        var scrollPos = $(window).scrollTop();
        var wh = $(window).height();

        if (scrollPos > boxOffset - wh + 20) {
            $(this).addClass(animated);
        }
    });

});