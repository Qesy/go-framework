package main

import (
	"net/http"
	"todo/controllers"

	"github.com/Qesy/qesygo"
)

// MyMux 定义数据类型
type MyMux struct{}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":12345", mux)
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	entry := controllers.Entry{W: w, R: r, Controller: "index", Action: "index", Params: []string{}, URL: qesygo.Substr(r.URL.Path, 1, 0)}
	entry.Run()

}
