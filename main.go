package main

import (
	//"time"
	"net/http"
	//"reflect"
	"bufio"
    "strings"
	"flag"
	"io"
	"os"
	"sync"
	//"fmt"
    "io/ioutil"
	"github.com/golang/glog"
)

var wait_group sync.WaitGroup

var flag_method *string = flag.String("method", "", "GET or POST")

var flag_get_urls_file *string = flag.String("get_urls", "", "the urls filename for GET requests")
var flag_post_host *string = flag.String("post_url", "", "the url for POST requests. for example: http://127.0.0.1:80/")
var flag_post_data_file *string = flag.String("post_data_file", "", "the data filename for POST")
var flag_coroutine_number *int  = flag.Int("coroutine_number", 0, "the coroutine_number for every http request")
var flag_loop_count *int = flag.Int("loop_count", 0, "the count for looping")
var flag_verbose *bool = flag.Bool("verbose", false, "verbose , if verbose is true, printing more details")
var flag_post_body_type *string = flag.String("post_body_type", "text/plain", "")

func exitAfterUsage() {
	flag.Usage()
	os.Exit(1)
}

func readContent(res *http.Response) ([]byte, error) {
    defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
    return contents, err
}

func runHttpGet(url string){
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

func runHttpPost(url string, content_type string, filename string) {
	defer wait_group.Done()

	file, err := os.Open(filename)
	if err != nil {
		glog.Error(err)
		exitAfterUsage()
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

func getLinesFromFile(filename string) []string {
	var lines = make([]string, 0)
	file, err := os.Open(filename)
	if err != nil {
		glog.Error(err)
		exitAfterUsage()
	}
	defer file.Close()

	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		lines = append(lines, strings.TrimRight(line, "\n"))
	}
	glog.Infof("read lines[%d]", len(lines))
	return lines
}

func main() {
	flag.Parse()
	defer wait_group.Wait()
	if *flag_method == "" {
		glog.Error("you should set option -method ")
		exitAfterUsage()
	}
	if *flag_method == "GET" {
		urls := getLinesFromFile(*flag_get_urls_file)
		for _, url := range urls {
			glog.Error(url)
			wait_group.Add(1)
			go runHttpGet(url)
		}
	} else if *flag_method == "POST" {
        go runHttpPost(*flag_post_host, *flag_post_body_type, *flag_post_data_file)
	}
}
