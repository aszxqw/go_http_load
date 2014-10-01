package httpload

import (
    "os"
    //"flag"
    "io/ioutil"
    "net/http"
    "bufio"
    "io"
    "strings"
	"github.com/golang/glog"
)


func readContent(res *http.Response) ([]byte, error) {
    defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
    return contents, err
}

func getLinesFromFile(filename string) []string {
	var lines = make([]string, 0)
	file, err := os.Open(filename)
	if err != nil {
		glog.Error(err)
        return lines
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
	//glog.Infof("read total lines count %d", len(lines))
	return lines
}
