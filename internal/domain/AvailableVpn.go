package domain

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

type OvpnRepo struct {
	ovpns map[string]map[string][]*OvpnFile
}

func (rep *OvpnRepo) getOvpnsByParam(country, proto string) []*OvpnFile {
	res := make([]*OvpnFile, 0)
	copy(res, rep.ovpns[country][proto])
	return res
}

func DownloadAllOvpnFiles() {
	log.Println("Start downloading ovpn files...")
	urls, err := GetDownloadUrlsFromSource()
	wg := new(sync.WaitGroup)
	wg.Add(len(urls))

	if err != nil {
		log.Println("Failed to pull download urls")
	}
	for _, val := range urls {
		go downloadFile(wg, fmt.Sprintf("https://ipspeed.info/%s", val[1:]), val[1:])
	}

	wg.Wait()
	log.Println("All ovpn downloaded...")
}

func downloadFile(wg *sync.WaitGroup, urlFrom, fileTo string) error {
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

func Ping(ipAddres string) bool { //TODO
	isHostAlive := true
	return isHostAlive
}
