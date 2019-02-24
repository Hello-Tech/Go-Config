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

func TestComments(t *testing.T) {
	path := getCurrentPath() + "/testdata/default.ini"
	log.Println("ini file is ", getCurrentPath())
	iniReader := NewIniReader()
	if iniReader.LoadIni(path) {
		log.Println("load success!")

		testV := iniReader.GetGlobalKey("Age")
		if testV == nil {
			t.Log("Age is comments, success!!")
		} else {
			t.Fatal("Age key exist , failed")
		}

		testV = iniReader.GetGlobalKey(";Age")
		if testV == nil {
			t.Log(";Age is not exist, success!!")
		} else {
			t.Fatal(";Age exist , failed")
		}

		testV = iniReader.GetGlobalKey("Name")
		if testV == nil {
			t.Log("Hello is comments, success!!")
		} else {
			t.Fatal("Name exist , failed")
		}

		testV = iniReader.GetGlobalKey("#Name")
		if testV == nil {
			t.Log("Hello is comments, success!!")
		} else {
			t.Fatal("#Name is a key, failed")
		}

	} else {
		log.Fatal("load ini failed")
	}
}
