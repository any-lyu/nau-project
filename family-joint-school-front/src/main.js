import Vue from 'vue'
import App from '@/App.vue'
import router from '@/router'
import store from '@/store'
import '@/extend/filter'
import '@/components/components_use.js'

Vue.config.productionTip = false

router.beforeEach(async (to, from, next) => {
  const isAuthenticated = store.getters.isLogin
  if (
    // 检查用户是否已登录
    !isAuthenticated &&
    // ❗️ 避免无限重定向
    to.path !== '/login'
  ) {
    // 将用户重定向到登录页面
    next({ path: '/login', replace: true })
  } else {
    next()
  }
})
Vue.prototype.Bus = new Vue()

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
