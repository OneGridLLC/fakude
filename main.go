package main

import (
	"log"
	"net/http"

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

	http.HandleFunc("/", baseHandler) // ping

	http.HandleFunc("/auth", authHandler) // login with RBAC

	http.HandleFunc("/domains", domainsHandler) // get info on active domains

	http.HandleFunc("/manage", manageHandler) // start, stop, create, reinstall, delete

	http.HandleFunc("/image", imageHandler) // upload VM images to S3 from template or ISO

	http.HandleFunc("/console", consoleHandler) // VNC endpoint

	log.Fatal(http.ListenAndServe(":9090", nil))
}
