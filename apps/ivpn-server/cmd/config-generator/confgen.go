package main

import (
	"os"
	"text/template"

	"github.com/instant-vpn/instant-vpn/apps/ivpn-server/internal"
)

func generateFreeAddress() string {
	return "10.0.0.100"
}

func main() {
	print("generating config")
	clientConfig := internal.ClientConfig{
		Interface: struct {
			Address string
			DNS     string
			MTU     string
		}{
			Address: "10.0.0.1/22",
			DNS:     "8.8.8.8, 8.8.4.4",
			MTU:     "1280",
		},
		Peer: struct {
			PublicKey  string
			AllowedIPs string
			EndPoint   string
		}{
			PublicKey:  "public-key",
			AllowedIPs: "0.0.0.0/0, ::/0",
			EndPoint:   "192.168.1.109:51820",
		},
	}
	templatefile := "client_config.tmpl"
	tmpl, err := template.New(templatefile).ParseFiles(templatefile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, clientConfig)
	if err != nil {
		panic(err)
	}
}
