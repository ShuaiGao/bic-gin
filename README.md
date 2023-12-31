# bic-gin

使用 protoc-gen-bic 生成 gin 接口代码的示例项目，它对应的前端项目是 bic-vue

## 接口文档

本地swagger html文档 http://localhost:8080/swagger/index.html#/

本地swagger docs文档 http://localhost:8080/swagger/doc.json

## 使用protobuf

- 交互协议定义使用[git submodule,子模块](https://git-scm.com/book/zh/v2/Git-%E5%B7%A5%E5%85%B7-%E5%AD%90%E6%A8%A1%E5%9D%97)引入

- 项目地址 https://github.com/ShuaiGao/bic-proto

- 交互协议生成使用 protoc-gen-bic, 工具[github地址](https://github.com/ShuaiGao/protoc-gen-bic)

在使用 git pull 命令拉取代码后，项目中会存在一个 `.gitmodule` 文件，它存储了子模块的所有信息。但子模块代码不会自动拉取，需使用下面命令:

```shell
git submodule init
git submodule update
```

## 生成代码命令

参照 makefile 文件，gin 接口文件生成命令

```shell
make go
```

## license

bic-gin is licensed under the [MIT](https://github.com/ShuaiGao/bic-gin/blob/master/LICENSE) license