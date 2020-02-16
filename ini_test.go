package goini

import (
	"strings"
	"testing"
)

func TestIniFileUnits(t *testing.T) {
	tests := map[string][2]string{
		`file="test.csv"`: [2]string{"file", "test.csv"},
		`tes="test.csv`:   [2]string{"tes", `"test.csv`},
		`empty=`:          [2]string{"empty", ""},
		`;comment=`:       [2]string{"comment", ""},
		`enclose="true"`:  [2]string{"enclose", `true`},
	}

	for input, expect := range tests {
		f, err := NewIniFile(strings.NewReader(input))
		if err != nil {
			t.Errorf("Parse Error %v on test [%v] -> [%v] ", err, input, expect[1])
		}
		if f.Get(expect[0]) != expect[1] {
			t.Errorf("Assertion Error on test [%v] -> [%v]", input, expect[1])
		}
	}
}

func TestIniFileSample(t *testing.T) {
	config := `
; last modified by John Doe

name=John Doe

[owner]
name=John Doe
organization=Acme Widgets Inc.

[database]
; use IP address in case network name resolution is not working
server=192.0.2.62     
port=9080
file="payroll.dat"

`
	f, err := NewIniFile(strings.NewReader(config))
	if err != nil {
		t.Errorf("Parser Error %v", err)
	}

	f.MoveSection("database")
	if f.Get("server") != "192.0.2.62" {
		t.Errorf("Get value error [server] ")
	}
	if f.Get("port") != "9080" {
		t.Errorf("Get value error [port] ")
	}
	if f.Get("organization") != "" {
		t.Errorf("Get value error [organization] ")
	}

	f.ResetSection()
	if f.Get("name") != "John Doe" {
		t.Errorf("Get value error [name]")
	}

}
