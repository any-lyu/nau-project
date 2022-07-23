/*
 * @Author: your name
 * @Date: 2021-07-12 17:25:55
 * @LastEditTime: 2021-07-12 17:40:21
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /ins-web/vue.config.js
 */
const path = require('path')
const packageInfo = require('./package.json')

const projectHost = {
  dev: `https://devstatic.ymm56.com/${packageInfo.serverName}`,
  beta: `https://qastatic.mbib.com.cn/${packageInfo.serverName}`,
  production: `https://static.mbib.com.cn/${packageInfo.serverName}`,
}

const configureWebpack = {
  resolve: {
    alias: {
      src: path.join(__dirname, 'src')
    }
  },
  externals: {
  },
  plugins: [
  ],
  resolveLoader: {
    extensions: ['.vue', '.js', '.json'],
  }
}

if (process.env.NODE_ENV === 'production') {
  configureWebpack.externals.vue = 'Vue'
  configureWebpack.externals['vue-router'] = 'VueRouter'
  configureWebpack.externals.vuex = 'Vuex'
}

module.exports = {
  publicPath: '/',
  // publicPath: process.env.NODE_ENV === 'production' ? `/${packageInfo.serverName}/` : '/',
  assetsDir: 'static',
  lintOnSave: process.env.NODE_ENV !== 'production',
  pages: {
    index: {
      entry: './src/main.js',
      filename: 'index.html',
      template: 'public/index.html',
      title: packageInfo.description,
      version: packageInfo.version,
      serverName: packageInfo.serverName
    },
  },
  configureWebpack,
  chainWebpack: config => {
    config.plugins.delete('preload-index')
    config.plugins.delete('prefetch-index')

    config.plugin('define').tap(args => {
      args[0]['process.env'].CODE_ENV = JSON.stringify(process.env.CODE_ENV)
      return args
    })

    // config.when(process.env.CODE_ENV === 'production', config => {
    //   config.optimization.minimizer('terser').tap(options => {
    //     options[0].terserOptions.compress.drop_console = true
    //     return options
    //   })
    // })

    // config.when(process.env.NODE_ENV === 'production', config => {
    //   config.module.rule('images').use('url-loader').tap(options => {
    //     options.fallback.options.publicPath = projectHost[process.env.CODE_ENV]
    //     return options
    //   })
    // })
  }
}
