package domain

import (
	utils "AvailableVpn/internal"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

const (
	OVPNFILE_PAYLOAD_END = "-----BEGIN CERTIFICATE-----"
)

type OvpnFile struct {
	fileName string
	country  string
	protocol string
	ip       string
}

func (ovpn *OvpnFile) GetFilename() string {
	return ovpn.fileName
}

func (ovpn *OvpnFile) GetCountry() string {
	return ovpn.country
}

func (ovpn *OvpnFile) GetProtocol() string {
	return ovpn.protocol
}

type OvpnRepo struct {
	ovpns map[string]map[string][]*OvpnFile
}

func CreateOvpnRepository() *OvpnRepo {
	ovpnRepo := &OvpnRepo{}
	ovpnRepo.load()
	return ovpnRepo
}

func (rep *OvpnRepo) load() {
	mmap := make(map[string]map[string][]*OvpnFile)
	ovpns := ParseAllOvpnInDir("./ovpn")
	for _, val := range ovpns {
		if _, ok := mmap[val.country]; !ok {
			mmap[val.country] = make(map[string][]*OvpnFile)
		}
		if _, ok := mmap[val.protocol]; !ok {
			mmap[val.country][val.protocol] = make([]*OvpnFile, 0)
		}
		if utils.Ping(val.ip) {
			mmap[val.country][val.protocol] = append(mmap[val.country][val.protocol], val)
		}
	}
	rep.ovpns = mmap
}

func (rep *OvpnRepo) GetOvpnsByParam(country, proto string) []*OvpnFile {
	res := make([]*OvpnFile, 0)
	copy(res, rep.ovpns[country][proto])
	return rep.ovpns[country][proto]
}

func (rep *OvpnRepo) GetAvailableCountries() []string {
	allCountries := make([]string, 0)
	for key := range rep.ovpns {
		allCountries = append(allCountries, key)
	}

	return allCountries
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
		go utils.DownloadFile(wg, fmt.Sprintf("https://ipspeed.info/%s", val[1:]), val[1:])
	}

	wg.Wait()
	log.Println("All ovpn downloaded...")
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
		}
		if counter == len(s) {
			break
		}
	}
	return ovpnFiles
}

func ParseOvpnFileA(ch chan *OvpnFile, path string) *OvpnFile {
	ovpnFile := ParseOvpnFile(path)
	ch <- ovpnFile
	return ovpnFile
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
	return &OvpnFile{
		fileName: strings.Split(path, "/")[2],
		country:  GetCountryByIp(parsedFile["remote"]),
		protocol: parsedFile["proto"],
		ip:       parsedFile["remote"],
	}
}
