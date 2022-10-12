package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"libvirt.org/go/libvirt"
)

type DomainAPIReturnType struct {
	Name    string
	Running libvirt.DomainState
}

var libConn *libvirt.Connect

func main() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		panic(err)
	}

	libConn = conn

	defer conn.Close()

	r := mux.NewRouter()
	r.HandleFunc("/domains", domains)

	r.HandleFunc("/startDomain", startDomain)

	r.HandleFunc("/stopDomain", stopDomain)

	log.Fatal(http.ListenAndServe(":9090", r))
}
