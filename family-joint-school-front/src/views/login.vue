<template>
  <div class="login-box">
    <h1 class="app-name">家校互联</h1>
    <van-form @submit="onSubmit">
      <van-field v-model="account" name="用户名" label="用户名" placeholder="用户名" :rules="[{ required: true, message: '请填写用户名' }]" />
      <van-field v-model="password" type="password" name="密码" label="密码" placeholder="密码" :rules="[{ required: true, message: '请填写密码' }]" />
      <div style="margin: 30px 0">
        <van-button round block type="info" native-type="submit">登录</van-button>
      </div>
    </van-form>
  </div>
</template>
<script>
import BasePage from '@/extend/BasePage'
import Server from '@/extend/Server'
export default {
  mixins: [BasePage],
  name: 'Login',
  data() {
    return {
      account: '',
      password: '',
    }
  },
  methods: {
    onSubmit() {
      Server({
        url: 'login',
        method: 'post',
        data: {
          account: this.account,
          password: this.password,
        },
      }).then(data => {
        const userInfo = {
          token: data.token,
          permission: data.permission,
          permissionName: data.permission_name,
          ...data.user,
        }
        window.localStorage.setItem('userInfo', JSON.stringify(userInfo))
        this.$store.dispatch('initUserInfo', userInfo)
        setTimeout(() => {
          this.$router.replace({ path: '/' })
        }, 300)
      })
    },
  },
}
</script>
<style lang="less" type="text/less" rel="stylesheet/less" scoped>
.login-box {
  display: flex;
  justify-content: center;
  align-items: center;
  box-sizing: border-box;
  flex-direction: column;
  margin-top: 20vh;
  .app-name {
    margin-top: 0;
    margin-bottom: 10vh;
  }
}
</style>
