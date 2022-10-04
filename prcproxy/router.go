package prcproxy

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func (p *PrcProxy) CreateRouter() *http.ServeMux {
	r := http.NewServeMux()
	if p.BlockRequests {
		r.HandleFunc("/", p.blockAllHandler)
	}

	if p.ProxyMakesRequests {
		r.HandleFunc("/", p.routeAllRequestsHandler)
	}

	return r
}

func (p *PrcProxy) blockAllHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "FORBIDDEN", http.StatusForbidden)
}

func (p *PrcProxy) routeAllRequestsHandler(w http.ResponseWriter, r *http.Request) {
	client := http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", res.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", res.Header.Get("Content-Length"))
	io.Copy(w, res.Body)
	res.Body.Close()
}
