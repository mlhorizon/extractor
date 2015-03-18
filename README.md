# extractor
中文网页正文内容提取
基于[《基于行块分布函数的通用网页正文抽取算法》](http://cx-extractor.googlecode.com/files/%E5%9F%BA%E4%BA%8E%E8%A1%8C%E5%9D%97%E5%88%86%E5%B8%83%E5%87%BD%E6%95%B0%E7%9A%84%E9%80%9A%E7%94%A8%E7%BD%91%E9%A1%B5%E6%AD%A3%E6%96%87%E6%8A%BD%E5%8F%96%E7%AE%97%E6%B3%95.pdf)实现

## 安装
```
	go get github.com/yqingp/extractor
```

## 使用
```go
	import (
		"github.com/yqingp/extractor"
	)
	....

	extract_worker := extractor.NewExtractor(url)
	content, err := extract_worker.Extract()
	
	if err != nil {
		fmt.Println(content)
	}
```	
## server方式启动
```
	go run  example/server.go
```	

```ruby
	require 'rest_client'
	RestClient.post("http://localhost:8000/work", {:url => "http://www.baidu.com"})
```

