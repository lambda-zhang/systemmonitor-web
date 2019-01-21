import Vue from 'vue'
import Router from 'vue-router'
import System from '@/views/system/system.vue'

Vue.use(Router)

var router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'system',
      component: System
    }
  ]
})

router.beforeEach((to, from, next) => {
  if (to.matched.length === 0) {
    from.name ? next({
      name: from.name
    }) : next('/')
  } else {
    next()
  }
})

export default router
