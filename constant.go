package extractor

import "errors"

var (
	REPLACE_REGX = []string{`<!DOCTYPE.*?>`,
		`<!--[\s\S]*?-->`,
		`<script[^>]*>[\s\S]*?</script>`,
		`<style[^>]*>[\s\S]*?</style>`,
		`<.*?>`,
		`&.{1,5};|&#.{1,5};`,
	}
	SPLITERS = []string{"\r\n", "\r", "\n"}
)

const (
	DEFAULT_PAGE_ENCODING = "UTF-8"
	BLANK_REGX            = `\s+`
	NEW_SPLITER           = "\n"
	META_REGX             = `<meta.*?charset=([^"']+)`
	CHARSET_REGX          = `charset=([\w\d-]+);?`
	URL_REGEX             = `^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`
)

var (
	ERROR_URL_IS_UNMATCH   = errors.New("URL不符合要求")
	ERROR_URL_REQUEST_FAIL = errors.New("URL请求失败")
	ERROR_ICONV_ERROR      = errors.New("Iconv未知错误")
	ERROR_UNKONWN_CHARSET  = errors.New("未知网页编码")
)
