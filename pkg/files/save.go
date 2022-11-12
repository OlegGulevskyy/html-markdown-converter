package files

import (
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SaveStringAsFile(name string, d string) {
	b := []byte(d)
	err := ioutil.WriteFile(name, b, 0644)
	check(err)
}
