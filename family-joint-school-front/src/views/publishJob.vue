<template>
  <div class="publish-job">
    <van-form @submit="onSubmit">
      <van-field v-model="form.title" name="标题" label="标题" placeholder="标题" :rules="[{ required: true, message: '请填写标题' }]" />
      <van-field
        v-model="form.content"
        name="内容"
        label="内容"
        placeholder="内容"
        :rules="[{ required: true, message: '请填写内容' }]"
        :rows="4"
        type="textarea"
      />
      <van-field name="uploader" label="图片上传" :rules="[{ required: true, message: '请上传图片' }]">
        <template v-slot:input>
          <van-uploader v-model="form.fileList" :after-read="afterRead" @delete="onDelete" multiple :max-count="9" />
        </template>
      </van-field>
      <div style="margin: 16px">
        <van-button round block type="info" native-type="submit">提交</van-button>
      </div>
    </van-form>
  </div>
</template>
<script>
import BasePage from '@/extend/BasePage'
import Server from '@/extend/Server'
export default {
  isShowNavBar: true,
  pageCName: '发布作业',
  mixins: [BasePage],
  name: 'PublishJob',
  data() {
    return {
      form: {
        title: '',
        content: '',
        fileList: [],
        pics: [],
      },
    }
  },
  mounted() {},
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
        url: '/cms/homework/publish',
        method: 'post',
        data: {
          title: this.form.title,
          content: this.form.content,
          pics: this.form.pics.filter(p => !!p.url),
        },
        needLoading: true,
      }).then(data => {
        this.$alert({
          message: '提交成功',
        })
        this.form = {
          title: '',
          content: '',
          fileList: [],
          pics: [],
        }
      })
    },
  },
}
</script>
<style lang="less" scoped>
.publish-job {
}
</style>
