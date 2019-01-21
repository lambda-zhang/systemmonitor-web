export function getWebSocketUrl () {
  if (window && window.location && window.location.hostname) {
    var wsaddr = (window.location.protocol === 'https') ? 'wss://' : 'ws://'
    var wsport = '9000'
    if (process.env.NODE_ENV !== 'development') {
      wsport = window.location.port.length > 0 ? window.location.port : '80'
    }

    wsaddr = wsaddr + window.location.hostname + ':' + wsport + '/ws'
    return wsaddr
  }
  return ''
}

export function getAPIUrl () {
  if (window && window.location && window.location.hostname) {
    var apiaddr = (window.location.protocol === 'https') ? 'https://' : 'http://'
    var apiport = '9000'
    if (process.env.NODE_ENV !== 'development') {
      apiport = window.location.port.length > 0 ? window.location.port : '80'
    }

    apiaddr = apiaddr + window.location.hostname + ':' + apiport
    return apiaddr
  }
  return ''
}
