// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import VueNativeSock from 'vue-native-websocket'
import store from './store/store'
import {getWebSocketUrl} from './utils/window_location.js'

import {post, get} from './utils/http.js'

Vue.config.productionTip = false

Vue.prototype.$http = {
  'get': get,
  'post': post
}

var wsaddr = getWebSocketUrl()
if (wsaddr && wsaddr.length > 1) {
  Vue.use(VueNativeSock, wsaddr, {
    store: store,
    reconnection: true,
    reconnectionAttempts: 5,
    reconnectionDelay: 3000
  })
}

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
