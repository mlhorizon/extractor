package extractor

import (
	"regexp"
	"strings"
)

var charsets = map[string]string{
	"UTF-8":  "UTF-8",
	"UTF8":   "UTF-8",
	"GBK":    "GB18030",
	"GB2312": "GB18030",
}

func guess(body string) (string, error) {
	regex, _ := regexp.Compile(META_REGX)
	tmp := regex.FindString(body)

	if tmp == "" {
		return "", ERROR_UNKONWN_CHARSET
	}

	regex, _ = regexp.Compile(CHARSET_REGX)
	tmp = regex.FindString(tmp)

	if tmp == "" {
		return "", ERROR_UNKONWN_CHARSET
	}

	foundCharset := strings.ToUpper(strings.Replace(tmp, "charset=", "", -1))

	defaultChartset := charsets[foundCharset]

	if defaultChartset == "" {
		return "", ERROR_UNKONWN_CHARSET
	}

	return defaultChartset, nil
}
