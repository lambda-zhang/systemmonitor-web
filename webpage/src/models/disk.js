import {color1, color2, lineItemStyle, isupdated, tooltipPercent, tooltipDefault} from '../consts/echarts.js'
export var getDiskInfo = _getDiskInfo

function _getDiskInfo (_this, updatedat, cb) {
  _this.$http.get('/diskusage').then((res) => {
    if (!res || !res.status || res.status !== 200) {
      return cb && cb.call(_this, null, null, null, 'get networkusage failed', 0)
    }
    var diskstate = res.data
    var legendData = []

    for (var keyname in diskstate.BytesUsedPermillage) {
      legendData.push(keyname)
    }
    var BytesUsedSeries = []
    var IOKBpsSeries = []
    var IOKBpsLegendData = []
    var IOKBpsMaxKBytes = diskstate.MaxKBytes
    var IORpsSeries = []
    var IORpsLegendData = []
    var IORpsMaxRequests = diskstate.MaxRequests
    var x = []

    var updatedts = isupdated(updatedat, diskstate.UpdatedAtNewest)
    if (updatedts === 0) {
      return cb && cb.call(_this, null, null, null, null, updatedts)
    }

    for (var i = 0; i < diskstate.UpdatedAt.length; i++) {
      var ts = new Date(diskstate.UpdatedAt[i])
      x.push('' + ts.getHours() + ':' + ts.getMinutes() + ':' + ts.getSeconds())
    }

    for (i = 0; i < legendData.length; i++) {
      var key = legendData[i]
      BytesUsedSeries.push({name: key, color: color1[i % color1.length], data: diskstate.BytesUsedPermillage[key], symbol: 'none', type: 'line', smooth: 0.3, itemStyle: lineItemStyle})
      IOKBpsSeries.push({name: key + '_read', color: color1[i % color1.length], data: diskstate.ReadKBytes[key], symbol: 'none', type: 'line', smooth: 0.3, itemStyle: lineItemStyle})
      IOKBpsSeries.push({name: key + '_write', color: color2[i % color2.length], data: diskstate.WriteKBytes[key], symbol: 'none', type: 'line', smooth: 0.3, itemStyle: lineItemStyle})
      IOKBpsLegendData.push(key + '_read')
      IOKBpsLegendData.push(key + '_write')

      IORpsSeries.push({name: key + '_read', color: color1[i % color1.length], data: diskstate.ReadRequests[key], symbol: 'none', type: 'line', smooth: 0.3, itemStyle: lineItemStyle})
      IORpsSeries.push({name: key + '_write', color: color2[i % color2.length], data: diskstate.WriteRequests[key], symbol: 'none', type: 'line', smooth: 0.3, itemStyle: lineItemStyle})
      IORpsLegendData.push(key + '_read')
      IORpsLegendData.push(key + '_write')
    }

    var chartoptusage = {
      title: {text: '硬盘使用率'},
      animation: false,
      xAxis: {boundaryGap: false, data: x},
      yAxis: {max: 100},
      legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: legendData},
      tooltip: tooltipPercent,
      series: BytesUsedSeries
    }
    var chartoptIOKBps = {
      title: {text: '磁盘读写数据量(KBps)'},
      animation: false,
      xAxis: {boundaryGap: false, data: x},
      yAxis: {max: IOKBpsMaxKBytes + 1},
      legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: IOKBpsLegendData},
      tooltip: tooltipDefault,
      series: IOKBpsSeries
    }
    var chartoptIORps = {
      title: {text: '磁盘读写请求次数(Rps)'},
      animation: false,
      xAxis: {boundaryGap: false, data: x},
      yAxis: {max: IORpsMaxRequests + 1},
      legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: IORpsLegendData},
      tooltip: tooltipDefault,
      series: IORpsSeries
    }
    return cb && cb.call(_this, chartoptusage, chartoptIOKBps, chartoptIORps, null, updatedts)
  })
}
