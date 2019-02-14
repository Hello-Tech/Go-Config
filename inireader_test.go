package goconfig

import (
	"log"
	"path"
	"runtime"
	"strings"
	"testing"
)

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func TestHello(t *testing.T) {
	path := getCurrentPath() + "/testdata/default.ini"
	log.Println("ini file is ", getCurrentPath())
	iniReader := NewIniReader()
	if iniReader.LoadIni(path) {
		log.Println("load success!")

		port := iniReader.GetGlobalKey("port")
		if port == nil {
			log.Fatal("read port error")
		}

		if strings.Compare(*port, "212") != 0 {
			log.Fatal("read port error, not eq ")
		}

	} else {
		log.Fatal("load ini failed")
	}
}
