package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"libvirt.org/go/libvirt"
)

func domains(w http.ResponseWriter, r *http.Request) {
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

func startDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	domainName := vars["name"]
	dom, err := libConn.LookupDomainByName(domainName)
	if err != nil {
		dom.Free()
		json.NewEncoder(w).Encode(false)
	}

	dom.Create()
	dom.Free()
	json.NewEncoder(w).Encode(true)
}

func stopDomain(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	domainName := vars["name"][0]
	dom, err := libConn.LookupDomainByName(domainName)
	if err != nil {
		json.NewEncoder(w).Encode(false)
	} else {
		dom.Destroy()
		dom.Free()
		json.NewEncoder(w).Encode(true)
	}

}
