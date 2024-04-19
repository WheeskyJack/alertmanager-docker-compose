package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Custom handler struct
type myHandler struct {
	url  string
	name string
}

// ServeHTTP method for myHandler
func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, err)
		w.WriteHeader(500)
	} else {
		fmt.Println(h.name, h.url, string(body))
		w.WriteHeader(200)
	}
}

type cfg struct {
	r *http.ServeMux
	h *myHandler
}

func main() {
	url1 := ":8886"
	url2 := ":8887"
	url3 := ":8888"
	url4 := ":8889"

	cfg1 := getServerCfg(url1, "webhook_team1")
	cfg2 := getServerCfg(url2, "webhook_vector")
	cfg3 := getServerCfg(url3, "webhook_kafka")
	cfg4 := getServerCfg(url4, "webhook_critical_recv")

	go func() { log.Fatal(http.ListenAndServe(cfg1.h.url, cfg1.r)) }()
	go func() { log.Fatal(http.ListenAndServe(cfg2.h.url, cfg2.r)) }()
	go func() { log.Fatal(http.ListenAndServe(cfg3.h.url, cfg3.r)) }()
	go func() { log.Fatal(http.ListenAndServe(cfg4.h.url, cfg4.r)) }()

	select {}
}

func getServerCfg(url string, name string) cfg {
	mh := &myHandler{url: url, name: name}
	mr := http.NewServeMux()
	mr.Handle("/", mh)
	return cfg{r: mr, h: mh}
}
