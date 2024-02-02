package main

import (
	cfg "AvailableVpn/config"
	"AvailableVpn/internal/domain"
	"fmt"
)

func main() {
	c := cfg.LoadConfig()
	fmt.Println(c)
	r := domain.Ping("126.0.0.1")
	fmt.Println(r)
}
