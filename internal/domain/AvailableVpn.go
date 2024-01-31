package domain

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "github.com/gocolly/colly/v2"
)

func DownloadFile(urlFrom, fileTo string) error {
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
	w, _ := io.Copy(file, resp.Body)
	fmt.Println(w)
	defer file.Close()
	defer resp.Body.Close()

	return nil
}
