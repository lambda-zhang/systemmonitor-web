import echarts from 'echarts'

export var color1 = ['#53FF53', '#FF5151', '#FF77FF', '#00FFFF', '#FFFF37', '#FF5809']
export var color2 = ['#00BB00', '#AE0000', '#D200D2', '#009393', '#A6A600', '#A23400']

export var lineItemStyle = {normal: {lineStyle: {width: 1.6}}}
export var areaStyleMem = {
  normal: {
    type: 'default',
    color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
      offset: 0,
      color: color1[0]
    }, {
      offset: 1,
      color: color1[0]
    }], false)
  }
}
export var areaStyleSwap = {
  normal: {
    type: 'default',
    color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
      offset: 0,
      color: color1[1]
    }, {
      offset: 1,
      color: color1[1]
    }], false)
  }
}
var formatterPercent = function (params) {
  var relVal = params[0].name
  for (var i = 0, l = params.length; i < l; i++) {
    relVal += '<br/>' + params[i].seriesName + ' : ' + params[i].value + ' %'
  }
  return relVal
}
export var tooltipPercent = {
  show: true,
  textStyle: {align: 'left'},
  trigger: 'axis',
  formatter: formatterPercent
}
export var tooltipDefault = {
  textStyle: {align: 'left'},
  show: true,
  trigger: 'axis'
}

export var isupdated = function (tslocal, tsfromserver) {
  var ts = new Date(tsfromserver).getTime()
  if (ts > tslocal) {
    return ts
  } else {
    return 0
  }
}
