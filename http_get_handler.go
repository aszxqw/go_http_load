package main

import (
	"github.com/golang/glog"
    "net/http"
)

func httpGetHandler() {
    var get_urls_chan chan string = make(chan string , *flag_coroutine_number)
    glog.Infof("urls_chan buffer size %d", *flag_coroutine_number)
    for i := 0; i < *flag_coroutine_number; i++ {
        go httpGetChanWorker(get_urls_chan)
    }
    urls := getLinesFromFile(*flag_get_urls_file)
    for _, url := range urls {
        for i := 0; i < *flag_loop_count; i++ {
            glog.Info(url)
            wait_group.Add(1)
            get_urls_chan<- url
        }
    }
}

func httpGetChanWorker(urls_chan <-chan string) {
    for url := range urls_chan {
        httpGetWorker(url)
    }
}

func httpGetWorker(url string) {
	defer wait_group.Done()
	res, err := http.Get(url)
	if err != nil {
		glog.Error(err)
		return
	}
    content, err:= readContent(res)
    if err != nil {
        glog.Error(err)
        return
    }
	glog.Infof("%s response body size %d", res.Status, len(content))
    if *flag_verbose {
        glog.Info(string(content))
    }
}
