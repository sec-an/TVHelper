# 影视助手

## `v1.0.0` -> `v1.1.1` 升级说明

- 可以使用一键脚本更新
- `v1.0.0`项目结构不够规范，`v1.1.1`进行了规范化
- 如果已经安装了`v1.0.0`版本，请将项目根目录下的`config`、`live`、`source_config`目录移至`configs`目录下
- `configs/config.yaml`可修改TVHelper启动端口、启用Redis进行缓存等
- 项目运行日志默认会存储在`logs`目录下，请确保TVHelper对该目录具有写权限

## 痛点

- 使用在线配置，不方便对配置进行个性化修改
- 在线配置缓存至本地，担心更新不及时

## 功能（详见`configs/config/sample.json`）

- 极高的自定义程度
- 多源整合、处理（在线缝合）
  - `http://你的IP:16214/config/sample`
  - 去重
  - :star2:失效自动切换备用源
  - 多`Jar`
  - 同时支持本地、在线订阅
    - `http://你的IP:16214/config/src/demo/config.json`
    - `http://你的IP:16214/config/src/demo/custom_spider.jar;md5;a84fef826cb82da525469e8acf1e7d9a`
- 点播源黑名单:u7981:，指定名称的点播源不再展示
- 直播源替换，本地文件服务器
  - `http://你的IP:16214/live/IPTV.m3u`
- 豆瓣主页
  - `http://你的IP:16214/home?douban=你的豆瓣id`
- ...

## 一键脚本

### 安装

```shell
curl -fsSL "https://gh-proxy.com/https://raw.githubusercontent.com/sec-an/TVHelper/main/v1.sh" | bash -s install
```

### 更新

```shell
curl -fsSL "https://gh-proxy.com/https://raw.githubusercontent.com/sec-an/TVHelper/main/v1.sh" | bash -s update
```

### 卸载

```shell
curl -fsSL "https://gh-proxy.com/https://raw.githubusercontent.com/sec-an/TVHelper/main/v1.sh" | bash -s uninstall
```

### 自定义路径

默认安装在`/opt/TVHelper`中，自定义安装路径，将安装路径作为第二个参数添加，必须是绝对路径（如果路径以`TVHelper`结尾，则直接安装到给定路径，否则会安装在给定路径`TVHelper`目录下），如安装到`/root`：

```shell
# Install
curl -fsSL "https://gh-proxy.com/https://raw.githubusercontent.com/sec-an/TVHelper/main/v1.sh" | bash -s install /root
# update
curl -fsSL "https://gh-proxy.com/https://raw.githubusercontent.com/sec-an/TVHelper/main/v1.sh" | bash -s update /root
# Uninstall
curl -fsSL "https://gh-proxy.com/https://raw.githubusercontent.com/sec-an/TVHelper/main/v1.sh" | bash -s uninstall /root
```

- 启动：`systemctl start TVHelper`
- 关闭：`systemctl stop TVHelper`
- 状态：`systemctl status TVHelper`
- 重启：`systemctl restart TVHelper`

## 手动安装

### 获取TVHelper

打开[TVHelper Release](https://github.com/sec-an/TVHelper/releases)下载待部署系统对应的文件。

### 运行

```shell
# 解压下载的文件，得到可执行文件：
tar -zxvf TVHelper_*.tar.gz
# 授予程序执行权限：
chmod +x TVHelper
# 运行程序
./TVHelper
```
### 守护进程（Linux systemd）

使用任意方式编辑`/usr/lib/systemd/system/TVHelper.service`并添加如下内容，其中`path_TVHelper`为`TVHelper`所在的路径

```bash
[Unit]
Description=TVHelper service
Wants=network.target
After=network.target network.service
 
[Service]
Type=simple
WorkingDirectory=path_TVHelper
ExecStart=path_TVHelper/TVHelper
KillMode=process
Restart=always
 
[Install]
WantedBy=multi-user.target
```

然后，执行`systemctl daemon-reload`重载配置，现在你可以使用这些命令来管理程序：
- 启动：`systemctl start TVHelper`
- 关闭：`systemctl stop TVHelper`
- 配置开机自启：`systemctl enable TVHelper`
- 取消开机自启：`systemctl disable TVHelper`
- 状态：`systemctl status TVHelper`
- 重启：`systemctl restart TVHelper`

### 守护进程（OpenWrt等 procd）

使用任意方式编辑`/etc/init.d/TVHelper`并添加如下内容，其中`PROG`为`TVHelper`所在的路径

```bash
#!/bin/sh /etc/rc.common

START=99
STOP=98

USE_PROCD=1

DAEMON=TVHelper
PROG=/opt/TVHelper

start_service() {
    procd_open_instance TVHelper
    procd_set_param command $PROG/$DAEMON -d $PROG
    procd_set_param respawn
    procd_close_instance
    echo "service TVHelper start"
}

stop_service() {
    killall $DAEMON
    echo "service TVHelper stop"
}
```

然后，你可以使用这些命令来管理程序：
- 启动：`/etc/init.d/TVHelper start`
- 关闭：`/etc/init.d/TVHelper stop`
- 配置开机自启：`/etc/init.d/TVHelper enable`
- 取消开机自启：`/etc/init.d/TVHelper disable`
- 状态：`/etc/init.d/TVHelper status`

## 如何更新
在[TVHelper Release](https://github.com/sec-an/TVHelper/releases)下载新版TVHelper，替换之前的版本即可。

## 配置示例
[无注释模板](https://github.com/sec-an/TVHelper/blob/main/configs/config/default.json)

[配置样例](https://github.com/sec-an/TVHelper/blob/main/configs/config/sample.json)

```json5
// 请自行新建配置，本配置仅供参考，请勿修改，后期更新可能会覆盖！
// 请自行新建配置，本配置仅供参考，请勿修改，后期更新可能会覆盖！
// 请自行新建配置，本配置仅供参考，请勿修改，后期更新可能会覆盖！

// 最终配置地址为：http://你的ip:16214/config/文件名
// 本配置的地址为：http://你的ip:16214/config/sample

// 当前程序提供的豆瓣主页在：http://你的ip:16214/home?douban=你的豆瓣id
// 当前程序提供的直播文件服务器为：http://你的ip:16214/live/文件名.后缀
// 直播文件示例：http://你的ip:16214/live/IPTV.m3u
// 若订阅地址为本地文件，请在source_config目录下新建目录并放置在新建目录中
// 本地订阅：http://你的ip:16214/config/src/新建的文件夹名/文件名.后缀
// 本地订阅json示例：http://你的ip:16214/config/src/demo/config.json
// 本地订阅jar示例：http://你的ip:16214/config/src/demo/custom_spider.jar;md5;a84fef826cb82da525469e8acf1e7d9a"

{
  // 订阅地址列表
  "subscribe": [
    {
      // 订阅地址
      "url": "https://hutool.ml/tang",
      // multi-jar为false时，将采用第一个订阅地址中的spider或者本配置文件中新的spider
      "multi-jar": false,
      "always-on": false
    },
    {
      "url": "http://饭太硬.ga/x/o.json",
      // multi-jar为true时
      // 本订阅中所有原本未定义多jar的站点
      // 将同一设置多jar为本订阅的spider
      "multi-jar": true,
      // 若always-on为false，且该订阅地址之前存在有效订阅，则不展示该订阅
      "always-on": false
    },
    {
      // 本地配置文件放置在source_config目录中，且需要在该目录下新建子文件夹
      // 配置地址为：http://你的ip:16214/config/src/新建的文件夹名/文件名.后缀
      "url": "http://127.0.0.1:16214/config/src/demo/config.json",
      // multi-jar为true时
      // 本订阅中所有原本未定义多jar的站点
      // 将同一设置多jar为本订阅的spider
      "multi-jar": true,
      // 该字段会替换该订阅中的spider
      // 本配置地址为：http://你的ip:16214/config/src/新建的文件夹名/文件名.后缀
      "jar": "http://127.0.0.1:16214/config/src/demo/custom_spider.jar;md5;a84fef826cb82da525469e8acf1e7d9a",
      // 若always-on为true，无论是否已存在有效订阅，都将展示该订阅
      "always-on": true
    }
  ],

  // 直播源，替换订阅配置，非增加，仅支持一组
  // 如果不需要替换，则改为："lives": []
  "lives": [
    {
      "name": "直播",
      // txt为0，json为1(仅影视支持json)
      "type": 1,
      // 可以在live文件夹中添加本地直播文件
      // 格式为：http://你的ip:16214/live/文件名.后缀
      // 示例：http://你的ip:16214/live/IPTV.m3u
      "url": "https://gh-proxy.com/https://raw.githubusercontent.com/FongMi/CatVodSpider/main/json/live.json",
      "epg": "http://epg.51zmt.top:8000/api/diyp/?ch={epg}&date={date}",
      "logo": "http://epg.51zmt.top:8000/{logo}",
      // 是否自动开启(仅影视支持)
      "boot": true,
      // 播放器，1：IJK，2：EXO
      "playerType": 1
    }
  ],

  // 点播黑名单，名称(name)在该列表中的点播源将被移除
  // 示例为移除”哔哩“点播源
  // 如果不需要替换，则改为："sites-blacklist": []
  "sites-blacklist": ["哔哩"],

  // 点播源，数组合并至订阅配置sites数组 前
  // 如果不需要增加，则改为："sites-prepend": []
  "sites-prepend": [
    {
      "key": "T4-douban",
      "name": "豆瓣主页",
      "type": 4,
      // 本程序默认运行在当前设备16214端口且自带豆瓣主页服务
      // 该api可替换成http(s)://你的ip或域名(:16214)/home?douban=你的豆瓣id
      "api": "https://t4.secan.icu/vod?douban=你的豆瓣id",
      "searchable": 0,
      "filterable": 1
    }
  ],

  // 点播源，数组合并至订阅配置sites数组 后
  // 如果不需要增加，则改为："sites-append": []
  "sites-append": [
    {
      "key": "Live",
      "name": "直播",
      "type": 3,
      "api": "csp_Live",
      "searchable": 0,
      "filterable": 0,
      // 设定延迟多少毫秒后进入直播
      "ext": "2000",
      // 首屏直播需要使用FongMi的底包，可能需要使用多jar
      "jar": "https://gh-proxy.com/https://raw.githubusercontent.com/FongMi/CatVodSpider/main/jar/custom_spider.jar"
    },
    {
      "key": "易搜",
      "name": "易搜",
      "type": 3,
      "api": "csp_YiSo",
      "searchable": 1,
      "filterable": 0,
      // 换源，0：关闭，1：开启
      "switchable": 1,
      // 播放器，1：IJK，2：EXO
      "playerType": 2,
      "ext": "http://我不是.肥猫.love:63/token.php"
    }
  ],

  // 自定义爬虫，非空则替换订阅中的爬虫
  // 如果不需要替换，则改为："spider": ""
  "spider": "",

  // 壁纸，非空则替换订阅中的壁纸
  // 如果不需要替换，则改为："wallpaper": ""
  "wallpaper": "",

  // VIP解析标识，对象合并至订阅配置flags数组后
  // 如果不需要替换，则改为："mix-flags": []
  "mix-flags": [],

  // 解析地址，对象合并至订阅配置parses数组后
  // 如果不需要替换，则改为："mix-parses": []
  "mix-parses": [],

  // 解析广告过滤，对象合并至订阅配置ads数组后
  // 如果不需要替换，则改为："mix-ads": []
  "mix-ads": []
}
```

## 优先适配:thumbsup:
[:tv: FongMi / TV](https://github.com/FongMi/TV)
