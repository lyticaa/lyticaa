import FusionCharts from 'fusioncharts/core'
import ZoomLine     from 'fusioncharts/viz/zoomline'
import Sparkline    from 'fusioncharts/viz/sparkline'
import FusionTheme  from 'fusioncharts/themes/es/fusioncharts.theme.fusion'

FusionCharts.addDep(ZoomLine)
FusionCharts.addDep(Sparkline)
FusionCharts.addDep(FusionTheme)

#
# Charts helper.
#
export default class ChartsHelper
  constructor: ->

  #
  # Line.
  #
  line: (container, yaxisname, xaxisname, categories, dataset) ->
    dataSource =
      chart:
        outCnvBaseFont: 'Overpass'
        outCnvBaseFontSize: '11'
        baseFont: 'Overpass'
        baseFontSize: '11'
        yaxisname: yaxisname
        xaxisname: xaxisname
        interactiveLegend: '1'
        forceaxislimits: '1'
        pixelsperpoint: '0'
        pixelsperlabel: '30'
        compactdatamode: '1'
        dataseparator: '|'
        theme: 'fusion'
      categories: categories
      dataset: dataset

    FusionCharts.ready ->
      myChart = new FusionCharts(
        type: 'zoomline'
        renderAt: container
        width: '100%'
        height: '400'
        dataFormat: 'json'
        dataSource: dataSource).render()
      return

    return

  #
  # Sparkline.
  #
  sparkline: (container, dataset) ->
    dataSource =
      chart:
        charttopmargin: '10'
        theme: 'fusion'
        showclosevalue: '1'
        showopenvalue: '1'
        setadaptiveymin: '1'
        dataset: dataset

    FusionCharts.ready ->
      myChart = new FusionCharts(
        type: 'sparkline'
        renderAt: container
        width: '100%'
        height: '100%'
        dataFormat: 'json'
        dataSource: dataSource).render()

    return
