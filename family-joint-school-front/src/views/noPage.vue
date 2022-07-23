<template>
  <div class="nopage">404</div>
</template>

<script>
import BasePage from '@/extend/BasePage'
export default {
  mixins: [BasePage],
  methods: {
    gohome() {
      this.$router.replace('/home')
    },
  },
}
</script>

<style lang="less" scoped>
.nopage {
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  .status-img {
    margin-top: 60px;
  }
  .nopage-text {
    margin-top: 100px;
  }
}
</style>
