package goini

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

const ()

// IniFileAccessor  is a type for access ini file instance
type IniFileAccessor interface {
	MoveSection(name string)
	ResetSection()
	Get(key string) string
	Set(key string, value string)
	ToString() string
}

// IniFile is an object holding ini file content
type IniFile struct {
	base     map[string]string            // base store
	sections map[string]map[string]string //section store
	section  string                       // current browsing section
}

// NewIniFile returns an instace of IniFileAccessor
func NewIniFile(r io.Reader) (IniFileAccessor, error) {
	ini := &IniFile{
		base:     make(map[string]string),
		sections: make(map[string]map[string]string),
	}
	err := ini.parse(r)
	if err != nil {
		return nil, err
	}

	return ini, nil
}

func (ini *IniFile) parse(r io.Reader) error {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	section := ""
	for i, line := range lines {
		line = strings.Trim(line, " ")

		// empty or comment
		if len(line) == 0 || strings.HasPrefix(line, ";") {
			continue
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			section = strings.Trim(line, "[]")
			ini.sections[section] = make(map[string]string)
			continue
		}

		if strings.ContainsAny(line, "=") != true {
			return fmt.Errorf("Invalid key value pair found @ %v", i+1)
		}

		loc := strings.Index(line, "=")
		key := strings.Trim(line[0:loc], " ")
		val := strings.Trim(line[loc+1:], " ")
		if strings.HasPrefix(val, "\"") && strings.HasSuffix(val, "\"") {
			val = strings.Trim(val, "\"")
		}
		if section != "" {
			ini.sections[section][key] = val
			continue
		}
		ini.base[key] = val
	}

	return nil
}

// MoveSection sets the current navigation section
func (ini *IniFile) MoveSection(name string) {
	ini.section = name
}

// ResetSection clears the current navigation section
func (ini *IniFile) ResetSection() {
	ini.section = ""
}

// Get gets a value fron the ini file
func (ini *IniFile) Get(key string) string {
	if ini.section != "" {
		return ini.sections[ini.section][key]
	}
	return ini.base[key]
}

// Set sets a value on the ini file
func (ini *IniFile) Set(key string, value string) {
	if ini.section != "" {
		ini.sections[ini.section][key] = value
	} else {
		ini.base[key] = value
	}
}

// ToString returns a string reprensentatioon of the init file
func (ini *IniFile) ToString() string {
	content := "; main \n"
	for k, v := range ini.base {
		content += fmt.Sprintf("%v=%v \n", k, v)
	}
	content += "; sections \n"
	for s, v := range ini.sections {
		content += fmt.Sprintf("[%v] \n", s)
		for vk, vv := range v {
			content += fmt.Sprintf("%v=%v \n", vk, vv)
		}
	}
	return content
}
