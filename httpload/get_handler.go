package httpload

import (
	"github.com/golang/glog"
    "net/http"
    "sync"
    "fmt"
    "time"
)


type GetHandler struct {
    waitGroup sync.WaitGroup
    urlChan chan string
}

func NewGetHandler() *GetHandler{
    this := new(GetHandler)
    this.urlChan = make(chan string, *flag_coroutine_number)
    for i := 0; i < *flag_coroutine_number; i++ {
        go this.chanWork()
    }
    glog.V(2).Infof("urlChan buffer size %d", *flag_coroutine_number)
    return this
}

func (this *GetHandler) Run() {
    urls := getLinesFromFile(*flag_get_urls_file)
    sum := 0
    start := time.Now()
    for _, url := range urls {
        for i := 0; i < *flag_loop_count; i++ {
            glog.V(2).Info(url)
            this.waitGroup.Add(1)
            sum ++
            this.urlChan <- url
        }
    }
    this.waitGroup.Wait()
    fmt.Printf("The Number of Queries:%d\n", sum)
    consumed := time.Now().Sub(start)
    fmt.Printf("The Time Consumed: %.3f s\n", consumed.Seconds())
    fmt.Printf("Query Per Second: %.3f q/s\n", float64(sum) / consumed.Seconds() )
}

func (this *GetHandler) chanWork() {
    for url := range this.urlChan {
        this.query(url)
        this.waitGroup.Done()
    }
}

func (this *GetHandler) query(url string) {
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
    glog.V(1).Info(string(content))
}
