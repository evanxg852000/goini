## GoIni

A very simple go package for parsing init files


```go
package main

import (
	"fmt"
	"strings"

	goini "github.com/evanxg852000/goini"
)

const config string = `
; last modified by John Doe
name=John Doe
age= 23

[owner]
name=John Doe
organization=Acme Widgets Inc.

[database]
; server IP address in case ...
server=192.0.2.62     
port=9080
file="payroll.dat"
`

func main() {
	f, err := goini.NewIniFile(strings.NewReader(config))
	if err != nil {
		fmt.Println("Parser Error ", err)
	}

	// currently at root section
	fmt.Println(f.Get("name")) // John Doe

	f.MoveSection("database")          // navigate to database section
	fmt.Println(f.Get("server"))       //  192.0.2.62
	fmt.Println(f.Get("port"))         //  9080
	fmt.Println(f.Get("organization")) //  empty string as organization is defined in another section

	f.MoveSection("owner")             // navigate to owner section
	fmt.Println(f.Get("organization")) // Acme Widgets Inc.

	f.ResetSection()          // navigate back to root section
	fmt.Println(f.Get("age")) // 23

}
```
