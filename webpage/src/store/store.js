import Vue from 'vue'
import Vuex from 'vuex'
import {updateOsInfo} from './osinfo.js'
import {updateThermalInfo} from './thermal.js'
import {updateCpuMemInfo} from './cpumem.js'
import {updateNetWorkInfo, updateNetIfInfo} from './net.js'
import {updateDiskInfo} from './disk.js'

const chartLen = 50

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    socket: {
      isConnected: false,
      message: '',
      reconnectError: false
    },
    osinfo: {},
    thermalinfo: {latestx: 0},
    cpumeminfo: {latestx: 0},
    networkinfo: {latestx: 0},
    netifinfo: {latestx: 0},
    diskinfo: {latestx: 0}
  },
  mutations: {
    SOCKET_ONOPEN (state, event) {
      Vue.prototype.$socket = event.currentTarget
      state.socket.isConnected = true
      console.log('SOCKET_ONOPEN')
    },
    SOCKET_ONCLOSE (state, event) {
      state.socket.isConnected = false
      console.log('SOCKET_ONCLOSE')
    },
    SOCKET_ONERROR (state, event) {
      console.error(state, event)
      console.log('SOCKET_ONERROR')
    },
    // default handler called for all methods
    SOCKET_ONMESSAGE (state, message) {
      state.socket.message = message
      if (message && message.data) {
        var obj = JSON.parse(message.data)
        if (obj.os && obj.os.uptime > 0) {
          state.osinfo = updateOsInfo(state.osinfo, obj.os)
        }
        if (obj.thermal && obj.thermal.temp && obj.thermal.temp.length > 0) {
          state.thermalinfo = updateThermalInfo(state.thermalinfo, obj.thermal, chartLen)
        }
        if (obj.cpumem && obj.cpumem.updateat > 0) {
          state.cpumeminfo = updateCpuMemInfo(state.cpumeminfo, obj.cpumem, chartLen)
        }
        if (obj.network && obj.network.updateat > 0) {
          state.networkinfo = updateNetWorkInfo(state.networkinfo, obj.network, chartLen)
        }
        if (obj.netif && obj.netif.updateat > 0) {
          state.netifinfo = updateNetIfInfo(state.netifinfo, obj.netif, chartLen)
        }
        if (obj.disk && obj.disk.updateat > 0) {
          state.diskinfo = updateDiskInfo(state.diskinfo, obj.disk, chartLen)
        }
      }
    },
    // mutations for reconnect methods
    SOCKET_RECONNECT (state, count) {
      console.info(state, count)
    },
    SOCKET_RECONNECT_ERROR (state) {
      state.socket.reconnectError = true
      console.log('SOCKET_RECONNECT_ERROR')
    }
  },
  actions: {
    sendMessage: function (context, message) {
      Vue.prototype.$socket.send(message)
    }
  }
})
