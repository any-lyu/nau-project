const bars = {
  0: [
    { path: '/notice', label: '学校通知', icon: 'comment-o' },
    { path: '/homework', label: '家庭作业', icon: 'shopping-cart-o' },
    { path: '/userCenter', label: '个人中心', icon: 'contact' },
  ],
  1: [
    { path: '/publishJob', label: '发布作业', icon: 'home-o' },
    { path: '/announce', label: '发布通知', icon: 'comment-circle-o' },
    { path: '/jobList', label: '作业列表', icon: 'todo-list-o' },
    { path: '/userCenter', label: '个人中心', icon: 'contact' },
  ],
}

export const isLogin = state => {
  return !!state.userInfo?.token
}

export const isAdmin = state => {
  return state.userInfo?.permission === 2
}

export const tabBars = state => {
  return bars[state.userInfo?.permission === 2 ? 1 : 0]
}
