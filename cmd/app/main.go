package main

import (
	"AvailableVpn/internal/domain"
	"fmt"
)

func main() {
	s := domain.GetCountryByIp("58.239.20.112")
	fmt.Println(s)
}
