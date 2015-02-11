# go_http_load

[![Build Status](https://travis-ci.org/yanyiwu/go_http_load.svg?branch=master)](https://travis-ci.org/yanyiwu/go_http_load)
[![GoDoc](https://godoc.org/github.com/yanyiwu/go_http_load?status.svg)](https://godoc.org/github.com/yanyiwu/go_http_load)
[![RTD](https://readthedocs.org/projects/go-http-load/badge/?version=latest)](http://go-http-load.readthedocs.org/en/latest/)

## Introduction

http load testing tool (using golang)。

## Usage

```
go get github.com/yanyiwu/go_http_load
```

### Show Usage

```
go_http_load
```

### Example

HTTP GET

```
go_http_load -method=GET -get_urls="urls"
go_http_load -method=GET -get_urls="urls" -loop_count=100 -goroutines=2
```

HTTP POST

```
go_http_load -method=POST -post_url="http://127.0.0.1:11267" -post_data_file=README.md
go_http_load -method=POST -post_url="http://127.0.0.1:11267" -post_data_file=README.md -loop_count=100 -goroutines=2
```

## Contact

`wuyanyi09@foxmail.com`
