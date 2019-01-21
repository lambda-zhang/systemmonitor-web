import {isupdated} from '../consts/echarts.js'
export var getOsInfo = _getOsInfo

function _getOsInfo (_this, updatedat, cb) {
  _this.$http.get('/osinfo').then((res) => {
    var updatedts = 0

    if (!res || !res.status || res.status !== 200) {
      return cb && cb.call(_this, null, 'get cpuusage failed', updatedts)
    }
    var data = res.data

    updatedts = isupdated(updatedat, data.UpdatedAt)
    if (updatedts === 0) {
      return cb && cb.call(_this, null, null, updatedts)
    }
    var netifname = []
    for (var keyname in data.TotalInKBytes) {
      netifname.push(keyname)
    }
    var netinfo = []
    for (var index = 0; index < netifname.length; index++) {
      var key = netifname[index]
      netinfo.push({
        name: key,
        TotalInKBytes: data.TotalInKBytes[key],
        TotalOutKBytes: data.TotalOutKBytes[key],
        TotalInPackages: data.TotalInPackages[key],
        TotalOutPackages: data.TotalOutPackages[key]
      })
    }
    data.netinfo = netinfo

    var diskname = []
    for (keyname in data.MBytesAll) {
      diskname.push(keyname)
    }
    var diskinfo = []
    for (index = 0; index < diskname.length; index++) {
      key = diskname[index]
      diskinfo.push({
        name: key,
        TotalMBytes: data.MBytesAll[key],
        BytesUsage: data.BytesUsage[key]
      })
    }
    data.diskinfo = diskinfo

    return cb && cb.call(_this, data, null, updatedts)
  })
}
