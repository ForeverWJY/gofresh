# gofresh

golang 热编译工具，实时监测代码变动并编译运行. fork来源： https://github.com/gravityblast/fresh


原fresh 工具用起来不是很灵活，不能单独指定监控的目录和指定编译的文件，这在多个main文件项目开发中不方便，因此fork一份小改


# 使用


1, 全局安装 gofresh 命令
```bash
go get -u github.com/jangozw/gofresh
```



2, 在你的项目跟目录添加一个.gofresh 的配置文件, 内容如下:


```text
# 监控代码变动的目录 . 表示当前目录
watch_dir:          .

# 编译的文件
build_file :        cmds/app/main.go

# 存放编译文件的临时目录
tmp_path:          ./tmp

# 编译的二进制名称
build_name:        runner-build

# 编译日志
build_log:         runner-build-errors.log

# 监控的文件扩展名
valid_ext:         .go, .ini

# 这些扩展名的文件忽略
no_rebuild_ext:    .tpl, .tmpl, .html

# 忽略目录
ignored:           assets, tmp, deploy

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

gofresh 将按照.gofresh 中的配置实时监控代码变动并运行

```bash

gofresh

```

