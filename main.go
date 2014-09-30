package main

import (
	//"time"
	//"reflect"
	"flag"
	"sync"
	"github.com/golang/glog"
	//"fmt"
)

var wait_group sync.WaitGroup

var flag_method *string = flag.String("method", "", "GET or POST")

var flag_get_urls_file *string = flag.String("get_urls", "", "the urls filename for GET requests")
var flag_post_url *string = flag.String("post_url", "", "the url for POST requests. for example: http://127.0.0.1:80/")
var flag_post_data_file *string = flag.String("post_data_file", "", "the data filename for POST")
var flag_coroutine_number *int  = flag.Int("coroutine_number", 1, "the coroutine_number for every http request")
var flag_loop_count *int = flag.Int("loop_count", 1, "the count for looping")
var flag_verbose *bool = flag.Bool("verbose", false, "verbose , if verbose is true, printing more details")
var flag_post_body_type *string = flag.String("post_body_type", "text/plain", "")

func main() {
	flag.Parse()
    if *flag_loop_count <= 0 {
        glog.Error("loop_count should be larger than zero.")
        exitAfterUsage()
    }
    if *flag_coroutine_number <= 0 {
        glog.Error("coroutine_number should be larger than zero.")
        exitAfterUsage()
    }
	defer wait_group.Wait()
	if *flag_method == "" {
		glog.Error("you should set option -method ")
		exitAfterUsage()
	}
	if *flag_method == "GET" {
        httpGetHandler()
	} else if *flag_method == "POST" {
        httpPostHandler()
	}
}
