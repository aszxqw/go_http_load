# go_http_load

[![Build Status](https://travis-ci.org/aszxqw/go_http_load.svg?branch=master)](https://travis-ci.org/aszxqw/go_http_load)
[![GoDoc](https://godoc.org/github.com/aszxqw/go_http_load?status.svg)](https://godoc.org/github.com/aszxqw/go_http_load)

## 简介

HTTP服务的压力测试工具(using golang)。

## 用法

### 下载和安装

```
go get github.com/aszxqw/go_http_load
```

### 显示用法

```
go_http_load
```

### 示例

HTTP GET

```
go_http_load -method=GET -get_urls="urls"
go_http_load -method=GET -get_urls="urls" -loop_count=4 -lcoroutine_number=2
```

HTTP POST

```
go_http_load -method=POST -post_url="http://127.0.0.1:11267" -post_data_file=README.md
```

