package main

import (
    "net/http"
    "os"
    "bufio"
    "github.com/golang/glog"
)

type postArgs struct {
    Url string
    BodyType string
    Filename string
}

func httpPostHandler() {
    if *flag_post_url == "" {
        glog.Error("post_url must not be empty.")
        exitAfterUsage()
    }
    if *flag_post_data_file == "" {
        glog.Error("post_data_file must not be empty.")
        exitAfterUsage()
    }
    wait_group.Add(1)
    var args postArgs = postArgs{*flag_post_url, *flag_post_body_type, *flag_post_data_file}
    go httpPostWorker(args)
}

func httpPostChanWorker(post_chan <-chan postArgs) {
    for item := range post_chan {
        httpPostWorker(item)
    }
}

func httpPostWorker(args postArgs) {
    var url string = args.Url
    var content_type string = args.BodyType
    var filename string = args.Filename

	defer wait_group.Done()

	file, err := os.Open(filename)
	if err != nil {
		glog.Error(err)
        return
	}
	defer file.Close()

	rd := bufio.NewReader(file)
	res, err := http.Post(url, content_type, rd)
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

