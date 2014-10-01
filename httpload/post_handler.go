package httpload

import (
    "net/http"
    //"bufio"
    "strings"
    "io/ioutil"
    "github.com/golang/glog"
    "sync"
    "fmt"
    "time"
)


type PostHandler struct {
    waitGroup sync.WaitGroup
    dataChan chan postArgs
}

func NewPostHandler() *PostHandler {
    this := new(PostHandler)
    this.dataChan = make(chan postArgs, *flag_coroutine_number)
    for  i:= 0; i < *flag_coroutine_number; i++ {
        go this.chanWork()
    }
    glog.V(2).Infof("dataChan buffer size %d", *flag_coroutine_number)
    return this
}

type postArgs struct {
    Url string
    BodyContent string
}

func (this *PostHandler) Run() {
    if *flag_post_url == "" {
        glog.Error("post_url must not be empty.")
        return
    }
    if *flag_post_data_file == "" {
        glog.Error("post_data_file must not be empty.")
        return
    }
    body_content, err  := ioutil.ReadFile(*flag_post_data_file)
    if err != nil {
        glog.Error(err)
        return
    }
    var args postArgs = postArgs {
        *flag_post_url,
        string(body_content),
    }
    sum := 0
    start := time.Now()
    for i := 0; i < *flag_loop_count; i++ {
        glog.V(2).Info(*flag_post_url)
        this.waitGroup.Add(1)
        sum ++
        this.dataChan <- args
    }
	this.waitGroup.Wait()
    fmt.Printf("The Number of Queries:%d\n", sum)
    consumed := time.Now().Sub(start)
    fmt.Printf("The Time Consumed: %.3f s\n", consumed.Seconds())
    fmt.Printf("Query Per Second: %.3f q/s\n", float64(sum) / consumed.Seconds() )
}

func (this *PostHandler) chanWork() {
    for data := range this.dataChan {
        this.query(data)
        this.waitGroup.Done()
    }
}

func (this *PostHandler) query(args postArgs) {
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
	glog.Infof("%s response body size %d", res.Status, len(content))
    glog.V(1).Info(string(content))
}

