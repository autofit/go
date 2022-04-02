package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/json-iterator/go"
)

type Handle struct {
	Path string `json:"path,omitempty"`
	Host string `json:"host,omitempty"`
	Port string `json:"port,omitempty"`
}

type Iface struct {
	Pem       string   `json:"pem,omitempty"`
	Key       string   `json:"key,omitempty"`
	LocalPort string   `json:"localport,omitempty"`
	Remote    []Handle `json:"remote,omitempty"`
}

func main() {
	ReadInitFile()
	mux := http.NewServeMux()
	mux.HandleFunc("/", ServeHTTP)
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		for range ticker.C {
			ReadInitFile()
		}
	}()
	if If.Pem != "" && If.Key != "" {
		err := http.ListenAndServeTLS(If.LocalPort, If.Pem, If.Key, mux)
		if err != nil {
			log.Println(err)
		}
	} else {
		err := http.ListenAndServe(If.LocalPort, mux)
		if err != nil {
			log.Println(err)
		}
	}
}

var If Iface

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var target *url.URL
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println(r.RequestURI)
	log.Println(r.Host)
	log.Println(r.RemoteAddr)
	log.Println(r.Header.Get("Referer"), r.Header.Get("Origin"))
	i := 0
	Path := ""
	if r.Header.Get("Referer") != "" {
		path := strings.Split(strings.Split(r.Header.Get("Referer"), ".")[0], "//")[1]
		for _, v := range If.Remote {
			if r.RequestURI[1:] == v.Path || path == v.Path {
				Path = v.Host + ":" + v.Port
				i++
			}
		}
	}
	if i == 0 {
		if r.Header.Get("Origin") != "" {
			path := strings.Split(strings.Split(r.Header.Get("Origin"), ".")[0], "//")[1]
			for _, v := range If.Remote {
				if r.RequestURI[1:] == v.Path || path == v.Path {
					Path = v.Host + ":" + v.Port
					i++
				}
			}
		}
		if i == 0 {
			if r.Host != "" {
				path := strings.Split(r.Host, ".")[0]
				for _, v := range If.Remote {
					if r.RequestURI[1:] == v.Path || path == v.Path {
						Path = v.Host + ":" + v.Port
					}
				}
			}
		}
	}
	if Path[len(Path)-1:] == ":" {
		Path = Path[:len(Path)-1]
	}
	target, _ = url.Parse(Path)
	log.Println(target)
	if target == nil {
		return
	}
	if r.RequestURI == "/favicon.ico" {
		io.WriteString(w, "Request path Error")
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w, r)
}
func ReadInitFile() {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	file, _ := os.Open("./iface.ini")
	bytes, _ := ioutil.ReadAll(file)
	log.Println(string(bytes))
	if err := json.Unmarshal(bytes, &If); err != nil {
		log.Println(err)
	}
	log.Println(If)
}
