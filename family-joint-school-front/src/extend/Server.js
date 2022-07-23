import axios from 'axios'
import Vue from 'vue'
import Config from '@/config'
import store from '@/store'
const instance = axios.create({
  baseURL: Config.host,
  timeout: 2 * 60 * 1000,
  headers: {},
  trimNull: true, // 是否去除空值
  withCredentials: true, // default
  needLoading: false, // 是否需要加载效果
  ignoreCode: false, // 是否忽略服务端的错误提示
})

instance.interceptors.request.use(
  function (config) {
    if (config.needLoading) {
      Vue.prototype.Bus.$emit('loading.show')
    }
    const token = store.state.userInfo?.token
    if (token) {
      config.headers.Authorization = 'Bearer ' + token
    }
    config.data = config.data || {}
    if (config.trimNull && !(config.data instanceof window.FormData) && Object.prototype.toString.call(config.data) === '[object Object]') {
      const _data = Object.assign({}, config.data)
      isNull(_data)
      config.data = _data
    }
    if (config.download) {
      config.responseType = 'arraybuffer'
    }
    return config
    function isNull(obj) {
      if (obj !== undefined && obj !== null) {
        for (const [key, value] of Object.entries(obj)) {
          if (typeof value == 'object' && !(value instanceof Date)) isNull(value)
          if (typeof value != 'boolean' && !value && typeof value != 'number') {
            delete obj[key]
          }
        }
      }
    }
  },
  function (error) {
    return Promise.reject(error)
  }
)

instance.interceptors.response.use(
  function (response) {
    if (response.config.needLoading) {
      Vue.prototype.Bus.$emit('loading.hide')
    }
    const code = response.data.code
    if (code !== 0 && !response.config.ignoreCode) {
      Vue.prototype
        .$alert({
          message: response.data.msg,
        })
        .then(() => {})
      return Promise.reject(response)
    } else {
      return response.data?.data || {}
    }
  },
  function (error) {
    Vue.prototype.Bus.$emit('loading.hide')
    const status = error.response && error.response.status
    const isIgnore = error.response && error.response.data && error.response.data.ignoreCode
    if (status !== 200 && !isIgnore) {
      const message = error.message
      if (status == 401) {
        Vue.prototype
          .$alert({
            message: '认证失败, 请重新登录',
          })
          .then(() => {
            Vue.prototype.Bus.$emit('logout')
          })
      } else {
        Vue.prototype
          .$alert({
            message: message,
          })
          .then(() => {
            // on close
          })
      }
    }
    return Promise.reject(error)
  }
)

export default instance
