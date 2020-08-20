package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/markbates/pkger"
	"github.com/webview/webview"
)

type pkgerServer struct{}

func (p *pkgerServer) Open(name string) (http.File, error) {
	return pkger.Open(name)
}

func httpServer() *httptest.Server {
	return httptest.NewServer(
		http.FileServer(
			&pkgerServer{}))
}

func main() {
	pkger.Include("/frontend/public")

	srv := httpServer()
	defer srv.Close()

	url := fmt.Sprintf("%s/frontend/public/index.html", srv.URL)
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(url)
	w.Run()
}
