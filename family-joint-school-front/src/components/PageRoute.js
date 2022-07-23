import Vue from 'vue'
import noPage from '@/views/noPage'

export default {
  data() {
    return {
      instance: null,
    }
  },
  watch: {
    $route: {
      async handler(_, oldRoute) {
        this.instance = await this.loadPage(oldRoute)
      },
      immediate: true,
    },
  },
  computed: {
    state() {
      return this.$store.state
    },
  },
  methods: {
    getPath() {
      const params = this.$route.params
      const arr = Object.values(params)
      return arr.join('/')
    },
    async loadPage(oldRoute) {
      const path = this.getPath() || (this.$store.getters.isAdmin ? 'publishJob' : 'notice')
      try {
        const module = await import('../views/' + path + '')
        const tempModule = Vue.extend(module.default)
        return tempModule.extend({
          data() {
            return {
              oldRoute$$: oldRoute,
            }
          },
        })
      } catch (e) {
        return noPage
      }
    },
  },
  render(createElement) {
    return createElement(this.instance)
  },
}
