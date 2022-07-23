<template>
  <div class="user-center">
    <section class="user-info">
      <div class="header-img">
        <van-image round width="1rem" height="1rem" :src="userInfo.avatar" />
      </div>
      <div class="brief">
        <div class="user-name">{{ userInfo.nickname }}</div>
        <div>{{ userInfo.mobile }}</div>
      </div>
    </section>
    <section class="sub-column">
      <van-cell icon="shop-o" title="意见反馈" is-link @click="$toast('意见反馈')" />
      <van-divider />
      <van-cell icon="shop-o" title="关于我们" is-link @click="$toast('关于我们')" />
      <van-divider />
      <van-cell icon="shop-o" title="分享" is-link @click="showShare = true" />
      <van-divider />
    </section>
    <div style="margin: 30px">
      <van-button round block type="info" @click="logout">退出</van-button>
    </div>
    <van-share-sheet v-model="showShare" title="立即分享给好友" :options="options" @select="onSelect" />
  </div>
</template>
<script>
import BasePage from '@/extend/BasePage'
import { mapState } from 'vuex'
export default {
  isShowNavBar: true,
  pageCName: '个人中心',
  mixins: [BasePage],
  name: 'UserCenter',
  data() {
    return {
      showShare: false,
      options: [
        [
          { name: '微信', icon: 'wechat' },
          { name: '微博', icon: 'weibo' },
          { name: 'QQ', icon: 'qq' },
        ],
        [
          { name: '复制链接', icon: 'link' },
          { name: '二维码', icon: 'qrcode' },
          { name: '小程序码', icon: 'weapp-qrcode' },
        ],
      ],
    }
  },
  computed: {
    ...mapState(['userInfo']),
  },
  mounted() {},
  methods: {
    onSelect(option) {
      this.$toast(option.name)
      this.showShare = false
    },
    logout() {
      this.Bus.$emit('logout')
    },
  },
}
</script>
<style lang="less" scoped>
.user-center {
  .user-info {
    background-color: #fff;
    padding: 16px 0;
    height: 1rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
    .header-img {
      width: 1.5rem;
      display: flex;
      justify-content: center;
      align-items: center;
      flex-shrink: 0;
    }
    .brief {
      width: 100%;
    }
  }
  .user-name {
    font-size: 0.48rem;
    font-weight: 700;
  }
  .sub-column {
    margin: 12px 0;
  }
  .van-divider {
    margin: 0;
  }
  /deep/ .van-share-sheet__options {
    justify-content: space-evenly;
  }
}
</style>
