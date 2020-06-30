import AlertsHelper   from '../helpers/alerts'
import ChartsHelper   from '../helpers/charts'
import TemplateHelper from '../helpers/template'
import URLHelper      from '../helpers/url'
import FiltersHelper  from '../helpers/filters'

#
# Dashboard: Overview.
#
export default class DashboardOverview
  constructor: ->
    this.alerts = new AlertsHelper()
    this.charts = new ChartsHelper()
    this.template = new TemplateHelper()
    this.url = new URLHelper()
    this.filters = new FiltersHelper()

  #
  # Initialize.
  #
  init: ->
    this.load()
    this.filters.watchDate(this)
    this.reload()

    return

  #
  # Load.
  #
  load: ->
    this.alerts.reset()

    d = this

    $('button.loading').fadeIn()

    $.ajax(
      type: 'GET'
      url: d.url.clean() + 'dashboard/metrics/filter/' + $('.date-filter.active').data('range')
      timeout: 10000
      statusCode:
        200: (j) ->
          $('button.loading').fadeOut(400, ->
            if j.totalSales.line.categories[0].category.length == 0
              $('.alert.dashboard-total-sales-chart-error').fadeIn()
            else
              d.charts.line(
                'dashboard-total-sales-chart',
                'SALES',
                'DATE',
                j.totalSales.line.categories,
                j.totalSales.line.dataSets
              )
            )

          if j.unitsSold.chart.sparkline.data.length > 0
            d.charts.sparkline('dashboard-units-sold-sparkline', [j.unitsSold.chart.sparkline])
          else
            $('#dashboard-units-sold-sparkline').html('')

          if j.amazonCosts.chart.sparkline.data.length > 0
            d.charts.sparkline('dashboard-amazon-costs-sparkline', [j.amazonCosts.chart.sparkline])
          else
            $('#dashboard-amazon-costs-sparkline').html('')

          if j.productCosts.chart.sparkline.data.length > 0
            d.charts.sparkline('dashboard-product-costs-sparkline', [j.productCosts.chart.sparkline])
          else
            $('#dashboard-product-costs-sparkline').html('')

          if j.advertisingSpend.chart.sparkline.data.length > 0
            d.charts.sparkline('dashboard-advertising-spend-sparkline', [j.advertisingSpend.chart.sparkline])
          else
            $('#dashboard-advertising-spend-sparkline').html('')

          if j.refunds.chart.sparkline.data.length > 0
            d.charts.sparkline('dashboard-refunds-sparkline', [j.refunds.chart.sparkline])
          else
            $('#dashboard-refunds-sparkline').html('')

          if j.shippingCredits.chart.sparkline.data.length > 0
            d.charts.sparkline('dashboard-shipping-credits-sparkline', [j.shippingCredits.chart.sparkline])
          else
            $('#dashboard-shipping-credits-sparkline').html('')

          if j.promotionalRebates.chart.sparkline.data.length > 0
            d.charts.sparkline('dashboard-promotional-rebates-sparkline', [j.promotionalRebates.chart.sparkline])
          else
            $('#dashboard-promotional-rebates-sparkline').html('')

          if j.totalCosts.chart.sparkline.data.length > 0
            d.charts.sparkline('dashboard-total-costs-sparkline', [j.totalCosts.chart.sparkline])
          else
            $('#dashboard-total-costs-sparkline').html('')

          if j.grossMargin.chart.sparkline.data.length > 0
            d.charts.sparkline('dashboard-gross-margin-sparkline', [j.grossMargin.chart.sparkline])
          else
            $('#dashboard-gross-margin-sparkline').html('')

          if j.netMargin.chart.sparkline.data.length > 0
            d.charts.sparkline('dashboard-net-margin-sparkline', [j.netMargin.chart.sparkline])
          else
            $('#dashboard-net-margin-sparkline').html('')

          d.paintMetric(
            'dashboard-units-sold',
            j.unitsSold.total.value,
            j.unitsSold.total.diff
          )
          d.paintMetric(
            'dashboard-amazon-costs',
            j.amazonCosts.total.value,
            j.amazonCosts.total.diff
          )
          d.paintMetric(
            'dashboard-product-costs',
            j.productCosts.total.value,
            j.productCosts.total.diff
          )
          d.paintMetric(
            'dashboard-advertising-spend',
            j.advertisingSpend.total.value,
            j.advertisingSpend.total.diff
          )
          d.paintMetric(
            'dashboard-refunds',
            j.refunds.total.value,
            j.refunds.total.diff
          )
          d.paintMetric(
            'dashboard-shipping-credits',
            j.shippingCredits.total.value,
            j.shippingCredits.total.diff
          )
          d.paintMetric(
            'dashboard-promotional-rebates',
            j.promotionalRebates.total.value,
            j.promotionalRebates.total.diff
          )
          d.paintMetric(
            'dashboard-total-costs',
            j.totalCosts.total.value,
            j.totalCosts.total.diff
          )
          d.paintMetric(
            'dashboard-gross-margin',
            j.grossMargin.total.value,
            j.grossMargin.total.diff
          )
          d.paintMetric(
            'dashboard-net-margin',
            j.netMargin.total.value,
            j.netMargin.total.diff
          )

          d.template.renderIcons()
        422: ->
          $('button.loading').fadeOut(400, ->
            $('.alert.dashboard-load-error').fadeIn()
          )
        500: ->
          $('button.loading').fadeOut(400, ->
            $('.alert.dashboard-load-error').fadeIn()
          )
    ).fail ->
      $('button.loading').fadeOut(400, ->
        $('.alert.dashboard-load-error').fadeIn()
      )

    return

  #
  # Reload.
  #
  reload: ->
    p = this
    $('button.reload').on 'click', (e) ->
      p.load()
      return

    return

  #
  # Paint metric.
  #
  paintMetric: (cssClass, total, diff) ->
    $("h3.#{cssClass}").text(total)
    $("span.diff.#{cssClass}").text("#{diff}%")
    $("span.diff.#{cssClass}").next().remove()

    if diff > 0
      arrow = 'arrow-up'
      $("span.diff.#{cssClass}").parent().removeClass().addClass 'text-success'
    else if diff == 0
      arrow = 'minus'
      $("span.diff.#{cssClass}").parent().removeClass().addClass 'text-warning'
    else
      arrow = 'arrow-down'
      $("span.diff.#{cssClass}").parent().removeClass().addClass 'text-danger'

    $("span.diff.#{cssClass}").parent().append(
      $('<i>')
        .attr('data-feather', arrow)
        .addClass('icon-sm')
        .addClass('mb-1')
    )

    return
