import Vue from 'vue'
import Router from 'vue-router'

const originalPush = Router.prototype.push
Router.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

const PageRouter = require('@/components/PageRoute.js')
const component = function component(resolve) {
  if (component.installed) return resolve(PageRouter)
  component.installed = true
  setTimeout(resolve.bind(null, PageRouter))
}

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/:name/:name1/:name2/:name3/:name4/name5',
      component: component,
    },
    {
      path: '/:name/:name1/:name2/:name3/:name4',
      component: component,
    },
    {
      path: '/:name/:name1/:name2/:name3',
      component: component,
    },
    {
      path: '/:name/:name1/:name2',
      component: component,
    },
    {
      path: '/:name/:name1',
      component: component,
    },
    {
      path: '/:name',
      component: component,
    },
    {
      path: '/',
      component: component,
    },
  ],
})
