<template>
  <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
    <van-list v-model="loading" :finished="finished" :immediate-check="false" finished-text="没有更多了" @load="onLoad">
      <van-cell-group v-for="item in list" :key="item.id">
        <van-cell :label="item.content">
          <template v-slot:title>
            <van-badge :color="['red', 'orange', 'transparent'][item.level - 1]" dot>
              <span style="font-size: 16px">{{ item.title }}</span>
            </van-badge>
            <span style="position: relative; float: right">{{ item.updated_at | datetime }}</span>
          </template>
        </van-cell>
      </van-cell-group>
    </van-list>
  </van-pull-refresh>
</template>
<script>
import BasePage from '@/extend/BasePage'
import Server from '@/extend/Server'
export default {
  isShowNavBar: true,
  pageCName: '学校通知',
  mixins: [BasePage],
  data() {
    return {
      list: [],
      loading: false,
      finished: false,
      refreshing: false,
      cursor: '',
    }
  },
  mounted() {
    this.onRefresh()
  },
  methods: {
    onLoad() {
      Server({
        url: '/api/notice/list',
        method: 'post',
        data: {
          cursor: this.cursor,
        },
      })
        .then(data => {
          this.cursor = data.cursor
          if (this.refreshing) {
            this.list = []
            this.refreshing = false
          }
          const arr = data?.list || []
          this.list.push(...arr)
          this.loading = false

          if (!arr.length) {
            this.finished = true
          }
        })
        .catch(() => {
          this.loading = false
          this.finished = true
        })
    },
    onRefresh() {
      this.cursor = ''
      // 清空列表数据
      this.finished = false

      // 重新加载数据
      // 将 loading 设置为 true，表示处于加载状态
      this.loading = true
      this.onLoad()
    },
  },
}
</script>
<style lang="less"></style>
