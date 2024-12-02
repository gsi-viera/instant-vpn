package internal

type ClientConfig struct {
	Interface struct {
		Address string
		DNS     string
		MTU     string
	}
	Peer struct {
		PublicKey  string
		AllowedIPs string
		EndPoint   string
	}
}
