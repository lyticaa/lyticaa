require 'bootstrap'
require 'clipboard'

$ = require('jquery')
jQuery = require('jquery')
feather = require('feather-icons')

$(document).on 'turbolinks:load', ->
  (($) ->
    'use strict'
    $ ->
      body = $('body')
      sidebar = $('.sidebar')
      
      iconSidebar = (e) ->
        if e.matches
          body.addClass 'sidebar-folded'
        else
          body.removeClass 'sidebar-folded'
        return

      addActiveClass = (element) ->
        if current == ''
          #for root url
          if element.attr('href').indexOf('index.html') != -1
            element.parents('.nav-item').last().addClass 'active'
            if element.parents('.sub-menu').length
              element.closest('.collapse').addClass 'show'
              element.addClass 'active'
        else
          #for other url
          if element.attr('href').indexOf(current) != -1
            element.parents('.nav-item').last().addClass 'active'
            if element.parents('.sub-menu').length
              element.closest('.collapse').addClass 'show'
              element.addClass 'active'
            if element.parents('.submenu-item').length
              element.addClass 'active'
        return

      feather.replace()

      # Sidebar toggle to sidebar-folded
      $('.sidebar-toggler').on 'click', (e) ->
        $(this).toggleClass 'active'
        $(this).toggleClass 'not-active'
        if window.matchMedia('(min-width: 992px)').matches
          e.preventDefault()
          body.toggleClass 'sidebar-folded'
        else if window.matchMedia('(max-width: 991px)').matches
          e.preventDefault()
          body.toggleClass 'sidebar-open'
        return

      desktopMedium = window.matchMedia('(min-width:992px) and (max-width: 1199px)')
      desktopMedium.addListener iconSidebar
      iconSidebar desktopMedium
      current = location.pathname.split('/').slice(-1)[0].replace(/^\/|\/$/g, '')
      $('.nav li a', sidebar).each ->
        $this = $(this)
        addActiveClass $this
        return

      $('.horizontal-menu .nav li a').each ->
        $this = $(this)
        addActiveClass $this
        return

      #  open sidebar-folded when hover
      $('.sidebar .sidebar-body').hover (->
        if body.hasClass('sidebar-folded')
          body.addClass 'open-sidebar-folded'
        return
      ), ->
        if body.hasClass('sidebar-folded')
          body.removeClass 'open-sidebar-folded'
        return

      # initializing popover
      $('[data-toggle="popover"]').popover()

      # checkbox and radios
      $('.form-check label,.form-radio label').append '<i class="input-frame"></i>'

      # Horizontal menu in mobile
      $('[data-toggle="horizontal-menu-toggle"]').on 'click', ->
        $('.horizontal-menu .bottom-navbar').toggleClass 'header-toggled'
        return
      
      # Horizontal menu navigation in mobile menu on click
      navItemClicked = $('.horizontal-menu .page-navigation >.nav-item')
      navItemClicked.on 'click', (event) ->
        if window.matchMedia('(max-width: 991px)').matches
          if !$(this).hasClass('show-submenu')
            navItemClicked.removeClass 'show-submenu'
          $(this).toggleClass 'show-submenu'
        return
        
      $(window).scroll ->
        if window.matchMedia('(min-width: 992px)').matches
          header = $('.horizontal-menu')
          if $(window).scrollTop() >= 60
            $(header).addClass 'fixed-on-scroll'
          else
            $(header).removeClass 'fixed-on-scroll'
        return
      return
    return
  ) jQuery
  return
