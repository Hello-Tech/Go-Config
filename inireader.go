package goconfig

import (
	"io/ioutil"
	"log"
	"strings"
)

var globalSection = "____GLOBAL____"

type IniReader struct {
	_defaultContent map[string]string
	_iniContent     map[string]interface{}
}

func NewIniReader() *IniReader {
	return &IniReader{}
}

func (self *IniReader) GetSectionValues(section string) map[string]string {
	if val, ok := self._iniContent[section]; ok {
		return val.(map[string]string)
	}
	return nil
}

func (self *IniReader) GetGlobalKey(key string) *string {
	if v, ok := self._defaultContent[key]; ok {
		return &v
	}
	return nil
}

func (self *IniReader) GetSectionKey(section, key string) *string {
	keyVals := self.GetSectionValues(section)
	if v, ok := keyVals[key]; ok {
		return &v
	}
	return nil
}

// LoadConfig 加载配置文件
func (self *IniReader) LoadIni(filePath string) bool {
	keyVals := self.loadIni(filePath)
	self._iniContent = keyVals
	self._defaultContent = self.GetSectionValues(globalSection)

	if keyVals == nil || len(keyVals) == 0 {
		log.Println("the content for the config is nil!")
		return true
	}

	return true
}

func (self *IniReader) IsValid() bool {
	return true
}

func makesureExistArray(key string, data map[string]interface{}) []interface{} {
	val, ok := data[key]
	if ok {

	}

	val = make([]interface{}, 0)
	data[key] = val
}

func (self *IniReader) loadIni(filePath string) map[string]interface{} {
	result := make(map[string]interface{})

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panic("can not found the ini config at ", filePath)
	}

	content := string(buf)

	lines := strings.Split(content, "\n")
	if len(lines) == 0 {
		return result
	}

	currentSection := globalSection
	currentArr := make(map[string]string, 0)
	for _, line := range lines {
		line = strings.Replace(line, "\r", "", -1)
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, ";") {
			continue
		}

		if strings.HasPrefix(line, "[") {
			if strings.HasSuffix(line, "]") {
				tempSection := line[1 : len(line)-1]
				if len(tempSection) == 0 {
					// 容错，空section
					continue
				}

				// 上一个section 保存
				result[currentSection] = currentArr

				currentSection = line[1 : len(line)-1]
				currentArr = make(map[string]string, 0)
				continue
			} else {
				log.Panic("INI File is not valid, wrong section")
			}
		}

		pos := strings.Index(line, "=")
		if pos > 0 && pos < len(line)-1 {
			key := line[:pos]
			val := line[pos+1:]

			key = strings.TrimSpace(key)
			val = strings.TrimSpace(val)

			if strings.HasSuffix(key, ":") {
				// extend ini, use := as an array
				if len(key) > 1 {

				} else {
					// the key is :?, error data, ignore
				}
			} else {
				currentArr[key] = val
			}
		}
	}

	result[currentSection] = currentArr

	return result
}
