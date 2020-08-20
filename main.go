package main

import (
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/http/httptest"
	"os"
	"path"

	"github.com/markbates/pkger"
	"github.com/webview/webview"
)

func httpServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := path.Join("/frontend/public", r.URL.Path)
		file, err := pkger.Open(url)
		if err != nil && errors.Is(err, os.ErrNotExist) {
			http.NotFound(w, r)
		} else if err != nil {
			panic(err)
		}
		defer file.Close()

		w.Header().Set("Content-Type", mime.TypeByExtension(url))
		_, err = io.Copy(w, file)
		if err != nil {
			panic(err)
		}
	}))
}

func main() {
	pkger.Include("/frontend/public")

	srv := httpServer()
	defer srv.Close()

	url := fmt.Sprintf("%s/index.html", srv.URL)
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(url)
	w.Run()
}
