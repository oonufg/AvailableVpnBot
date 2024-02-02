package domain

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	OVPNFILE_PAYLOAD_END = "-----BEGIN CERTIFICATE-----"
)

type OvpnFile struct {
	fileName string
	country  string
	protocol string
}

func ParseOvpnFile(path string) *OvpnFile {
	file, err := os.Open(path)
	if err != nil {
		log.Println("Ошибка во время чтения ovpn файла ", path)
		return nil
	}
	defer file.Close()

	ovpnFile := bufio.NewReader(file)
	parsedFile := make(map[string]string)

	for {
		line, _, err := ovpnFile.ReadLine()
		if err == io.EOF || string(line) == OVPNFILE_PAYLOAD_END {
			break
		}
		sepStr := strings.Split(string(line), " ")
		if len(sepStr) >= 2 {
			parsedFile[sepStr[0]] = sepStr[1]
		}
	}
	fmt.Println(parsedFile["remote"])
	return &OvpnFile{
		fileName: strings.Split(path, "/")[1],
		country:  GetCountryByIp(parsedFile["remote"]),
		protocol: parsedFile["proto"],
	}
}

func ParseAllOvpnInDir(dirPath string) []*OvpnFile {
	s, _ := os.ReadDir(dirPath)
	ovpnFiles := make([]*OvpnFile, 0)
	c := make(chan *OvpnFile)
	counter := 0
	for _, val := range s {
		go ParseOvpnFileA(c, fmt.Sprintf("%s/%s", dirPath, val.Name()))
	}

	for {
		select {
		case ovp := <-c:
			ovpnFiles = append(ovpnFiles, ovp)
			counter++
			fmt.Println(ovp)
		}
		if counter == len(s) {
			break
		}
	}
	fmt.Println(len(ovpnFiles))
	return ovpnFiles
}

func ParseOvpnFileA(ch chan *OvpnFile, path string) *OvpnFile {
	ovpnFile := ParseOvpnFile(path)
	ch <- ovpnFile
	return ovpnFile
}
