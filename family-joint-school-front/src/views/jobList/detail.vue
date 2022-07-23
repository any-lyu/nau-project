<template>
  <div class="detail">
    <van-cell-group inset>
      <van-cell :title="detail.title" :label="detail.content" />
    </van-cell-group>
    <van-image @click="previewImg(detail)" width="100" height="100" lazy-load :src="detail.pics && detail.pics[0] ? detail.pics[0].url : ''" />
    <van-image-preview v-model="show" :images="images" @change="onChange" @close="onClose">
      <template v-slot:index>第{{ index + 1 }}页</template>
    </van-image-preview>
  </div>
</template>
<script>
import BasePage from '@/extend/BasePage'
import Server from '@/extend/Server'
export default {
  isShowNavBar: true,
  isShowLeftArrow: true,
  pageCName: '作业详情',
  name: 'JobDetail',
  mixins: [BasePage],
  data() {
    return {
      detail: {
        title: '语文背诵作业',
        content: '鹅鹅鹅，曲项向天歌，白毛浮绿水，红掌拨清波',
        pics: [
          {
            url: 'https://img.iplaysoft.com/wp-content/uploads/2019/free-images/free_stock_photo.jpg',
          },
        ],
      },
      activeNames: ['1'],
      show: false,
      images: [],
      index: 0,
    }
  },
  mounted() {
    this.getDetail()
  },
  methods: {
    onClose() {
      this.index = 0
      this.images = []
    },
    onChange(index) {
      this.index = index
    },
    previewImg(it) {
      this.show = true
      this.images = it?.pics?.map(p => p.url) || []
    },
    getDetail() {
      Server({
        url: '/api/homework/user/detail',
        method: 'post',
        data: {
          homework_id: ~~this.$route.query?.id,
        },
      }).then(data => {
        this.detail = data.homework
      })
    },
  },
}
</script>
<style lang="less" scoped>
.detail {
  padding: 0 12px 12px;
  /deep/ .van-cell__value {
    display: none;
  }
  .van-image {
    margin-top: 12px;
  }
}
</style>
