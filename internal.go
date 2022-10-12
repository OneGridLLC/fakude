package main

import (
	"libvirt.org/go/libvirt"
)

func startDomain(domainName string) (*libvirt.Domain, error) {
	dom, err := libConn.LookupDomainByName(domainName)
	if err != nil {
		dom.Free()
		return dom, err
	}

	dom.Create()
	dom.Free()
	return dom, nil
}

func stopDomain(domainName string) (*libvirt.Domain, error) {
	dom, err := libConn.LookupDomainByName(domainName)
	if err != nil {
		return dom, err
	}

	dom.Destroy()
	dom.Free()
	return dom, nil
}
