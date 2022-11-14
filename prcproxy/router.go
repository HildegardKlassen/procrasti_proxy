package prcproxy

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (p *PrcProxy) CreateRouter() *mux.Router {
	r := mux.NewRouter()

	if p.BlockAllRequests {
		r.HandleFunc("/", p.blockAllHandler)
		return r
	}

	r.HandleFunc("/admin/{command}/{host}", p.adminHostConfigurationHandler)
	r.HandleFunc("/", p.blockRequestsFromList)
	return r
}

func (p *PrcProxy) blockAllHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "FORBIDDEN", http.StatusForbidden)
}

func (p *PrcProxy) routeAllRequestsHandler(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Write(body)
}
func (p *PrcProxy) blockRequestsFromList(w http.ResponseWriter, r *http.Request) {
	if p.timeIsInWindow(time.Now()) {
		if p.isHostInBlockList(r.URL.Host) {
			p.blockAllHandler(w, r)
		}
	}

	p.routeAllRequestsHandler(w, r)
}

func (p *PrcProxy) adminHostConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	adminMethod := mux.Vars(r)["command"]
	switch adminMethod {
	case "":
		http.Error(w, "no valid admin command", http.StatusBadRequest)
		return
	case "block":
		host := mux.Vars(r)["host"]
		if host == "" {
			http.Error(w, "no valid host to block", http.StatusBadRequest)
			return
		}
		if !p.isHostInBlockList(host) {
			p.BlockList = append(p.BlockList, host)
		}
		w.Write([]byte(fmt.Sprintf("the host %s was successfully blocked", host)))
		return
	case "unblock":
		host := mux.Vars(r)["host"]
		if host == "" {
			http.Error(w, "no valid host to unblock", http.StatusBadRequest)
			return
		}
		if p.isHostInBlockList(host) {
			p.removeHostFromBlockList(host)
		}
		w.Write([]byte(fmt.Sprintf("the host %s was successfully unblocked", host)))
	}
}

func (p *PrcProxy) removeHostFromBlockList(host string) {
	for i, v := range p.BlockList {
		if v == host {
			p.BlockList = append(p.BlockList[:i], p.BlockList[i+1:]...)
			break
		}
	}
}

func (p *PrcProxy) isHostInBlockList(host string) bool {
	for _, s := range p.BlockList {
		if s == host {
			return true
		}
	}
	return false
}

func (p *PrcProxy) timeIsInWindow(tnow time.Time) bool {
	currentTime := time.Date(int(0000), time.January, int(1), tnow.Hour(), tnow.Minute(), tnow.Second(), tnow.Nanosecond(), time.Now().Local().Location())
	if currentTime.After(p.BlockStartTime) && currentTime.Before(p.BlockEndTime) {
		return true
	}
	return false
}
