import {dateformat1, dateformat2} from '../utils/dateformat.js'
import {KBytes2Human} from '../utils/byteformat.js'

export function updateOsInfo (origin, doc) {
  if (origin.uptime >= doc.origin) {
    return origin
  }
  var info = doc

  info.memtotalmb = KBytes2Human(doc.memtotalmb * 1024)
  info.swaptotalmb = KBytes2Human(doc.swaptotalmb * 1024)
  for (var i = 0; i < info.netinfo.length; i++) {
    var item = info.netinfo[i]
    info.netinfo[i].inkb = KBytes2Human(item.inkb)
    info.netinfo[i].outkb = KBytes2Human(item.outkb)
  }
  for (var j = 0; j < info.diskinfo.length; j++) {
    info.diskinfo[j].totalmb = KBytes2Human(info.diskinfo[j].totalmb * 1024)
  }

  info.uptime = dateformat1(info.uptime)
  info.starttime = dateformat2(info.starttime)
  if (info.diskinfo) {
    info.diskinfo.sort(function (a, b) {
      return a.name > b.name
    })
  }
  if (info.netinfo) {
    info.netinfo.sort(function (a, b) {
      return a.name > b.name
    })
  }

  return info
}
