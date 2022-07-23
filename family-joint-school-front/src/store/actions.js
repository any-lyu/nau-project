import axios from 'axios'
import Config from '@/config'
axios.defaults.withCredentials = true

export const initUserInfo = ({ commit }, data) => {
  commit('initUserInfo', data)
}

export const toggleNavBar = ({ commit }, data) => {
  commit('toggleNavBar', data)
}

export const uploadFile = async ({ commit }, data) => {
  const param = new FormData()
  param.append('file', data.file)
  const config = {
    headers: { 'Content-Type': 'multipart/form-data' },
  }

  const result = await axios.post(Config.host + 'upload', param, config)
  if (result?.data?.code === 0) {
    return Promise.resolve({ url: result.data.data?.url })
  } else {
    console.log('上传失败:' + data.file)
    return Promise.resolve({ url: '' })
  }
}
