package main

import (
	"AvailableVpn/internal/domain"
	"fmt"
)

func main() {
	err := domain.DownloadFile("https://ipspeed.info/ovpn/219.100.37.83_udp_1195.ovpn", "./ovpn/test.ovpn")
	if err != nil {
		fmt.Println(err)
	}
}
