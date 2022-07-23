import Vue from 'vue'
import Vuex from 'vuex'
import * as mutations from './mutations'
import * as actions from './actions'
import * as getters from './getters'

Vue.use(Vuex)
let userInfo = {}
try {
  userInfo = JSON.parse(window.localStorage.getItem('userInfo'))
} catch (err) {
  console.log(err)
}

export default new Vuex.Store({
  state: {
    userInfo: userInfo,
    noticeNum: 0,
    homeworkNum: 0,
    isShowNavBar: false,
    isShowLeftArrow: false,
    navBarTitle: '',
    leftText: '',
  },
  getters: getters,
  mutations: mutations,
  actions: actions,
})
