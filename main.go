package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/miekg/dns"
	"github.com/xujiajun/nutsdb"
)

var opt nutsdb.Options
var db *nutsdb.DB

func main() {
	if err := InitConfig(); err != nil {
		log.Printf("Read config error:%v\n", err)
		os.Exit(-1)
	}
	var err error
	opt = nutsdb.DefaultOptions
	opt.Dir = Config.CacheFile //这边数据库会自动创建这个目录文件
	db, err = nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	go InitWebServer()
	InitDNS()
}
func InitWebServer() {
	http.HandleFunc("/", index)
	fmt.Printf("Web Console Listen and Serve %s:%d\n", Config.WebAddress, Config.WebPort)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", Config.WebAddress, Config.WebPort), nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
func InitDNS() {
	tcpHandler := dns.NewServeMux()
	tcpHandler.HandleFunc(".", TCPHandle)
	udpHandler := dns.NewServeMux()
	udpHandler.HandleFunc(".", UDPHandle)
	tcpServer := &dns.Server{
		Addr:    fmt.Sprintf("%s:%d", Config.Address, Config.Port),
		Net:     "tcp",
		Handler: tcpHandler,
	}
	udpServer := &dns.Server{
		Addr:    fmt.Sprintf("%s:%d", Config.Address, Config.Port),
		Net:     "udp",
		Handler: udpHandler,
	}
	go tcpServer.ListenAndServe()
	udpServer.ListenAndServe()
}
