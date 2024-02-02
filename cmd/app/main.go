package main

import (
	"AvailableVpn/internal/domain"
	"fmt"
)

func main() {
	o := domain.ParseAllOvpnInDir("./ovpn")
	for _, v := range o {
		fmt.Println(v)
	}
	fmt.Println("Выполнено")
}
