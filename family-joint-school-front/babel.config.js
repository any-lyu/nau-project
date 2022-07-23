module.exports = {
  comments: true,
  presets: [
    ['@vue/cli-plugin-babel/preset', { useBuiltIns: false }]
  ],
  plugins: [
    ['component', {
      libraryName: 'truck-lib',
      style: false
    }, 'truck-lib'],
    ['component', {
      libraryName: 'truck-mint-ui',
      style: true
    }, 'truck-mint-ui'],
    ['import', {
      libraryName: 'vant',
      libraryDirectory: 'es',
      style: true,
    }, 'vant']
  ]
}
