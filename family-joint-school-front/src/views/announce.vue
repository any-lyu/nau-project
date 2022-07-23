<template>
  <div class="announce">
    <van-form @submit="onSubmit">
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
      <van-field name="radio" label="等级">
        <template #input>
          <van-radio-group v-model="form.level" direction="horizontal">
            <van-radio :name="1">高</van-radio>
            <van-radio :name="2">中</van-radio>
            <van-radio :name="3">低</van-radio>
          </van-radio-group>
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
  pageCName: '发布通知',
  mixins: [BasePage],
  name: 'Announce',
  data() {
    return {
      isShowNavBar: true,
      form: {
        title: '',
        content: '',
        level: 1,
        fileList: [],
      },
    }
  },
  methods: {
    onSubmit() {
      Server({
        url: '/cms/notice/publish',
        method: 'post',
        data: {
          title: this.form.title,
          content: this.form.content,
          level: this.form.level,
        },
        needLoading: true,
      }).then(data => {
        this.$alert({
          message: '提交成功',
        })
        this.form = {
          title: '',
          content: '',
          level: 1,
          fileList: [],
        }
      })
    },
  },
}
</script>
<style lang="less" scoped>
.announce {
}
</style>
