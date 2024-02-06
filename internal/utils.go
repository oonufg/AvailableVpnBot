package utils

import (
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func DownloadFile(wg *sync.WaitGroup, urlFrom, fileTo string) error {
	_, err := os.Stat(fileTo)
	var file *os.File
	if os.IsNotExist(err) {
		file, _ = os.OpenFile(fileTo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		return errors.New("File already exists")
	}
	resp, err := http.Get(urlFrom)
	if resp.StatusCode != 200 {
		return errors.New("Failed to download file")
	}
	io.Copy(file, resp.Body)

	defer func() {
		file.Close()
		resp.Body.Close()
		wg.Done()
	}()
	return nil
}

func Ping(target string) bool {
	var isHostAlive bool
	ip, err := net.ResolveIPAddr("ip4", target)
	if err != nil {
		log.Println("Failed to resolve ip..")
		return false
	}
	conn, err := icmp.ListenPacket("udp4", "0.0.0.0")
	if err != nil {
		log.Println("Failed to listen icmp packet..")
		return false
	}
	defer conn.Close()

	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte(""),
		},
	}
	msg_bytes, err := msg.Marshal(nil)
	if err != nil {
		return false
	}

	// Write the message to the listening connection
	if _, err := conn.WriteTo(msg_bytes, &net.UDPAddr{IP: net.ParseIP(ip.String())}); err != nil {
		return false
	}

	err = conn.SetReadDeadline(time.Now().Add(time.Second * 1))
	if err != nil {
		return false
	}
	reply := make([]byte, 1500)
	n, _, err := conn.ReadFrom(reply)

	if err != nil {
		return false
	}
	parsed_reply, err := icmp.ParseMessage(1, reply[:n])

	if err != nil {
		return false
	}

	switch parsed_reply.Code {
	case 0:
		isHostAlive = true
	default:
		isHostAlive = false
	}
	return isHostAlive
}
