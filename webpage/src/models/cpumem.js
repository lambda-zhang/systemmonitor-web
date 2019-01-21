import {color1, color2, lineItemStyle, isupdated, areaStyleMem, areaStyleSwap, tooltipPercent, tooltipDefault} from '../consts/echarts.js'

export var getSysInfo = _getSysInfo

function _getSysInfo (_this, updatedat, cb) {
  _this.$http.get('/sysusage').then((res) => {
    var updatedts = 0
    var x = []
    var cpunum = 0

    if (!res || !res.status || res.status !== 200) {
      return cb && cb.call(_this, null, null, null, 'get cpuusage failed', updatedts)
    }
    var data = res.data
    cpunum = data.NumCpu
    var yMaxLoad = cpunum
    for (var i = 0; i < data.UpdatedAt.length; i++) {
      var ts = new Date(data.UpdatedAt[i])
      x.push('' + ts.getHours() + ':' + ts.getMinutes() + ':' + ts.getSeconds())
    }

    updatedts = isupdated(updatedat, data.UpdatedAtNewest)
    if (updatedts === 0) {
      return cb && cb.call(_this, null, null, null, null, updatedts)
    }

    var chartopt = {
      title: { text: 'CPU使用率' },
      animation: false,
      legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: ['CPU']},
      tooltip: tooltipPercent,
      xAxis: {boundaryGap: false, data: x},
      yAxis: {max: 100},
      series: [{name: 'CPU', symbol: 'none', type: 'line', color: ['#66AEDE'], smooth: 0.3, data: data.Cpu_percent, itemStyle: lineItemStyle}]
    }
    var chartopt2 = {
      title: {text: '内存使用率'},
      animation: false,
      xAxis: {boundaryGap: false, data: x},
      yAxis: {max: 100},
      legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: ['内存', 'SWAP']},
      tooltip: tooltipPercent,
      series: [{name: '内存', symbol: 'none', type: 'line', itemStyle: lineItemStyle, color: color2[0], smooth: 0.3, areaStyle: areaStyleMem, data: data.MemUsepercent},
        {name: 'SWAP', symbol: 'none', type: 'line', itemStyle: lineItemStyle, color: color2[1], smooth: 0.3, areaStyle: areaStyleSwap, data: data.SwapUsepercent}]
    }
    var chartopt3 = {
      title: {text: '系统平均负载 (' + cpunum + '核)'},
      animation: false,
      calculable: true,
      legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: ['load_1m', 'load_5m', 'load_15m']},
      tooltip: tooltipDefault,
      xAxis: {boundaryGap: false, data: x},
      yAxis: {max: Math.round(yMaxLoad + 1)},
      series: [
        {name: 'load_1m', symbol: 'none', type: 'line', color: color1[0], smooth: 0.3, data: data.Avg1min, itemStyle: lineItemStyle},
        {name: 'load_5m', symbol: 'none', type: 'line', color: color1[1], smooth: 0.3, data: data.Avg5min, itemStyle: lineItemStyle},
        {name: 'load_15m', symbol: 'none', type: 'line', color: color1[2], smooth: 0.3, data: data.Avg15min, itemStyle: lineItemStyle}
      ]
    }
    return cb && cb.call(_this, chartopt, chartopt2, chartopt3, null, updatedts)
  })
}
