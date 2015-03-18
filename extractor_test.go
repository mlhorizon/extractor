package extractor

import "testing"
import "fmt"

func TestParse(t *testing.T) {
	temp := NewExtractor("http://china.huanqiu.com/article/2015-03/5934483.html")
	a, _ := temp.Extract()
	fmt.Println(a)
}
