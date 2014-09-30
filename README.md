# go_http_load

## Usage

### Download & Install

```
go get github.com/aszxqw/go_http_load
```

### Show Usage

```
go_http_load
```

### Example

HTTP GET

```
go_http_load -method=GET -get_urls="urls" -alsologtostderr=true
```

HTTP POST

```
go_http_load -method=POST -post_url="http://127.0.0.1:80" -post_data_file=README.md  -alsologtostderr=true
```

