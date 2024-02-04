package utils

import (
	"errors"
	"io"
	"net/http"
	"os"
	"sync"
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

func Ping(ipAddres string) bool { //TODO
	isHostAlive := true
	return isHostAlive
}
