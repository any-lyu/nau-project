# family-joint-school-front

##yarn.lock不能删,会影响包依赖版本,在部分安卓手机浏览器上会出现白屏

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run dev
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).


### sockjs 依赖版本

```

sockjs-client@^1.5.0:
  version "1.5.0"
  resolved "https://registry.npm.taobao.org/sockjs-client/download/sockjs-client-1.5.0.tgz#2f8ff5d4b659e0d092f7aba0b7c386bd2aa20add"
  dependencies:
    debug "^3.2.6"
    eventsource "^1.0.7"
    faye-websocket "^0.11.3"
    inherits "^2.0.4"
    json3 "^3.3.3"
    url-parse "^1.4.7"

sockjs@^0.3.21:
  version "0.3.21"
  resolved "https://registry.npm.taobao.org/sockjs/download/sockjs-0.3.21.tgz#b34ffb98e796930b60a0cfa11904d6a339a7d417"
  dependencies:
    faye-websocket "^0.11.3"
    uuid "^3.4.0"
    websocket-driver "^0.7.4"

```