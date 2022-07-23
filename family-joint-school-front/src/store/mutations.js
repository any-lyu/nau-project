export const initUserInfo = (state, userInfo) => {
  state.userInfo = userInfo
}

export const toggleNavBar = (state, data) => {
  state.isShowNavBar = data.isShowNavBar
  state.navBarTitle = data.title
  state.isShowLeftArrow = data.isShowLeftArrow
  state.leftText = data.leftText
  console.log(data)
}
