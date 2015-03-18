package extractor

import (
	"bytes"
	"github.com/qiniu/iconv"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type Extractor struct {
	url          string
	body         string
	pageEncoding string
	bodyLines    []string
	blockSize    int
	blockLens    []int
	textLines    []string
}

func NewExtractor(url string) *Extractor {
	return &Extractor{
		url:       url,
		blockSize: 3,
	}
}

func validateParams(url string) error {
	isMatch, _ := regexp.MatchString(URL_REGEX, url)
	if !isMatch {
		return ERROR_URL_IS_UNMATCH
	}

	return nil
}

func (this *Extractor) preProcess() {
	for _, v := range REPLACE_REGX {
		tmpRegex, _ := regexp.Compile(v)
		this.body = tmpRegex.ReplaceAllString(this.body, "")
	}
}

func (this *Extractor) formatBodyEncoding() error {
	if this.pageEncoding == DEFAULT_PAGE_ENCODING {
		return nil
	}

	cd, err := iconv.Open(DEFAULT_PAGE_ENCODING, this.pageEncoding)

	defer cd.Close()

	if err != nil {
		return ERROR_ICONV_ERROR
	}

	this.body = strings.ToLower(cd.ConvString(this.body))

	return nil
}

func (this *Extractor) guessEncoding() error {
	response, err := http.Get(this.url)

	if err != nil {
		return ERROR_URL_REQUEST_FAIL
	}

	body, _ := ioutil.ReadAll(response.Body)

	defer response.Body.Close()

	defaultChartset, _ := guess(string(body))

	this.body = string(body)
	this.pageEncoding = defaultChartset

	return nil
}

func (this *Extractor) bodyToLines() {
	for _, v := range SPLITERS {
		this.body = strings.Replace(this.body, v, NEW_SPLITER, -1)
	}

	lines := strings.Split(this.body, NEW_SPLITER)
	regex, _ := regexp.Compile(`\s+`)

	for _, v := range lines {
		this.textLines = append(this.textLines, regex.ReplaceAllString(v, ""))
	}
}

func (this *Extractor) calBlocksLen() {
	textLineNum := len(this.textLines)
	blockLen := 0
	for i := 0; i < this.blockSize; i++ {
		blockLen += len(this.textLines[i])
	}

	this.blockLens = append(this.blockLens, blockLen)

	for i := 1; i < (textLineNum - this.blockSize); i++ {
		blockLen = this.blockLens[i-1] + len(this.textLines[i-1+this.blockSize]) - len(this.textLines[i-1])
		this.blockLens = append(this.blockLens, blockLen)
	}

}

func (this *Extractor) Extract() (string, error) {
	if err := validateParams(this.url); err != nil {
		return "", err
	}

	if err := this.guessEncoding(); err != nil {
		return "", err
	}

	if err := this.formatBodyEncoding(); err != nil {
		return "", err
	}

	this.preProcess()
	this.bodyToLines()
	this.calBlocksLen()

	start, end := -1, -1
	i := 0
	maxTextLen := 0
	blockNum := len(this.blockLens)

	for i < blockNum {
		for i < blockNum && (this.blockLens[i] == 0) {
			i++
		}

		if i >= blockNum {
			break
		}

		tmp := i
		curTextLen := 0
		var buffer bytes.Buffer

		for i < blockNum && (this.blockLens[i] != 0) {
			if this.textLines[i] != "" {
				buffer.WriteString(this.textLines[i])
				buffer.WriteString("<br />")
				curTextLen += len(this.textLines[i])
			}

			i++
		}

		if curTextLen > maxTextLen {
			this.body = buffer.String()
			maxTextLen = curTextLen
			start = tmp
			end = i - 1
		}

		_ = start
		_ = end
	}

	return this.body, nil
}
