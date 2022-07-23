/**
 * 基本元数据配置
 */
/**
 * 功能模块权code限全局配置
 * 使用如<el-button v-if="hasPower(modulePowers.sendRedPackage)" type="warning">批量发红包</el-button>
 */
export const modulePowers = {
  sendRedPackage: ['2100012', '2100013'], // 发红包操作
  selectRedPackage: ['2100014'], // 红包查询
  // ... 其他模块权限code配置
}
