require("bootstrap");
require("clipboard");

import jQuery from "jquery";
import feather from "feather-icons"

(function($) {
    'use strict';
    $(function() {
        let body = $('body');
        let sidebar = $('.sidebar');

        // Enable feather-icons with SVG markup
        feather.replace();

        // Sidebar toggle to sidebar-folded
        $('.sidebar-toggler').on('click', function(e) {
            $(this).toggleClass('active');
            $(this).toggleClass('not-active');
            if (window.matchMedia('(min-width: 992px)').matches) {
                e.preventDefault();
                body.toggleClass('sidebar-folded');
            } else if (window.matchMedia('(max-width: 991px)').matches) {
                e.preventDefault();
                body.toggleClass('sidebar-open');
            }
        });

        // sidebar-folded on large devices
        function iconSidebar(e) {
            if (e.matches) {
                body.addClass('sidebar-folded');
            } else {
                body.removeClass('sidebar-folded');
            }
        }
        var desktopMedium = window.matchMedia('(min-width:992px) and (max-width: 1199px)');
        desktopMedium.addListener(iconSidebar);
        iconSidebar(desktopMedium);

        //Add active class to nav-link based on url dynamically
        function addActiveClass(element) {
            if (current === "") {
                //for root url
                if (element.attr('href').indexOf("index.html") !== -1) {
                    element.parents('.nav-item').last().addClass('active');
                    if (element.parents('.sub-menu').length) {
                        element.closest('.collapse').addClass('show');
                        element.addClass('active');
                    }
                }
            } else {
                //for other url
                if (element.attr('href').indexOf(current) !== -1) {
                    element.parents('.nav-item').last().addClass('active');
                    if (element.parents('.sub-menu').length) {
                        element.closest('.collapse').addClass('show');
                        element.addClass('active');
                    }
                    if (element.parents('.submenu-item').length) {
                        element.addClass('active');
                    }
                }
            }
        }

        var current = location.pathname.split("/").slice(-1)[0].replace(/^\/|\/$/g, '');
        $('.nav li a', sidebar).each(function() {
            var $this = $(this);
            addActiveClass($this);
        });

        $('.horizontal-menu .nav li a').each(function() {
            var $this = $(this);
            addActiveClass($this);
        })


        //  open sidebar-folded when hover
        $(".sidebar .sidebar-body").hover(
            function () {
                if (body.hasClass('sidebar-folded')){
                    body.addClass("open-sidebar-folded");
                }
            },
            function () {
                if (body.hasClass('sidebar-folded')){
                    body.removeClass("open-sidebar-folded");
                }
            });

        // close sidebar when click outside on mobile/table
        $(document).on('click touchstart', function(e){
            e.stopPropagation();

            // closing of sidebar menu when clicking outside of it
            if (!$(e.target).closest('.sidebar-toggler').length) {
                var sidebar = $(e.target).closest('.sidebar').length;
                var sidebarBody = $(e.target).closest('.sidebar-body').length;
                if (!sidebar && !sidebarBody) {
                    if ($('body').hasClass('sidebar-open')) {
                        $('body').removeClass('sidebar-open');
                    }
                }
            }
        });

        // initializing popover
        $('[data-toggle="popover"]').popover();

        //checkbox and radios
        $(".form-check label,.form-radio label").append('<i class="input-frame"></i>');

        //Horizontal menu in mobile
        $('[data-toggle="horizontal-menu-toggle"]').on("click", function() {
            $(".horizontal-menu .bottom-navbar").toggleClass("header-toggled");
        });
        // Horizontal menu navigation in mobile menu on click
        var navItemClicked = $('.horizontal-menu .page-navigation >.nav-item');
        navItemClicked.on("click", function(event) {
            if(window.matchMedia('(max-width: 991px)').matches) {
                if(!($(this).hasClass('show-submenu'))) {
                    navItemClicked.removeClass('show-submenu');
                }
                $(this).toggleClass('show-submenu');
            }
        })

        $(window).scroll(function() {
            if(window.matchMedia('(min-width: 992px)').matches) {
                var header = $('.horizontal-menu');
                if ($(window).scrollTop() >= 60) {
                    $(header).addClass('fixed-on-scroll');
                } else {
                    $(header).removeClass('fixed-on-scroll');
                }
            }
        });

    });
})(jQuery);
