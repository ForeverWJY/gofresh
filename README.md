# gofresh

golang 热编译工具，实时监测代码变动并编译运行. fork来源： https://github.com/gravityblast/fresh



用过很多热编译工具，各有各的坑，总之用起来问题不少，大多数坑在:

* 编译后杀死原进程不及时，运行新程序时候导致端口占用
* 要编译的包无法指定
* 要编译的临时目录无法指定
* 运行时参数无法指定
* 监控文件无法灵活指定
* 只编译不运行
* 卡顿延迟，意外退出



综合以上毛病，亲身体验, 在原fresh 的项目基础上小改一番，用了2年感觉很丝滑



# 使用


1, 全局安装 gofresh 命令
```bash
go get -u github.com/jangozw/gofresh
```



2, 在你的项目跟目录添加一个 ```.gofresh``` 的配置文件：



详情见 ```.gofresh``` 内容

```bash
touch .gofresh
```

添加内容：
```text

# 监控代码变动的目录 . 表示当前目录
watch_dir:          .

# 编译的文件
build_file :        cmd/app/main.go

# 存放编译文件的临时目录
tmp_path:          tmp

# 启动参数 模版: -p=conf/test.ini
run_arg:            -p=conf/test.ini

# 编译的二进制名称
build_name:        app-main-build

# 编译日志
build_log:         app-main-build-errors.log

# 监控的文件扩展名
valid_ext:         .go, .ini

# 这些扩展名的文件忽略
no_rebuild_ext:    .tpl, .tmpl, .html

# 忽略目录
ignored:           assets, tmp, deploy

# 编译前执行cmd脚本
pre_build_cmd:    docs.bat

# 监测到变动后编译延迟时间
build_delay:       600

colors:            1
log_color_main:    cyan
log_color_build:   yellow
log_color_runner:  green
log_color_watcher: magenta
log_color_app: green


```

3, 在项目根目录运行gofresh


```bash

gofresh

```


