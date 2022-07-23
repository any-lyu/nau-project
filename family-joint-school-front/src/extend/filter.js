import Vue from 'vue'
import dayjs from 'dayjs'

/**
 * 时间格式化
 */
Vue.filter('datetime', function (value, format) {
  if (!value) return
  return dayjs(value).format(format || 'YYYY-MM-DD HH:mm:ss')
})
