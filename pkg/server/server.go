package server

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/miekg/dns"
)

type Server struct {
	Port           int
	Host           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	controlChannel chan int
}

func (s *Server) Addr() string {
	return net.JoinHostPort(s.Host, strconv.Itoa(s.Port))
}

func (s *Server) Run() {

	s.controlChannel = make(chan int)
	handler := NewHandler()

	tcpHandler := dns.NewServeMux()
	tcpHandler.HandleFunc(".", handler.DoTCP)

	udpHandler := dns.NewServeMux()
	udpHandler.HandleFunc(".", handler.DoUDP)

	tcpServer := &dns.Server{Addr: s.Addr(),
		Net:          "tcp",
		Handler:      tcpHandler,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout}

	udpServer := &dns.Server{Addr: s.Addr(),
		Net:          "udp",
		Handler:      udpHandler,
		UDPSize:      65535,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout}

	go s.start(udpServer)
	go s.start(tcpServer)

	http.HandleFunc("/control", s.controlHandler)
	go http.ListenAndServe(":8090", nil)

	for {
		c := <-s.controlChannel
		if c == 10 {
			return
		}
	}
}

func (s *Server) start(ds *dns.Server) {
	err := ds.ListenAndServe()
	if err != nil {
		fmt.Printf("Start %s listener on %s failed:%s\n", ds.Net, s.Addr(), err.Error())
	}
}

func (s *Server) controlHandler(w http.ResponseWriter, req *http.Request) {
	s.controlChannel <- 10
}
