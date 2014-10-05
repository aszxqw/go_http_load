//See https://github.com/aszxqw/go_http_load/blob/master/README.md
package main

import (
	//"time"
	//"reflect"
	"flag"
	"github.com/golang/glog"
    "github.com/aszxqw/go_http_load/httpload"
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
    f := flag.Lookup("alsologtostderr")
    if f.DefValue == f.Value.String() {
        flag.Set("alsologtostderr", "true")
    }
	if *flag_method == "" {
		glog.Error("you should set option -method ")
		exitAfterUsage()
	}
	if *flag_method == "GET" {
        httpload.NewGetHandler().Run()
	} else if *flag_method == "POST" {
        httpload.NewPostHandler().Run()
	}
}
