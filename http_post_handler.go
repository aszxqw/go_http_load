package main

import (
    "net/http"
    //"bufio"
    "strings"
    "io/ioutil"
    "github.com/golang/glog"
    //"fmt"
)

type postArgs struct {
    Url string
    BodyContent string
}

func httpPostHandler() {
    var post_data_chan chan postArgs = make(chan postArgs, *flag_coroutine_number)
    glog.Infof("post_data_chan size %d", *flag_coroutine_number)
    for  i:= 0; i < *flag_coroutine_number; i++ {
        go httpPostChanWorker(post_data_chan)
    }
    if *flag_post_url == "" {
        glog.Error("post_url must not be empty.")
        exitAfterUsage()
    }
    if *flag_post_data_file == "" {
        glog.Error("post_data_file must not be empty.")
        exitAfterUsage()
    }
    body_content, err  := ioutil.ReadFile(*flag_post_data_file)
    if err != nil {
        glog.Error(err)
        exitAfterUsage()
    }
    var args postArgs = postArgs {
        *flag_post_url,
        string(body_content),
    }
    for i := 0; i < *flag_loop_count; i++ {
        glog.Info(*flag_post_url)
        wait_group.Add(1)
        post_data_chan <- args
    }
}

func httpPostChanWorker(post_chan <-chan postArgs) {
    for item := range post_chan {
        httpPostWorker(item)
    }
}

func httpPostWorker(args postArgs) {
	defer wait_group.Done()

    var url string = args.Url
    var body_content string = args.BodyContent

    rd := strings.NewReader(body_content)
	res, err := http.Post(url, *flag_post_body_type, rd)
	if err != nil {
		glog.Error(err)
		return
	}
    content, err:= readContent(res)
    if err != nil {
        glog.Error(err)
        return
    }
	glog.Info(string(content))
}

