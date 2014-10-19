package httpload

import (
    "flag"
)

var flag_goroutines *int  = flag.Int("goroutines", 1, "the goroutines for every http request")
var flag_loop_count *int = flag.Int("loop_count", 1, "the count for looping")
var flag_get_urls_file *string = flag.String("get_urls", "", "the urls filename for GET requests")
var flag_post_url *string = flag.String("post_url", "", "the url for POST requests. for example: http://127.0.0.1:80/")
var flag_post_data_file *string = flag.String("post_data_file", "", "the data filename for POST")
var flag_post_body_type *string = flag.String("post_body_type", "text/plain", "")
