<template>
  <div class="app" id="app">
    <van-nav-bar fixed v-if="isShowNavBar" :title="navBarTitle" :left-text="leftText" :left-arrow="isShowLeftArrow" @click-left="handleLeftClick" />
    <div class="content" :class="{ 'is-show-navbar': isShowNavBar }">
      <router-view />
    </div>
    <van-tabbar v-show="isShowTabbar" v-model="active" fixed>
      <van-tabbar-item v-for="b in tabBars" replace :to="b.path" :icon="b.icon" :key="b.path" :badge="getBadge(b)">{{ b.label }}</van-tabbar-item>
    </van-tabbar>
    <van-overlay :show="isLoading">
      <van-loading class="loading" size="24px" type="spinner" color="#1989fa" text-color="#0094ff">加载中...</van-loading>
    </van-overlay>
  </div>
</template>
<script>
import { mapState, mapGetters } from 'vuex'
export default {
  data() {
    return {
      active: -1,
      bars: [],
      isLoading: false,
    }
  },
  computed: {
    ...mapState(['noticeNum', 'homeworkNum', 'isShowNavBar', 'isShowLeftArrow', 'navBarTitle', 'leftText']),
    ...mapGetters(['isLogin', 'isAdmin', 'tabBars']),
    isShowTabbar({ isLogin, $route }) {
      return isLogin && $route.path !== '/login'
    },
  },
  watch: {
    $route(to) {
      this.setActiveIndex()
    },
  },
  created() {
    this.Bus.$on('loading.show', () => {
      this.isLoading = true
    })
    this.Bus.$on('loading.hide', () => {
      this.isLoading = false
    })
    this.Bus.$on('logout', () => {
      this.$router.replace('/login')
      window.localStorage.setItem('userInfo', '{}')
      this.$store.dispatch('initUserInfo', {})
    })
    this.setActiveIndex()
  },
  methods: {
    getBadge(item) {
      if (item.path === '/notice') {
        return this.noticeNum || ''
      }
      if (item.path === '/homework') {
        return this.homeworkNum || ''
      }
      return ''
    },
    setActiveIndex() {
      if (this.$route.path === '/') {
        this.active = 0
        return
      }
      const arr = this.isAdmin ? ['/publishJob', '/announce', 'jobList', '/userCenter'] : ['/notice', '/homework', '/userCenter']
      arr.forEach((item, idx) => {
        if (this.$route.path.indexOf(item) !== -1) {
          this.active = idx
        }
      })
    },
    handleLeftClick() {
      this.$router.go(-1)
    },
  },
}
</script>
<style lang="less">
@import '~@/assets/style/index.less';
.app {
  position: absolute;
  width: 100%;
  background-color: #f2f3f5;
  font-size: 0.32rem;
  .content {
    box-sizing: border-box;
    margin-top: 0;
    padding-bottom: 50px;
    &.is-show-navbar {
      margin-top: 62px;
    }
  }
  .loading {
    position: absolute;
    transform: translate(-50%, -50%);
    top: 50%;
    left: 50%;
  }
}
</style>
