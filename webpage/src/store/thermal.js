import {dateformat3} from '../utils/dateformat.js'

export function updateThermalInfo (origin, doc, chartLen) {
  if (origin.latestx >= doc.updateat) {
    return origin
  }
  var info = origin
  if (!info || !info.y || info.y.length < 1) {
    info.y = {}
  }
  for (var i = 0; i < doc.temp.length; i++) {
    var item = doc.temp[i]
    if (!info.y[item.name]) {
      info.y[item.name] = []
    }
    info.y[item.name].push(item.val)
  }
  if (!info.x || info.x.length < 1) {
    info.x = []
  }
  info.x.push(dateformat3(doc.updateat))
  info.latestx = doc.updateat

  if (info.x.length > chartLen) {
    for (i = (info.x.length - chartLen); i > 0; i--) {
      info.x.shift()
    }
  }
  info.maxTemp = 0
  for (var key in info.y) {
    var item2 = info.y[key]
    if (item2.length > 0) {
      var max = Math.max(...item2)
      info.maxTemp = Math.max(max, info.maxTemp)
    }
    if (item2.length > chartLen) {
      for (i = (item2.length - chartLen); i > 0; i--) {
        info.y[key].shift()
      }
    }
  }
  return info
}
