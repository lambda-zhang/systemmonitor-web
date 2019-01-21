import {dateformat3} from '../utils/dateformat.js'

export function updateNetWorkInfo (origin, doc, chartLen) {
  if (origin.latestx >= doc.updateat) {
    return origin
  }
  var info = origin
  if (!info.tcpall || info.tcpall.length < 1) {
    info.tcpall = []
  }
  if (!info.tcpest || info.tcpest.length < 1) {
    info.tcpest = []
  }
  if (!info.tcplis || info.tcplis.length < 1) {
    info.tcplis = []
  }
  if (!info.updateat || info.updateat.length < 1) {
    info.updateat = []
  }
  info.tcpall.push(doc.tcpall)
  info.tcpest.push(doc.tcpest)
  info.tcplis.push(doc.tcplis)
  info.updateat.push(dateformat3(doc.updateat))

  for (var key in info) {
    var item2 = info[key]
    if (item2.length > chartLen) {
      for (var i = (item2.length - chartLen); i > 0; i--) {
        info[key].shift()
      }
    }
  }
  info.maxtcp = Math.max(...info.tcpall)
  info.latestx = doc.updateat
  return info
}

export function updateNetIfInfo (origin, doc, chartLen) {
  if (origin.latestx >= doc.updateat) {
    return origin
  }
  var info = origin
  if (!info || !info.y || info.y.length < 1) {
    info.y = {}
  }
  for (var i = 0; i < doc.cards.length; i++) {
    var item = doc.cards[i]
    if (!info.y[item.name]) {
      info.y[item.name] = {}
      info.y[item.name]['inkb'] = []
      info.y[item.name]['inp'] = []
      info.y[item.name]['outkb'] = []
      info.y[item.name]['outp'] = []
    }
    info.y[item.name]['inkb'].push(item.inkb)
    info.y[item.name]['inp'].push(item.inp)
    info.y[item.name]['outkb'].push(item.outkb)
    info.y[item.name]['outp'].push(item.outp)
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
  info.maxkb = 0
  info.maxp = 0
  for (var key in info.y) {
    var item1 = info.y[key]
    var maxinkb = Math.max(...item1.inkb)
    var maxinp = Math.max(...item1.inp)
    var maxoutkb = Math.max(...item1.outkb)
    var maxoutp = Math.max(...item1.outp)
    info.maxkb = Math.max(info.maxkb, 10, maxinkb, maxoutkb)
    info.maxp = Math.max(info.maxp, 10, maxinp, maxoutp)

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
