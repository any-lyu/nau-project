module.exports = {
  root: true,
  env: {
    browser: true,
    es6: true,
  },
  extends: ['plugin:vue/essential', '@vue/standard'],
  plugins: ['prettier'],
  rules: {
    'no-extra-boolean-cast': 0,
    semi: 0,
    'no-eval': 0,
    'no-undef': 0,
    // allow paren-less arrow functions
    'arrow-parens': 0,
    // no-fallthrough
    'no-fallthrough': 0,
    'prefer-promise-reject-errors': 0,
    'space-before-function-paren': 0,
    // allow debugger during development
    'no-debugger': process.env.NODE_ENV === 'production' ? 2 : 0,
    eqeqeq: 0, // 必须使用全等
    'no-extend-native': 0,
    'no-multi-spaces': 0,
    indent: 0,
    'comma-dangle': [2, 'only-multiline'],
    'object-curly-spacing': 0,
    'no-mixed-operators': 0,
    'no-irregular-whitespace': 0,
    'no-cond-assign': 0,
  },
  parserOptions: {
    parser: "babel-eslint",
  },
}
