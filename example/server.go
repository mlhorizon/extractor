package main

import (
	"encoding/json"
	"github.com/yqingp/extractor"
	"net/http"
)

type RenderContent struct {
	Url       string `json:"url"`
	Content   string `json:"content"`
	ErrorInfo string `json:"error"`
}

func work(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	extract_worker := extractor.NewExtractor(url)
	content, err := extract_worker.Extract()
	var renderContent RenderContent

	if err != nil {
		renderContent.Url = url
		renderContent.ErrorInfo = err.Error()
	} else {
		renderContent.Url = url
		renderContent.Content = content
	}

	w.WriteHeader(http.StatusOK)
	js, _ := json.Marshal(renderContent)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(js)
}

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	http.HandleFunc("/work", work)
	http.ListenAndServe(":8000", nil)
}
