import {color1, lineItemStyle, isupdated, tooltipDefault} from '../consts/echarts.js'
export var getThermalInfo = _getThermalInfo

function _getThermalInfo (_this, updatedat, cb) {
  _this.$http.get('/thermalstate').then((res) => {
    var updatedts = 0
    if (!res || !res.status || res.status !== 200) {
      return cb && cb.call(_this, null, 'get ThermalInfo failed', updatedts)
    }
    var data = res.data
    var x = []
    updatedts = isupdated(updatedat, data.UpdatedAtNewest)
    if (updatedts === 0) {
      return cb && cb.call(_this, null, null, updatedts)
    }
    var thermalname = []
    for (var keyname in data.Temp) {
      thermalname.push(keyname)
    }
    for (var i = 0; i < data.UpdatedAt.length; i++) {
      var ts = new Date(data.UpdatedAt[i])
      x.push('' + ts.getHours() + ':' + ts.getMinutes() + ':' + ts.getSeconds())
    }
    var yData = []
    var legendData = []
    for (var index = 0; index < thermalname.length; index++) {
      var key = thermalname[index]
      yData.push({name: key, symbol: 'none', type: 'line', color: [color1[index % color1.length]], smooth: 0.3, data: data.Temp[key], itemStyle: lineItemStyle})
      legendData.push(key)
    }
    console.log(yData)

    var chartopt = {
      title: { text: '主板温度(℃)' },
      animation: false,
      calculable: true,
      legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: legendData},
      tooltip: tooltipDefault,
      xAxis: {boundaryGap: false, data: x},
      yAxis: {max: data.MaxTemp + 1},
      series: yData
    }
    return cb && cb.call(_this, chartopt, null, updatedts)
  })
}
