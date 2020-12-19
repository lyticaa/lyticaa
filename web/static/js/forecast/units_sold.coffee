import ForecastOverview   from './overview'

#
# Forecast: Total Sales.
#
export default class ForecastUnitsSold
  constructor: ->
    this.f = new ForecastOverview()

  #
  # Initialize.
  #
  init: ->
    this.f.init()
