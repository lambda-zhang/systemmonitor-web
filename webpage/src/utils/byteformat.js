export function KBytes2Human (kb) {
  var mb = kb / 1024
  if (mb < 1) {
    return '' + kb + 'KB'
  }
  var gb = kb / 1024 / 1024
  if (gb < 1) {
    return '' + mb.toFixed(1) + 'MB'
  }
  var tb = kb / 1024 / 1024 / 1024
  if (tb < 1) {
    return '' + gb.toFixed(1) + 'GB'
  } else {
    return '' + tb.toFixed(1) + 'TB'
  }
}
