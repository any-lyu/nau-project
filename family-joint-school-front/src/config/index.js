import BaseConfig from './base'
const Config = require('./' + process.env.CODE_ENV)
Object.assign(BaseConfig, Config.default)
export default BaseConfig
