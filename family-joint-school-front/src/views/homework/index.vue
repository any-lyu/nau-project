<template>
  <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
    <van-list v-model="loading" :finished="finished" finished-text="没有更多了" @load="onLoad">
      <van-cell-group v-for="item in list" :key="item.id">
        <van-cell @click.stop="go2Detail(item)">
          <template v-slot:title>
            <span style="font-size: 16px">{{ item.title }}</span>
            <span style="position: relative; float: right">{{ item.updated_at | datetime }}</span>
          </template>
          <template v-slot:label>
            <div>{{ item.content }}</div>
            <van-image width="60" height="60" lazy-load :src="item.pics && item.pics[0] ? item.pics[0].url : ''" />
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
  pageCName: '家庭作业',
  name: 'HomeworkList',
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
        url: '/api/homework/list',
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
    go2Detail(row) {
      this.$router.push({
        path: '/homework/detail',
        query: {
          id: row.id,
        },
      })
    },
  },
}
</script>
<style lang="less" scoped></style>
