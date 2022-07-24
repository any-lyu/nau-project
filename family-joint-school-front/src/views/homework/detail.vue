<template>
  <div class="homework-detail">
    <van-collapse v-model="activeNames">
      <van-collapse-item v-if="homework" :title="homework.title || ''" name="1">
        <div class="content">{{ homework.content }}</div>
        <van-image @click="previewImg(homework)" width="100" height="100" lazy-load :src="homework.pics && homework.pics[0] ? homework.pics[0].url : ''" />
      </van-collapse-item>
      <van-collapse-item style="margin-top: 16px" v-if="userHomework" :title="userHomework.title" name="1">
        <div class="content">{{ userHomework.content }}</div>
        <van-image
          @click="previewImg(userHomework)"
          width="100"
          height="100"
          lazy-load
          :src="userHomework.pics && userHomework.pics[0] ? userHomework.pics[0].url : ''"
        />
      </van-collapse-item>
    </van-collapse>
    <van-form v-if="!userHomework" style="margin-top: 24px" @submit="onSubmit">
      <van-field v-model="form.title" name="标题" label="标题" placeholder="标题" :rules="[{ required: true, message: '请填写标题' }]" />
      <van-field
        v-model="form.content"
        :rows="4"
        type="textarea"
        name="内容"
        label="内容"
        placeholder="内容"
        :rules="[{ required: true, message: '请填写内容' }]"
      />
      <van-field name="uploader" label="图片上传">
        <template v-slot:input>
          <van-uploader v-model="form.fileList" :after-read="afterRead" multiple :max-count="9" />
        </template>
      </van-field>
      <div style="margin: 16px 20vw">
        <van-button round block type="info" native-type="submit">提交</van-button>
      </div>
    </van-form>
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
  name: 'HomeworkDetail',
  mixins: [BasePage],
  data() {
    return {
      form: {
        title: '',
        content: '',
        fileList: [],
        pics: [],
      },
      homework: null,
      userHomework: null,
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
    afterRead(file) {
      if (Object.prototype.toString.call(file) === '[object Array]') {
        const arr = []
        file.forEach(f => {
          arr.push(this.$store.dispatch('uploadFile', f))
        })
        Promise.all(arr).then(result => {
          this.form.pics = this.form.pics.concat(result)
        })
      } else {
        this.$store.dispatch('uploadFile', file).then(r => {
          this.form.pics.push(r)
        })
      }
    },
    onDelete(file, detail) {
      this.form.pics.splice(detail.index, 1)
    },
    onSubmit() {
      Server({
        url: '/api/homework/submit',
        method: 'post',
        data: {
          homework_id: ~~this.$route.query?.id,
          title: this.form.title,
          content: this.form.content,
          pics: this.form.pics.filter(p => !!p.url),
        },
        needLoading: true,
      }).then(data => {
        this.$alert({
          message: '提交成功',
        })

        // this.$router.go(-1)
      })
    },
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
        this.homework = data.homework
        console.log(data)
        this.userHomework = data.user_homework
      })
    },
  },
}
</script>
<style lang="less" scoped>
.content {
  margin-bottom: 12px;
}
</style>
