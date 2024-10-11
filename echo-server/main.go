package main

import (
	"net/http"
	"os"
	"server/handler"
	"server/option"
)

func main() {
	opt := option.GetOption()
	os.Stdout.WriteString("Start listening on " + opt.Listen + "\n")

	h := http.HandlerFunc(handler.Handler)
	if len(opt.Cert) > 0 && len(opt.Key) > 0 {
		http.ListenAndServeTLS(opt.Listen, opt.Cert, opt.Key, h)
	} else {
		http.ListenAndServe(opt.Listen, h)
	}
}
