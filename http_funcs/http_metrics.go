package http_funcs

import(
	"net/http/httptrace"
	"time"
	"crypto/tls"
	"log"
)

var start, connect, dns, tlsHandshake time.Time

func GetMetricsObject() *httptrace.ClientTrace{
	return &httptrace.ClientTrace{
        DNSStart: func(dsi httptrace.DNSStartInfo) { dns = time.Now() },
        DNSDone: func(ddi httptrace.DNSDoneInfo) {
            log.Printf("DNS Done: %v\n", time.Since(dns))
        },

        TLSHandshakeStart: func() { tlsHandshake = time.Now() },
        TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
            log.Printf("TLS Handshake: %v\n", time.Since(tlsHandshake))
        },

        ConnectStart: func(network, addr string) { connect = time.Now() },
        ConnectDone: func(network, addr string, err error) {
            log.Printf("Connect time: %v\n", time.Since(connect))
        },

        GotFirstResponseByte: func() {
            log.Printf("Time from start to first byte: %v\n", time.Since(start))
        },
    }
}