//See https://github.com/yanyiwu/go_http_load/blob/master/README.md
package main

import (
	//"time"
	//"reflect"
	"flag"
	"github.com/golang/glog"
	"github.com/yanyiwu/go_http_load/httpload"
	"os"
	//"fmt"
)

var flag_method *string = flag.String("method", "", "GET or POST")

var flag_verbose *bool = flag.Bool("verbose", false, "verbose , if verbose is true, printing more details")

func exitAfterUsage() {
	flag.Usage()
	os.Exit(1)
}

func main() {
	flag.Parse()
	//f := flag.Lookup("alsologtostderr")
	//if f.DefValue == f.Value.String() {
	//	flag.Set("alsologtostderr", "true")
	//}
	if *flag_method == "" {
		glog.Error("you should set option -method ")
		exitAfterUsage()
	}
	var handler httpload.HandlerInterface
	switch *flag_method {
	case "GET":
		handler = httpload.NewGetHandler()
	case "POST":
		handler = httpload.NewPostHandler()
	default:
		glog.Error("-method is illegal.")
		exitAfterUsage()
	}
	handler.Run()
}
