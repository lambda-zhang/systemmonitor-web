import {dateformat3} from '../utils/dateformat.js'

export function updateCpuMemInfo (origin, doc, chartLen) {
  if (origin.latestx >= doc.updateat) {
    return origin
  }
  var info = origin
  if (!info.cpuusage || info.cpuusage.length < 1) {
    info.cpuusage = []
  }
  if (!info.memusage || info.memusage.length < 1) {
    info.memusage = []
  }
  if (!info.swapusage || info.swapusage.length < 1) {
    info.swapusage = []
  }
  if (!info.avg1min || info.avg1min.length < 1) {
    info.avg1min = []
  }
  if (!info.avg5min || info.avg5min.length < 1) {
    info.avg5min = []
  }
  if (!info.avg15min || info.avg15min.length < 1) {
    info.avg15min = []
  }
  if (!info.updateat || info.updateat.length < 1) {
    info.updateat = []
  }
  info.cpuusage.push(doc.cpuusage)
  info.memusage.push(doc.memusage)
  info.swapusage.push(doc.swapusage)
  info.avg1min.push(doc.avg1min)
  info.avg5min.push(doc.avg5min)
  info.avg15min.push(doc.avg15min)
  info.updateat.push(dateformat3(doc.updateat))

  for (var key in info) {
    var item2 = info[key]
    if (item2.length > chartLen) {
      for (var i = (item2.length - chartLen); i > 0; i--) {
        info[key].shift()
      }
    }
  }
  var max1 = Math.max(...info.avg1min)
  var max2 = Math.max(...info.avg5min)
  var max3 = Math.max(...info.avg15min)
  info.maxavg = Math.max(max1, max2, max3)
  info.latestx = doc.updateat
  return info
}
