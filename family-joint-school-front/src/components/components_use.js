/* eslint-disable */
/**
 * 该文件是为了按需加载，剔除掉了一些不需要的框架组件。
 * 减少了编译支持库包大小
 *
 * 当需要更多组件依赖时，在该文件加入即可
 */
import Vue from 'vue'
import {
  Tabbar,
  TabbarItem,
  PullRefresh,
  List,
  Cell,
  CellGroup,
  Form,
  Field,
  Button,
  Uploader,
  NavBar,
  Col,
  Row,
  Image as VanImage,
  ImagePreview,
  Loading,
  Overlay,
  Divider,
  Toast,
  ShareSheet,
  Dialog,
  Collapse,
  CollapseItem,
  Lazyload,
  RadioGroup,
  Radio,
  Badge,
} from 'vant'

Vue.use(Tabbar)
Vue.use(TabbarItem)
Vue.use(PullRefresh)
Vue.use(List)
Vue.use(Cell)
Vue.use(CellGroup)
Vue.use(Form)
Vue.use(Field)
Vue.use(Button)
Vue.use(Uploader)
Vue.use(NavBar)
Vue.use(Col)
Vue.use(Row)
Vue.use(VanImage)
Vue.use(ImagePreview)
Vue.use(Loading)
Vue.use(Overlay)
Vue.use(Divider)
Vue.use(Toast)
Vue.use(ShareSheet)
Vue.use(Dialog)
Vue.use(Collapse)
Vue.use(CollapseItem)
Vue.use(Lazyload)
Vue.use(RadioGroup)
Vue.use(Radio)
Vue.use(Badge)

Vue.prototype.$toast = Toast
Vue.prototype.$dialog = Dialog
Vue.prototype.$alert = Dialog.alert
