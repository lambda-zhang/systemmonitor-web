import {color1, color2, lineItemStyle, isupdated, tooltipDefault} from '../consts/echarts.js'
export var getNetIfInfo = _getNetIfInfo
export var getTcpInfo = _getTcpInfo

function _getNetIfInfo (_this, updatedat, cb) {
  _this.$http.get('/netifusage').then((res) => {
    var updatedts = 0
    if (!res || !res.status || res.status !== 200) {
      return cb && cb.call(_this, null, null, 'get netifusage failed', updatedts)
    }
    var data = res.data
    var x = []
    updatedts = isupdated(updatedat, data.UpdatedAtNewest)
    if (updatedts === 0) {
      return cb && cb.call(_this, null, null, null, updatedts)
    }
    var netifname = []
    for (var keyname in data.InKBytes) {
      netifname.push(keyname)
    }
    for (var i = 0; i < data.UpdatedAt.length; i++) {
      var ts = new Date(data.UpdatedAt[i])
      x.push('' + ts.getHours() + ':' + ts.getMinutes() + ':' + ts.getSeconds())
    }
    var yData = []
    var yDataPackages = []
    var legendData = []
    for (var index = 0; index < netifname.length; index++) {
      var key = netifname[index]
      yData.push({name: 'in_' + key, symbol: 'none', type: 'line', color: [color1[index % color1.length]], smooth: 0.3, data: data.InKBytes[key], itemStyle: lineItemStyle})
      yData.push({name: 'out_' + key, symbol: 'none', type: 'line', color: [color2[index % color2.length]], smooth: 0.3, data: data.OutKBytes[key], itemStyle: lineItemStyle})
      yDataPackages.push({name: 'in_' + key, symbol: 'none', type: 'line', color: [color1[index % color1.length]], smooth: 0.3, data: data.InPackages[key], itemStyle: lineItemStyle})
      yDataPackages.push({name: 'out_' + key, symbol: 'none', type: 'line', color: [color2[index % color2.length]], smooth: 0.3, data: data.OutPackages[key], itemStyle: lineItemStyle})
      legendData.push('in_' + key)
    }

    var chartopt = {
      title: { text: '网络流入流出速率 KBps' },
      animation: false,
      calculable: true,
      legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: legendData},
      tooltip: tooltipDefault,
      xAxis: {boundaryGap: false, data: x},
      yAxis: {max: data.MaxKBytes + 1},
      series: yData
    }

    var chartopt2 = {
      title: { text: '网络流入流出数据包数 pps' },
      animation: false,
      calculable: true,
      legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: legendData},
      tooltip: tooltipDefault,
      xAxis: {boundaryGap: false, data: x},
      yAxis: {max: data.MaxPackages + 1},
      series: yDataPackages
    }
    return cb && cb.call(_this, chartopt, chartopt2, null, updatedts)
  })
}

function _getTcpInfo (_this, updatedat, cb) {
  _this.$http.get('/tcpusage').then((res) => {
    var updatedts = 0
    if (!res || !res.status || res.status !== 200) {
      return cb && cb.call(_this, null, 'get tcpusage failed', updatedts)
    }
    var data = res.data
    var x = []
    updatedts = isupdated(updatedat, data.UpdatedAtNewest)
    if (updatedts === 0) {
      return cb && cb.call(_this, null, null, updatedts)
    }
    for (var i = data.UpdatedAt.length - 1; i > -1; i--) {
      var ts = new Date(data.UpdatedAt[i])
      x.push('' + ts.getHours() + ':' + ts.getMinutes() + ':' + ts.getSeconds())
    }
    var chartopt = {
      title: { text: 'TCP连接数' },
      animation: false,
      legend: {orient: 'horizontal', 'x': 'center', 'y': 'bottom', icon: 'circle', itemHeight: 10, data: ['TCP_total', 'Established']},
      tooltip: tooltipDefault,
      xAxis: {boundaryGap: false, data: x},
      yAxis: {max: data.MaxTcpConnections + 1},
      series: [
        {name: 'TCP_total', symbol: 'none', type: 'line', color: color1[0], smooth: 0.3, data: data.TcpConnections, itemStyle: lineItemStyle},
        {name: 'Established', symbol: 'none', type: 'line', color: color1[1], smooth: 0.3, data: data.Established, itemStyle: lineItemStyle},
        {name: 'TcpListen', symbol: 'none', type: 'line', color: color1[2], smooth: 0.3, data: data.TcpListen, itemStyle: lineItemStyle}
      ]
    }
    return cb && cb.call(_this, chartopt, null, updatedts)
  })
}
