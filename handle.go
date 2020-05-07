package main

import (
	"errors"
	"log"
	"net"

	"github.com/miekg/dns"
	"github.com/xujiajun/nutsdb"
)

func TCPHandle(w dns.ResponseWriter, req *dns.Msg) {
	handle("tcp", w, req)
}
func UDPHandle(w dns.ResponseWriter, req *dns.Msg) {
	handle("udp", w, req)
}
func handle(nettype string, w dns.ResponseWriter, req *dns.Msg) {
	question := req.Question[0]
	ips, err := getIP(question.Name)
	if err != nil {
		dns.HandleFailed(w, req)
		return
	}
	rr_header := dns.RR_Header{
		Name:   question.Name,
		Rrtype: question.Qtype,
		Class:  dns.ClassINET,
	}
	for _, ip := range ips {
		a := &dns.A{rr_header, ip}
		req.Answer = append(req.Answer, a)
	}
	w.WriteMsg(req)
}
func getIP(name string) ([]net.IP, error) {
	name = name[0 : len(name)-1]
	err := db.View(func(tx *nutsdb.Tx) error {
		_, err := tx.Get("gamekiller", []byte(name))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		ips, err := net.LookupIP(name)
		if err != nil {
			return nil, err
		}
		return ips, nil
	}
	return nil, errors.New("balck hostname")
}
func setHostname(hostname string) bool {
	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			key := []byte(hostname)
			val := []byte(hostname)
			bucket := "gamekiller"
			if err := tx.Put(bucket, key, val, 0); err != nil {
				return err
			}
			return nil
		}); err != nil {
		log.Printf("%v\n", err)
		return false
	}
	return true
}
