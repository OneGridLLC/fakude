package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"libvirt.org/go/libvirt"
)

func baseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		return
	}

	w.Header().Set("content-type", "text/html")
	w.WriteHeader(200)
	_, err := fmt.Fprint(w, "online")
	if err != nil {
		fmt.Println("Error with request:", r)
		fmt.Println(err)

		return
	}
}

func authHandler(w http.ResponseWriter, r *http.Request) {

}

func domainsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		return
	}

	doms, err := libConn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		fmt.Println(err)
	}

	var returnObject []DomainAPIReturnType

	for _, dom := range doms {
		name, err := dom.GetName()
		if err == nil {
			domainState, returnInt, err := dom.GetState()

			_ = returnInt

			if err == nil {
				valueToInsert := DomainAPIReturnType{Name: name, Running: domainState}
				returnObject = append(returnObject, valueToInsert)
			}
		}
		dom.Free()
	}

	json.NewEncoder(w).Encode(returnObject)
}

func manageHandler(w http.ResponseWriter, r *http.Request) {

}

func imageHandler(w http.ResponseWriter, r *http.Request) {

}

func consoleHandler(w http.ResponseWriter, r *http.Request) {

}
