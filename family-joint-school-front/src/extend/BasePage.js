/**
 * 所有页面基础类
 */

export default {
  isShowNavBar: false,
  isShowLeftArrow: false,
  leftText: '',
  /* 父类向下传递的参数数据 */
  props: [],
  data() {
    return {}
  },
  // 组件是你刚被创建,组件属性计算前
  beforeCreated() {
    console.log('component---beforeCreated', this.$options.name)
  },
  // 组件创建完成,属性已绑定,但是dom还没生产,$el还不存在
  created() {
    this.toggleNavBar({
      isShowNavBar: this.$options.isShowNavBar,
      title: this.$options.pageCName,
      isShowLeftArrow: this.$options.isShowLeftArrow,
      leftText: this.$options.isShowLeftArrow ? this.$options.leftText : '',
    })
  },
  // 模板编译挂载前
  beforeMount() {},
  // 模板编译挂载之后,不保证组件已经在document中。
  mounted() {},
  // 组件更新之前
  beforeUpdate() {},
  // 组件更新之后
  updated() {},
  // keep-alive 会用到
  // 组件被激活
  activated() {},
  // 组件被移除
  deactivated() {},
  beforeDestroy() {},
  destroyed() {},
  methods: {
    toggleNavBar(options) {
      this.$store.dispatch('toggleNavBar', options)
    },
  },
}
