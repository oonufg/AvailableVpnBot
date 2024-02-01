package domain

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	_ "github.com/gocolly/colly/v2"
)

func DownloadAllOvpnFiles() {
	urls, err := GetDownloadUrlsFromSource()
	wg := new(sync.WaitGroup)
	wg.Add(len(urls))

	if err != nil {
		log.Println("Failed to pull download urls")
	}
	for _, val := range urls {
		go DownloadFile(wg, fmt.Sprintf("https://ipspeed.info/%s", val[1:]), val[1:])
	}

	wg.Wait()
	fmt.Println("Downloaded!")

}

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
