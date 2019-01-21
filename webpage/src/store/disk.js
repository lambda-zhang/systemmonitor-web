import {dateformat3} from '../utils/dateformat.js'

export function updateDiskInfo (origin, doc, chartLen) {
  if (origin.latestx >= doc.updateat) {
    return origin
  }
  var info = origin
  if (!info || !info.y || info.y.length < 1) {
    info.y = {}
  }
  for (var i = 0; i < doc.disks.length; i++) {
    var item = doc.disks[i]
    if (!info.y[item.name]) {
      info.y[item.name] = {}
      info.y[item.name]['rkbps'] = []
      info.y[item.name]['wkbps'] = []
      info.y[item.name]['rrps'] = []
      info.y[item.name]['wrps'] = []
      info.y[item.name]['usage'] = []
    }
    info.y[item.name]['rkbps'].push(item.rkbps)
    info.y[item.name]['wkbps'].push(item.wkbps)
    info.y[item.name]['rrps'].push(item.rrps)
    info.y[item.name]['wrps'].push(item.wrps)
    info.y[item.name]['usage'].push(item.usage)
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
  info.maxkbps = 0
  info.maxrps = 0
  for (var key in info.y) {
    var item1 = info.y[key]
    var max1 = Math.max(...item1.rkbps)
    var max2 = Math.max(...item1.wkbps)
    var max3 = Math.max(...item1.rrps)
    var max4 = Math.max(...item1.wrps)
    info.maxkbps = Math.max(10, info.maxkbps, max1, max2)
    info.maxrps = Math.max(10, info.maxrps, max3, max4)
    for (var key2 in item1) {
      var item2 = item1[key2]
      if (item2.length > chartLen) {
        for (i = (item2.length - chartLen); i > 0; i--) {
          info.y[key][key2].shift()
        }
      }
    }
  }
  return info
}
