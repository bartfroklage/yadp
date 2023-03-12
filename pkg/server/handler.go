package server

import (
	"fmt"

	"github.com/miekg/dns"
)

type Handler struct {
}

func NewHandler() *Handler {
	dnsHandler := Handler{}
	return &dnsHandler
}

func (h *Handler) DoTCP(w dns.ResponseWriter, req *dns.Msg) {
	fmt.Println("TCP INCOMING")
}

func (h *Handler) DoUDP(w dns.ResponseWriter, req *dns.Msg) {

	fmt.Println("UDP INCOMING" + w.RemoteAddr().String())
}
